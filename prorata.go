package quanta

import (
	"fmt"
)

// ProRataTruncateStrategy implements allocation by truncating each proportional share.
//
// Algorithm:
//  1. Allocate proportional shares based on weights
//  2. Quantize each share with RoundDown (truncate)
//  3. Return parts with remaining dust
//
// Guarantees:
//   - sum(Parts.Multiplier()) + Dust.Quantize(rounding).Value.Multiplier() == total.Quantize(rounding).Value.Multiplier()
//   - Dust may be non-zero (contains the truncated amounts)
//   - Each part is rounded down, ensuring no over-allocation
//
// Use case: When you want explicit dust tracking for accounting purposes.
//
// Edge cases:
//   - Zero total → all parts are zero, dust is zero
//   - Single weight → returns quantized total in parts, dust is zero
//   - Negative total → allowed (refund allocation)
type ProRataTruncateStrategy struct{}

// Allocate distributes the total across weighted parts using truncation.
func (s ProRataTruncateStrategy) Allocate(
	total Measure,
	weights []int64,
	rounding Rounding,
) (AllocationResult, error) {
	// Validate weights
	if err := validateWeights(weights); err != nil {
		return AllocationResult{}, err
	}

	spec := total.Unit()

	// Edge case: single weight - just return the quantized total
	if len(weights) == 1 {
		result := total.Quantize(rounding)
		return AllocationResult{
			Parts: []Quantized{result.Value},
			Dust:  NewMeasureFrom(spec, NewDecimalFromInt64(0)),
		}, nil
	}

	// Edge case: zero total - all parts are zero
	if total.Quantity().IsZero() {
		parts := make([]Quantized, len(weights))
		zero := DeprecatedNewQuantizedFromMultiplier(spec, 0)
		for i := range parts {
			parts[i] = zero
		}
		return AllocationResult{
			Parts: parts,
			Dust:  NewMeasureFrom(spec, NewDecimalFromInt64(0)),
		}, nil
	}

	// Calculate total weight
	totalWeight := int64(0)
	for _, w := range weights {
		totalWeight += w
	}

	// Calculate proportional shares and truncate each
	parts := make([]Quantized, len(weights))
	sumMultipliers := int64(0)

	for i, weight := range weights {
		if weight == 0 {
			// Zero weight gets zero share
			parts[i] = DeprecatedNewQuantizedFromMultiplier(spec, 0)
			continue
		}

		// Calculate proportional share: (weight / totalWeight) × total
		weightDecimal := NewDecimalFromInt64(weight)
		totalWeightDecimal := NewDecimalFromInt64(totalWeight)
		proportion, err := weightDecimal.Div(totalWeightDecimal)
		if err != nil {
			return AllocationResult{}, fmt.Errorf("allocation: division failed: %w", err)
		}
		share, err := total.Quantity().Mul(proportion)
		if err != nil {
			return AllocationResult{}, fmt.Errorf("allocation: multiplication failed: %w", err)
		}

		// Quantize the share with RoundDown (truncate)
		shareMeasure := NewMeasureFrom(spec, share)
		result := shareMeasure.Quantize(RoundDown)

		parts[i] = result.Value
		sumMultipliers += result.Value.Multiplier()
	}

	// Calculate target (quantized total)
	targetResult := total.Quantize(rounding)
	target := targetResult.Value

	// Calculate dust: target - sum(parts)
	targetDecimal := target.Decimal()
	sumDecimal := NewDecimalFromInt64(sumMultipliers)
	quantumDecimal := spec.Quantum().Decimal()

	// Convert sum from multipliers to decimal value
	sumValue, err := sumDecimal.Mul(quantumDecimal)
	if err != nil {
		return AllocationResult{}, fmt.Errorf("allocation: multiplication failed: %w", err)
	}

	// Dust = target - sum
	dustDecimal, err := targetDecimal.Sub(sumValue)
	if err != nil {
		return AllocationResult{}, fmt.Errorf("allocation: subtraction failed: %w", err)
	}

	dust := NewMeasureFrom(spec, dustDecimal)

	return AllocationResult{
		Parts: parts,
		Dust:  dust,
	}, nil
}
