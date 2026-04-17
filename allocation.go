package quanta

import (
	"fmt"
)

// AllocationStrategy allocates a total amount across multiple weighted parts.
// The parts are quantized such that their sum equals the quantized total.
//
// Strategies handle the "dust" (rounding residuals) differently:
//   - LargestRemainderStrategy: Distributes all dust to parts (Dust == 0)
//   - ProRataTruncateStrategy: Returns dust separately (sum(Parts) + Dust == total)
type AllocationStrategy interface {
	// Allocate distributes the total amount across multiple weighted parts.
	//
	// Parameters:
	//   - total: The amount to allocate
	//   - weights: Integer weights for each part (must be non-negative, sum > 0)
	//   - rounding: Rounding mode for quantization
	//
	// Returns:
	//   - AllocationResult containing allocated parts and any remaining dust
	//   - Error if weights are invalid (empty, negative, or all zero)
	//
	// Guarantees:
	//   Strategy-specific, but generally:
	//   sum(Parts.Multiplier()) + Dust.Quantize(rounding).Value.Multiplier() == total.Quantize(rounding).Value.Multiplier()
	Allocate(
		total Measure,
		weights []int64,
		rounding Rounding,
	) (AllocationResult, error)
}

// AllocationResult contains the allocated parts and any remaining dust.
type AllocationResult struct {
	// Parts are the allocated amounts, one per weight.
	// Length always equals len(weights) passed to Allocate.
	// Each part is quantized to the spec's quantum.
	Parts []Quantized
	// Dust is the remaining amount that couldn't be distributed.
	// Semantics are strategy-dependent:
	//   - LargestRemainderStrategy: Always zero (all dust distributed to Parts)
	//   - ProRataTruncateStrategy: Non-zero (sum(Parts) + Dust == quantized total)
	Dust Measure
}

// validateWeights checks that weights are valid for allocation.
// Returns error if:
//   - weights is empty
//   - any weight is negative
//   - all weights are zero
func validateWeights(weights []int64) error {
	if len(weights) == 0 {
		return fmt.Errorf("allocation: weights cannot be empty")
	}

	totalWeight := int64(0)
	for i, w := range weights {
		if w < 0 {
			return fmt.Errorf("allocation: weight[%d] is negative: %d", i, w)
		}
		totalWeight += w
	}

	if totalWeight == 0 {
		return fmt.Errorf("allocation: all weights are zero")
	}

	return nil
}
