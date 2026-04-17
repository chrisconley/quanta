package quanta

import (
	"testing"
)

func TestLargestRemainderStrategy_Allocate(t *testing.T) {
	t.Run("with equal weights", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		strategy := LargestRemainderStrategy{}

		// Allocate $10.00 across 3 equal parts
		total := NewMeasureFrom(spec, MustNewDecimal("10.00"))
		weights := []int64{1, 1, 1}

		result, err := strategy.Allocate(total, weights, RoundHalfEven)
		if err != nil {
			t.Fatalf("Allocate() error: %v", err)
		}

		// Check number of parts
		if len(result.Parts) != 3 {
			t.Errorf("Parts length = %d, want 3", len(result.Parts))
		}

		// Check sum equals total
		target := total.Quantize(RoundHalfEven).Value
		sumMultipliers := int64(0)
		for _, part := range result.Parts {
			sumMultipliers += part.Multiplier()
		}
		if sumMultipliers != target.Multiplier() {
			t.Errorf("sum(parts) = %d, want %d", sumMultipliers, target.Multiplier())
		}

		// Check dust is zero
		if !result.Dust.Quantity().IsZero() {
			t.Errorf("Dust = %v, want zero", result.Dust.Quantity())
		}

		// Expected: $3.33, $3.33, $3.34 (first part gets extra penny due to remainder tie)
		// Each part should be close to $3.33
		for i, part := range result.Parts {
			decimal := part.Decimal()
			if decimal.Cmp(MustNewDecimal("3.32")) <= 0 || decimal.Cmp(MustNewDecimal("3.35")) >= 0 {
				t.Errorf("Part[%d] = %v, expected ~$3.33", i, decimal)
			}
		}
	})

	t.Run("with unequal weights", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		strategy := LargestRemainderStrategy{}

		// Allocate $10.00 with weights [5, 3, 2]
		total := NewMeasureFrom(spec, MustNewDecimal("10.00"))
		weights := []int64{5, 3, 2}

		result, err := strategy.Allocate(total, weights, RoundHalfEven)
		if err != nil {
			t.Fatalf("Allocate() error: %v", err)
		}

		// Check invariant: sum(parts) == total
		target := total.Quantize(RoundHalfEven).Value
		sumMultipliers := int64(0)
		for _, part := range result.Parts {
			sumMultipliers += part.Multiplier()
		}
		if sumMultipliers != target.Multiplier() {
			t.Errorf("sum(parts) = %d, want %d", sumMultipliers, target.Multiplier())
		}

		// Check dust is zero
		if !result.Dust.Quantity().IsZero() {
			t.Errorf("Dust = %v, want zero", result.Dust.Quantity())
		}

		// Expected proportions: 5/10 = $5.00, 3/10 = $3.00, 2/10 = $2.00
		expected := []string{"5.00", "3.00", "2.00"}
		for i, part := range result.Parts {
			decimal := part.Decimal()
			if !decimal.Equal(MustNewDecimal(expected[i])) {
				t.Errorf("Part[%d] = %v, want %s", i, decimal, expected[i])
			}
		}
	})

	t.Run("with single weight", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		strategy := LargestRemainderStrategy{}

		// Allocate $12.3456 with single weight
		total := NewMeasureFrom(spec, MustNewDecimal("12.3456"))
		weights := []int64{1}

		result, err := strategy.Allocate(total, weights, RoundHalfEven)
		if err != nil {
			t.Fatalf("Allocate() error: %v", err)
		}

		// Should just return quantized total
		expected := total.Quantize(RoundHalfEven).Value
		if !result.Parts[0].Decimal().Equal(expected.Decimal()) {
			t.Errorf("Part[0] = %v, want %v", result.Parts[0].Decimal(), expected.Decimal())
		}

		// Dust should be zero
		if !result.Dust.Quantity().IsZero() {
			t.Errorf("Dust = %v, want zero", result.Dust.Quantity())
		}
	})

	t.Run("with zero total", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		strategy := LargestRemainderStrategy{}

		// Allocate $0.00
		total := NewMeasureFrom(spec, MustNewDecimal("0"))
		weights := []int64{1, 1, 1}

		result, err := strategy.Allocate(total, weights, RoundHalfEven)
		if err != nil {
			t.Fatalf("Allocate() error: %v", err)
		}

		// All parts should be zero
		for i, part := range result.Parts {
			if part.Multiplier() != 0 {
				t.Errorf("Part[%d] = %d, want 0", i, part.Multiplier())
			}
		}

		// Dust should be zero
		if !result.Dust.Quantity().IsZero() {
			t.Errorf("Dust = %v, want zero", result.Dust.Quantity())
		}
	})

	t.Run("with negative total (refund)", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		strategy := LargestRemainderStrategy{}

		// Allocate -$10.00 (refund)
		total := NewMeasureFrom(spec, MustNewDecimal("-10.00"))
		weights := []int64{1, 1, 1}

		result, err := strategy.Allocate(total, weights, RoundHalfEven)
		if err != nil {
			t.Fatalf("Allocate() error: %v", err)
		}

		// Check invariant
		target := total.Quantize(RoundHalfEven).Value
		sumMultipliers := int64(0)
		for _, part := range result.Parts {
			sumMultipliers += part.Multiplier()
		}
		if sumMultipliers != target.Multiplier() {
			t.Errorf("sum(parts) = %d, want %d", sumMultipliers, target.Multiplier())
		}

		// All parts should be negative
		for i, part := range result.Parts {
			if part.Multiplier() >= 0 {
				t.Errorf("Part[%d] = %d, want negative", i, part.Multiplier())
			}
		}

		// Dust should be zero
		if !result.Dust.Quantity().IsZero() {
			t.Errorf("Dust = %v, want zero", result.Dust.Quantity())
		}
	})

	t.Run("with very unequal weights", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		strategy := LargestRemainderStrategy{}

		// Allocate $10.00 with very unequal weights [1, 1000000]
		total := NewMeasureFrom(spec, MustNewDecimal("10.00"))
		weights := []int64{1, 1000000}

		result, err := strategy.Allocate(total, weights, RoundHalfEven)
		if err != nil {
			t.Fatalf("Allocate() error: %v", err)
		}

		// Check invariant
		target := total.Quantize(RoundHalfEven).Value
		sumMultipliers := int64(0)
		for _, part := range result.Parts {
			sumMultipliers += part.Multiplier()
		}
		if sumMultipliers != target.Multiplier() {
			t.Errorf("sum(parts) = %d, want %d", sumMultipliers, target.Multiplier())
		}

		// First part should be tiny (almost zero)
		if result.Parts[0].Multiplier() > 1 {
			t.Errorf("Part[0] = %d, want 0 or 1 (very small)", result.Parts[0].Multiplier())
		}

		// Second part should be almost the entire total
		if result.Parts[1].Decimal().Cmp(MustNewDecimal("9.99")) < 0 {
			t.Errorf("Part[1] = %v, want ~$10.00", result.Parts[1].Decimal())
		}

		// Dust should be zero
		if !result.Dust.Quantity().IsZero() {
			t.Errorf("Dust = %v, want zero", result.Dust.Quantity())
		}
	})

	t.Run("with many parts", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		strategy := LargestRemainderStrategy{}

		// Allocate $100.00 across 100 equal parts
		total := NewMeasureFrom(spec, MustNewDecimal("100.00"))
		weights := make([]int64, 100)
		for i := range weights {
			weights[i] = 1
		}

		result, err := strategy.Allocate(total, weights, RoundHalfEven)
		if err != nil {
			t.Fatalf("Allocate() error: %v", err)
		}

		// Check invariant
		target := total.Quantize(RoundHalfEven).Value
		sumMultipliers := int64(0)
		for _, part := range result.Parts {
			sumMultipliers += part.Multiplier()
		}
		if sumMultipliers != target.Multiplier() {
			t.Errorf("sum(parts) = %d, want %d", sumMultipliers, target.Multiplier())
		}

		// Each part should be $1.00
		for i, part := range result.Parts {
			if !part.Decimal().Equal(MustNewDecimal("1.00")) {
				t.Errorf("Part[%d] = %v, want $1.00", i, part.Decimal())
			}
		}

		// Dust should be zero
		if !result.Dust.Quantity().IsZero() {
			t.Errorf("Dust = %v, want zero", result.Dust.Quantity())
		}
	})

	t.Run("with some zero weights", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		strategy := LargestRemainderStrategy{}

		// Allocate $10.00 with weights [5, 0, 5]
		total := NewMeasureFrom(spec, MustNewDecimal("10.00"))
		weights := []int64{5, 0, 5}

		result, err := strategy.Allocate(total, weights, RoundHalfEven)
		if err != nil {
			t.Fatalf("Allocate() error: %v", err)
		}

		// Check invariant
		target := total.Quantize(RoundHalfEven).Value
		sumMultipliers := int64(0)
		for _, part := range result.Parts {
			sumMultipliers += part.Multiplier()
		}
		if sumMultipliers != target.Multiplier() {
			t.Errorf("sum(parts) = %d, want %d", sumMultipliers, target.Multiplier())
		}

		// Middle part should be zero
		if result.Parts[1].Multiplier() != 0 {
			t.Errorf("Part[1] = %v, want $0.00", result.Parts[1].Decimal())
		}

		// First and third parts should be $5.00 each
		if !result.Parts[0].Decimal().Equal(MustNewDecimal("5.00")) {
			t.Errorf("Part[0] = %v, want $5.00", result.Parts[0].Decimal())
		}
		if !result.Parts[2].Decimal().Equal(MustNewDecimal("5.00")) {
			t.Errorf("Part[2] = %v, want $5.00", result.Parts[2].Decimal())
		}

		// Dust should be zero
		if !result.Dust.Quantity().IsZero() {
			t.Errorf("Dust = %v, want zero", result.Dust.Quantity())
		}
	})

	t.Run("with invalid weights", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		strategy := LargestRemainderStrategy{}
		total := NewMeasureFrom(spec, MustNewDecimal("10.00"))

		tests := []struct {
			name    string
			weights []int64
			wantErr string
		}{
			{
				name:    "empty weights",
				weights: []int64{},
				wantErr: "weights cannot be empty",
			},
			{
				name:    "negative weight",
				weights: []int64{5, -3, 2},
				wantErr: "weight[1] is negative",
			},
			{
				name:    "all zero weights",
				weights: []int64{0, 0, 0},
				wantErr: "all weights are zero",
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				_, err := strategy.Allocate(total, tt.weights, RoundHalfEven)
				if err == nil {
					t.Errorf("Allocate() expected error containing %q, got nil", tt.wantErr)
					return
				}
				if !contains(err.Error(), tt.wantErr) {
					t.Errorf("Allocate() error = %v, want substring %q", err, tt.wantErr)
				}
			})
		}
	})

	t.Run("tiebreaking by index", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		strategy := LargestRemainderStrategy{}

		// Allocate $10.00 with equal weights - all remainders tie
		// First parts should get the dust
		total := NewMeasureFrom(spec, MustNewDecimal("10.00"))
		weights := []int64{1, 1, 1}

		result, err := strategy.Allocate(total, weights, RoundHalfEven)
		if err != nil {
			t.Fatalf("Allocate() error: %v", err)
		}

		// Check invariant
		target := total.Quantize(RoundHalfEven).Value
		sumMultipliers := int64(0)
		for _, part := range result.Parts {
			sumMultipliers += part.Multiplier()
		}
		if sumMultipliers != target.Multiplier() {
			t.Errorf("sum(parts) = %d, want %d", sumMultipliers, target.Multiplier())
		}

		// Expected: one part gets $3.34, others get $3.33
		// Due to tiebreaking, first part should get extra penny
		counts := make(map[int64]int)
		for _, part := range result.Parts {
			counts[part.Multiplier()]++
		}

		// Should have 1 part with 334 cents and 2 parts with 333 cents
		if counts[334] != 1 {
			t.Errorf("Expected 1 part with $3.34 (334 cents), got %d", counts[334])
		}
		if counts[333] != 2 {
			t.Errorf("Expected 2 parts with $3.33 (333 cents), got %d", counts[333])
		}
	})
}
