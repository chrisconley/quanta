package quanta

import (
	"fmt"
)

// shareInfo tracks allocation information for a single weighted part.
type shareInfo struct {
	index     int
	quantized Quantized
	remainder Measure
}

// LargestRemainderStrategy implements allocation using the largest remainder method.
//
// Algorithm:
//  1. Allocate proportional shares based on weights
//  2. Quantize each share (may create "dust" due to rounding)
//  3. Distribute dust to parts with largest fractional remainders
//
// Guarantees:
//   - sum(Parts.Multiplier()) == total.Quantize(rounding).Value.Multiplier()
//   - Dust == 0 (all dust distributed to parts)
//   - Each part receives at most 1 quantum of dust
//
// Tiebreaking: When remainders are equal, earlier indices win.
//
// Edge cases:
//   - Zero total → all parts are zero
//   - Single weight → optimized (just quantize total)
//   - Negative total → allowed (refund allocation)
type LargestRemainderStrategy struct{}

// Allocate distributes the total across weighted parts using largest remainder method.
func (s LargestRemainderStrategy) Allocate(
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

	// Calculate proportional shares and track remainders
	shares := make([]shareInfo, len(weights))
	sumMultipliers := int64(0)

	for i, weight := range weights {
		if weight == 0 {
			// Zero weight gets zero share
			shares[i] = shareInfo{
				index:     i,
				quantized: DeprecatedNewQuantizedFromMultiplier(spec, 0),
				remainder: NewMeasureFrom(spec, NewDecimalFromInt64(0)),
			}
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

		// Quantize the share
		shareMeasure := NewMeasureFrom(spec, share)
		result := shareMeasure.Quantize(rounding)

		shares[i] = shareInfo{
			index:     i,
			quantized: result.Value,
			remainder: result.Remainder,
		}

		sumMultipliers += result.Value.Multiplier()
	}

	// Calculate target (quantized total)
	targetResult := total.Quantize(rounding)
	target := targetResult.Value
	targetMultiplier := target.Multiplier()

	// Calculate dust in terms of quanta
	dustQuanta := targetMultiplier - sumMultipliers

	// Distribute dust to parts with largest remainders
	// Sort indices by remainder.Frac() descending, tiebreak by index ascending
	if dustQuanta != 0 {
		// Find parts to receive dust
		indices := make([]int, len(shares))
		for i := range indices {
			indices[i] = i
		}

		// Sort by remainder fractional part (descending), then by index (ascending)
		sortByRemainderDesc(indices, shares)

		// Distribute dust one quantum at a time
		quantum := spec.Quantum().Decimal()
		dustRemaining := dustQuanta
		if dustRemaining < 0 {
			// Negative dust - subtract quanta (rounding went too high)
			quantum = quantum.Negate()
			dustRemaining = -dustRemaining
		}

		for i := 0; i < int(dustRemaining) && i < len(indices); i++ {
			idx := indices[i]
			// Add/subtract one quantum to this part
			newMultiplier := shares[idx].quantized.Multiplier()
			if dustQuanta > 0 {
				newMultiplier++
			} else {
				newMultiplier--
			}
			shares[idx].quantized = DeprecatedNewQuantizedFromMultiplier(spec, newMultiplier)
		}
	}

	// Extract final parts
	parts := make([]Quantized, len(shares))
	for i, share := range shares {
		parts[i] = share.quantized
	}

	// Verify invariant: sum(parts) == target
	finalSum := int64(0)
	for _, part := range parts {
		finalSum += part.Multiplier()
	}
	if finalSum != targetMultiplier {
		return AllocationResult{}, fmt.Errorf(
			"allocation invariant violated: sum(parts)=%d != target=%d",
			finalSum, targetMultiplier,
		)
	}

	return AllocationResult{
		Parts: parts,
		Dust:  NewMeasureFrom(spec, NewDecimalFromInt64(0)), // Always zero for LargestRemainder
	}, nil
}

// sortByRemainderDesc sorts indices by remainder fractional part (descending),
// with ties broken by index (ascending - earlier indices win).
func sortByRemainderDesc(indices []int, shares []shareInfo) {
	// Simple insertion sort - adequate for typical allocation counts
	for i := 1; i < len(indices); i++ {
		key := indices[i]
		keyRemainder := shares[key].remainder.Quantity().Abs().Frac()
		j := i - 1

		for j >= 0 {
			currentRemainder := shares[indices[j]].remainder.Quantity().Abs().Frac()
			// Compare: want descending remainder, ascending index
			if currentRemainder.Cmp(keyRemainder) > 0 {
				// Current has larger remainder - keep it first
				break
			}
			if currentRemainder.Equal(keyRemainder) && indices[j] < key {
				// Tie - earlier index wins
				break
			}
			// Move current element forward
			indices[j+1] = indices[j]
			j--
		}
		indices[j+1] = key
	}
}
