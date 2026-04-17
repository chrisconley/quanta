package quanta

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestAllocation_InvoiceSplit demonstrates allocating an invoice total across multiple line items
// with different weights (quantities).
func TestAllocation_InvoiceSplit(t *testing.T) {
	usdSpec := MustNewUnit("USD", "0.01")
	strategy := LargestRemainderStrategy{}

	// Invoice total: $127.47 to be split among 3 items with quantities [5, 3, 2]
	// Proportions: 50%, 30%, 20%
	total := NewMeasureFrom(usdSpec, MustNewDecimal("127.47"))
	weights := []int64{5, 3, 2}

	result, err := strategy.Allocate(total, weights, RoundHalfEven)
	assert.NoError(t, err)

	// Verify parts
	assert.Equal(t, 3, len(result.Parts))

	// Expected: $63.74 (50%), $38.24 (30%), $25.49 (20%)
	// Due to rounding, one part may get adjusted
	expectedAmounts := []string{"63.74", "38.24", "25.49"}
	for i, expected := range expectedAmounts {
		actual := result.Parts[i].Decimal()
		// Allow for small rounding adjustments
		if !actual.Equal(MustNewDecimal(expected)) {
			diff, _ := actual.Sub(MustNewDecimal(expected))
			// Difference should be at most 1 quantum ($0.01)
			if diff.Abs().Cmp(MustNewDecimal("0.01")) > 0 {
				t.Errorf("Part[%d] = %v, expected ~%s", i, actual, expected)
			}
		}
	}

	// Verify dust is zero (LargestRemainder distributes all dust)
	assert.True(t, result.Dust.Quantity().IsZero(), "Dust should be zero")

	// Verify invariant: sum(parts) == total
	target := total.Quantize(RoundHalfEven).Value
	sumMultipliers := int64(0)
	for _, part := range result.Parts {
		sumMultipliers += part.Multiplier()
	}
	assert.Equal(t, target.Multiplier(), sumMultipliers, "sum(parts) must equal target")
}

// TestAllocation_RevenueSplit demonstrates revenue sharing among partners.
func TestAllocation_RevenueSplit(t *testing.T) {
	usdSpec := MustNewUnit("USD", "0.01")
	strategy := LargestRemainderStrategy{}

	// Monthly revenue: $50,000.00 to be split among 4 partners
	// Shares: 40%, 30%, 20%, 10% (represented as weights: 4, 3, 2, 1)
	total := NewMeasureFrom(usdSpec, MustNewDecimal("50000.00"))
	weights := []int64{4, 3, 2, 1}

	result, err := strategy.Allocate(total, weights, RoundHalfEven)
	assert.NoError(t, err)

	// Verify parts
	assert.Equal(t, 4, len(result.Parts))

	// Expected: $20,000.00 (40%), $15,000.00 (30%), $10,000.00 (20%), $5,000.00 (10%)
	expectedAmounts := []string{"20000.00", "15000.00", "10000.00", "5000.00"}
	for i, expected := range expectedAmounts {
		actual := result.Parts[i].Decimal()
		assert.True(t, actual.Equal(MustNewDecimal(expected)),
			"Part[%d] = %v, want %s", i, actual, expected)
	}

	// Verify dust is zero
	assert.True(t, result.Dust.Quantity().IsZero(), "Dust should be zero")

	// Verify invariant
	target := total.Quantize(RoundHalfEven).Value
	sumMultipliers := int64(0)
	for _, part := range result.Parts {
		sumMultipliers += part.Multiplier()
	}
	assert.Equal(t, target.Multiplier(), sumMultipliers)
}

