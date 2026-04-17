package quanta

import (
	"testing"
)

func TestProRataTruncateStrategy_Allocate(t *testing.T) {
	t.Run("with equal weights", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		strategy := ProRataTruncateStrategy{}

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

		// With truncation, each part gets $3.33 (floor of $3.333...)
		// Total distributed: $9.99, dust: $0.01
		for i, part := range result.Parts {
			if !part.Decimal().Equal(MustNewDecimal("3.33")) {
				t.Errorf("Part[%d] = %v, want $3.33", i, part.Decimal())
			}
		}

		// Dust should be $0.01
		if !result.Dust.Quantity().Equal(MustNewDecimal("0.01")) {
			t.Errorf("Dust = %v, want $0.01", result.Dust.Quantity())
		}

		// Verify invariant: sum(parts) + dust == quantized total
		target := total.Quantize(RoundHalfEven).Value
		sumMultipliers := int64(0)
		for _, part := range result.Parts {
			sumMultipliers += part.Multiplier()
		}
		dustQuantized := result.Dust.Quantize(RoundHalfEven).Value
		totalDistributed := sumMultipliers + dustQuantized.Multiplier()

		if totalDistributed != target.Multiplier() {
			t.Errorf("sum(parts) + dust = %d, want %d", totalDistributed, target.Multiplier())
		}
	})

	t.Run("with unequal weights", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		strategy := ProRataTruncateStrategy{}

		// Allocate $10.00 with weights [5, 3, 2]
		total := NewMeasureFrom(spec, MustNewDecimal("10.00"))
		weights := []int64{5, 3, 2}

		result, err := strategy.Allocate(total, weights, RoundHalfEven)
		if err != nil {
			t.Fatalf("Allocate() error: %v", err)
		}

		// Expected proportions: 5/10 = $5.00, 3/10 = $3.00, 2/10 = $2.00
		// These divide evenly, so dust should be zero
		expected := []string{"5.00", "3.00", "2.00"}
		for i, part := range result.Parts {
			if !part.Decimal().Equal(MustNewDecimal(expected[i])) {
				t.Errorf("Part[%d] = %v, want %s", i, part.Decimal(), expected[i])
			}
		}

		// Dust should be zero
		if !result.Dust.Quantity().IsZero() {
			t.Errorf("Dust = %v, want zero", result.Dust.Quantity())
		}
	})

	t.Run("with single weight", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		strategy := ProRataTruncateStrategy{}

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
		strategy := ProRataTruncateStrategy{}

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
		strategy := ProRataTruncateStrategy{}

		// Allocate -$10.00 (refund)
		total := NewMeasureFrom(spec, MustNewDecimal("-10.00"))
		weights := []int64{1, 1, 1}

		result, err := strategy.Allocate(total, weights, RoundHalfEven)
		if err != nil {
			t.Fatalf("Allocate() error: %v", err)
		}

		// Verify invariant
		target := total.Quantize(RoundHalfEven).Value
		sumMultipliers := int64(0)
		for _, part := range result.Parts {
			sumMultipliers += part.Multiplier()
		}
		dustQuantized := result.Dust.Quantize(RoundHalfEven).Value
		totalDistributed := sumMultipliers + dustQuantized.Multiplier()

		if totalDistributed != target.Multiplier() {
			t.Errorf("sum(parts) + dust = %d, want %d", totalDistributed, target.Multiplier())
		}

		// All parts should be negative
		for i, part := range result.Parts {
			if part.Multiplier() >= 0 {
				t.Errorf("Part[%d] = %d, want negative", i, part.Multiplier())
			}
		}
	})

	t.Run("with dust accumulation", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		strategy := ProRataTruncateStrategy{}

		// Allocate $1.00 across 7 equal parts
		// Each part: $1.00 / 7 = $0.142857...
		// Truncated: $0.14 each
		// Total distributed: 7 x $0.14 = $0.98
		// Dust: $0.02
		total := NewMeasureFrom(spec, MustNewDecimal("1.00"))
		weights := []int64{1, 1, 1, 1, 1, 1, 1}

		result, err := strategy.Allocate(total, weights, RoundHalfEven)
		if err != nil {
			t.Fatalf("Allocate() error: %v", err)
		}

		// Each part should be $0.14 (truncated)
		for i, part := range result.Parts {
			if !part.Decimal().Equal(MustNewDecimal("0.14")) {
				t.Errorf("Part[%d] = %v, want $0.14", i, part.Decimal())
			}
		}

		// Dust should be $0.02
		if !result.Dust.Quantity().Equal(MustNewDecimal("0.02")) {
			t.Errorf("Dust = %v, want $0.02", result.Dust.Quantity())
		}

		// Verify invariant
		target := total.Quantize(RoundHalfEven).Value
		sumMultipliers := int64(0)
		for _, part := range result.Parts {
			sumMultipliers += part.Multiplier()
		}
		dustQuantized := result.Dust.Quantize(RoundHalfEven).Value
		totalDistributed := sumMultipliers + dustQuantized.Multiplier()

		if totalDistributed != target.Multiplier() {
			t.Errorf("sum(parts) + dust = %d, want %d", totalDistributed, target.Multiplier())
		}
	})

	t.Run("with some zero weights", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		strategy := ProRataTruncateStrategy{}

		// Allocate $10.00 with weights [5, 0, 5]
		total := NewMeasureFrom(spec, MustNewDecimal("10.00"))
		weights := []int64{5, 0, 5}

		result, err := strategy.Allocate(total, weights, RoundHalfEven)
		if err != nil {
			t.Fatalf("Allocate() error: %v", err)
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
		strategy := ProRataTruncateStrategy{}
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
}
