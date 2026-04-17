package quanta_test

import (
	"testing"

	"github.com/chrisconley/quanta"
)

// FuzzAllocationSumInvariant verifies that LargestRemainderStrategy always
// produces parts that sum to exactly the quantized total, regardless of input.
func FuzzAllocationSumInvariant(f *testing.F) {
	f.Add(int64(1000), int64(1), int64(1), int64(1))    // $10.00 equal split
	f.Add(int64(3), int64(1), int64(1), int64(1))       // $0.03 penny split
	f.Add(int64(-1000), int64(5), int64(3), int64(2))   // refund
	f.Add(int64(1), int64(1), int64(1000000), int64(1)) // extreme weight imbalance
	f.Add(int64(0), int64(1), int64(1), int64(1))       // zero total

	f.Fuzz(func(t *testing.T, totalCents, w1, w2, w3 int64) {
		if w1 <= 0 || w2 <= 0 || w3 <= 0 {
			t.Skip()
		}
		if totalCents > 1e12 || totalCents < -1e12 {
			t.Skip()
		}

		usd := quanta.MustNewUnit("USD", "0.01")
		quantum := quanta.MustNewDecimal("0.01")
		centsDec := quanta.NewDecimalFromInt64(totalCents)
		totalValue, err := centsDec.Mul(quantum)
		if err != nil {
			t.Skip()
		}

		total := quanta.NewMeasureFrom(usd, totalValue)
		strategy := quanta.LargestRemainderStrategy{}
		result, err := strategy.Allocate(total, []int64{w1, w2, w3}, quanta.RoundHalfEven)
		if err != nil {
			t.Skip()
		}

		// Invariant: sum(parts) == quantized(total)
		target := total.Quantize(quanta.RoundHalfEven).Value.Multiplier()
		sum := int64(0)
		for _, part := range result.Parts {
			sum += part.Multiplier()
		}
		if sum != target {
			t.Errorf("sum(parts)=%d != target=%d for totalCents=%d weights=[%d,%d,%d]",
				sum, target, totalCents, w1, w2, w3)
		}
	})
}

// FuzzQuantizeRoundtrip verifies two invariants for any value and quantum:
//  1. quantized + remainder == original (nothing lost)
//  2. |remainder| < quantum (rounding moved less than one quantum)
func FuzzQuantizeRoundtrip(f *testing.F) {
	f.Add("19.99", "0.01")
	f.Add("0.001", "0.001")
	f.Add("-5.55", "0.01")
	f.Add("123456.789", "0.01")
	f.Add("0", "0.01")
	f.Add("1", "1")

	f.Fuzz(func(t *testing.T, valueStr, quantumStr string) {
		unit, err := quanta.NewUnit("X", quantumStr)
		if err != nil {
			t.Skip()
		}
		valueDec, err := quanta.NewDecimal(valueStr)
		if err != nil {
			t.Skip()
		}

		m := quanta.NewMeasureFrom(unit, valueDec)
		result := m.Quantize(quanta.RoundHalfEven)

		// Invariant 1: quantized + remainder == original
		reconstructed, err := result.Value.Decimal().Add(result.Remainder.Quantity())
		if err != nil {
			t.Fatalf("reconstruction add failed: %v", err)
		}
		if !reconstructed.Equal(valueDec) {
			t.Errorf("roundtrip failed: %s + %s = %s, want %s",
				result.Value.Decimal(), result.Remainder.Quantity(), reconstructed, valueDec)
		}

		// Invariant 2: |remainder| < quantum
		quantumDec := unit.Quantum().Decimal()
		remainderAbs := result.Remainder.Quantity().Abs()
		if remainderAbs.Cmp(quantumDec) >= 0 {
			t.Errorf("|remainder| >= quantum: |%s| = %s >= %s",
				result.Remainder.Quantity(), remainderAbs, quantumDec)
		}
	})
}