// TestAllocation_RefundAllocation demonstrates allocating a refund (negative amount).
func TestAllocation_RefundAllocation(t *testing.T) {
	usdSpec := MustNewUnit("USD", "0.01")
	strategy := LargestRemainderStrategy{}

	// Refund total: -$89.97 to be split among 3 customers equally
	total := NewMeasureFrom(usdSpec, MustNewDecimal("-89.97"))
	weights := []int64{1, 1, 1}

	result, err := strategy.Allocate(total, weights, RoundHalfEven)
	assert.NoError(t, err)

	// Verify parts
	assert.Equal(t, 3, len(result.Parts))

	// Each should get ~-$29.99
	for i, part := range result.Parts {
		actual := part.Decimal()
		// Should be negative and close to -$29.99
		assert.True(t, actual.IsNegative(), "Part[%d] should be negative", i)

		diff, _ := actual.Add(MustNewDecimal("29.99"))
		// Allow for 1 quantum difference due to rounding
		if diff.Abs().Cmp(MustNewDecimal("0.01")) > 0 {
			t.Errorf("Part[%d] = %v, expected ~-$29.99", i, actual)
		}
	}

	// Verify dust is zero
	assert.True(t, result.Dust.Quantity().IsZero())

	// Verify invariant
	target := total.Quantize(RoundHalfEven).Value
	sumMultipliers := int64(0)
	for _, part := range result.Parts {
		sumMultipliers += part.Multiplier()
	}
	assert.Equal(t, target.Multiplier(), sumMultipliers)
}

// TestAllocation_LargeScale demonstrates allocation across many parts.
func TestAllocation_LargeScale(t *testing.T) {
	usdSpec := MustNewUnit("USD", "0.01")
	strategy := LargestRemainderStrategy{}

	// Allocate $1,000,000.00 across 1000 equal parts
	total := NewMeasureFrom(usdSpec, MustNewDecimal("1000000.00"))
	weights := make([]int64, 1000)
	for i := range weights {
		weights[i] = 1
	}

	result, err := strategy.Allocate(total, weights, RoundHalfEven)
	assert.NoError(t, err)

	// Verify parts count
	assert.Equal(t, 1000, len(result.Parts))

	// Each part should be exactly $1,000.00
	expected := MustNewDecimal("1000.00")
	for i, part := range result.Parts {
		if !part.Decimal().Equal(expected) {
			t.Errorf("Part[%d] = %v, want $1,000.00", i, part.Decimal())
		}
	}

	// Verify dust is zero
	assert.True(t, result.Dust.Quantity().IsZero())

	// Verify invariant
	target := total.Quantize(RoundHalfEven).Value
	sumMultipliers := int64(0)
	for _, part := range result.Parts {
		sumMultipliers += part.Multiplier()
	}
	assert.Equal(t, target.Multiplier(), sumMultipliers)
}

// TestAllocation_ProRataWithDustAccount demonstrates using ProRataTruncateStrategy
// when explicit dust tracking is needed.
func TestAllocation_ProRataWithDustAccount(t *testing.T) {
	usdSpec := MustNewUnit("USD", "0.01")
	strategy := ProRataTruncateStrategy{}

	// Allocate $100.00 across 7 equal parts
	// Each part: $14.2857... -> truncate to $14.28
	// Total distributed: 7 x $14.28 = $99.96
	// Dust: $0.04
	total := NewMeasureFrom(usdSpec, MustNewDecimal("100.00"))
	weights := []int64{1, 1, 1, 1, 1, 1, 1}

	result, err := strategy.Allocate(total, weights, RoundHalfEven)
	assert.NoError(t, err)

	// Each part should be $14.28 (truncated)
	expected := MustNewDecimal("14.28")
	for i, part := range result.Parts {
		assert.True(t, part.Decimal().Equal(expected),
			"Part[%d] = %v, want $14.28", i, part.Decimal())
	}

	// Dust should be $0.04
	assert.True(t, result.Dust.Quantity().Equal(MustNewDecimal("0.04")),
		"Dust = %v, want $0.04", result.Dust.Quantity())

	// Verify invariant: sum(parts) + dust == target
	target := total.Quantize(RoundHalfEven).Value
	sumMultipliers := int64(0)
	for _, part := range result.Parts {
		sumMultipliers += part.Multiplier()
	}
	dustQuantized := result.Dust.Quantize(RoundHalfEven).Value
	totalDistributed := sumMultipliers + dustQuantized.Multiplier()

	assert.Equal(t, target.Multiplier(), totalDistributed)
}

// TestAllocation_StrategyComparison compares LargestRemainder vs ProRataTruncate.
func TestAllocation_StrategyComparison(t *testing.T) {
	usdSpec := MustNewUnit("USD", "0.01")

	// Same input for both strategies
	total := NewMeasureFrom(usdSpec, MustNewDecimal("10.00"))
	weights := []int64{1, 1, 1}

	// LargestRemainder: distributes all dust to parts
	lrStrategy := LargestRemainderStrategy{}
	lrResult, err := lrStrategy.Allocate(total, weights, RoundHalfEven)
	assert.NoError(t, err)

	// ProRataTruncate: returns dust separately
	prStrategy := ProRataTruncateStrategy{}
	prResult, err := prStrategy.Allocate(total, weights, RoundHalfEven)
	assert.NoError(t, err)

	// LargestRemainder: parts sum to total, dust is zero
	lrSum := int64(0)
	for _, part := range lrResult.Parts {
		lrSum += part.Multiplier()
	}
	target := total.Quantize(RoundHalfEven).Value
	assert.Equal(t, target.Multiplier(), lrSum, "LargestRemainder: parts should sum to total")
	assert.True(t, lrResult.Dust.Quantity().IsZero(), "LargestRemainder: dust should be zero")

	// ProRataTruncate: parts + dust sum to total
	prSum := int64(0)
	for _, part := range prResult.Parts {
		prSum += part.Multiplier()
	}
	prDust := prResult.Dust.Quantize(RoundHalfEven).Value
	assert.Equal(t, target.Multiplier(), prSum+prDust.Multiplier(),
		"ProRataTruncate: parts + dust should sum to total")
	assert.False(t, prResult.Dust.Quantity().IsZero(), "ProRataTruncate: dust should be non-zero")

	// LargestRemainder gives more to some parts to eliminate dust
	// ProRataTruncate gives less to all parts and tracks dust separately
	lrMax := int64(0)
	for _, part := range lrResult.Parts {
		if part.Multiplier() > lrMax {
			lrMax = part.Multiplier()
		}
	}

	prMax := int64(0)
	for _, part := range prResult.Parts {
		if part.Multiplier() > prMax {
			prMax = part.Multiplier()
		}
	}

	// At least one LR part should be larger than all PR parts (due to dust distribution)
	assert.True(t, lrMax > prMax, "LargestRemainder should give more to at least one part")
}

// TestAllocation_EdgeCase_VerySmallTotal demonstrates allocation of very small amounts.
func TestAllocation_EdgeCase_VerySmallTotal(t *testing.T) {
	usdSpec := MustNewUnit("USD", "0.01")
	strategy := LargestRemainderStrategy{}

	// Allocate $0.03 across 5 parts
	// Only 3 parts can get $0.01, others get $0.00
	total := NewMeasureFrom(usdSpec, MustNewDecimal("0.03"))
	weights := []int64{1, 1, 1, 1, 1}

	result, err := strategy.Allocate(total, weights, RoundHalfEven)
	assert.NoError(t, err)

	// Count parts with non-zero amounts
	nonZero := 0
	for _, part := range result.Parts {
		if !part.Decimal().IsZero() {
			nonZero++
			assert.True(t, part.Decimal().Equal(MustNewDecimal("0.01")),
				"Non-zero parts should be $0.01")
		}
	}

	// Exactly 3 parts should be non-zero
	assert.Equal(t, 3, nonZero, "Exactly 3 parts should get $0.01")

	// Verify dust is zero
	assert.True(t, result.Dust.Quantity().IsZero())

	// Verify invariant
	target := total.Quantize(RoundHalfEven).Value
	sumMultipliers := int64(0)
	for _, part := range result.Parts {
		sumMultipliers += part.Multiplier()
	}
	assert.Equal(t, target.Multiplier(), sumMultipliers)
}
