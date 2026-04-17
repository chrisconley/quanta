package quanta_test

import (
	"fmt"
	"testing"

	"github.com/chrisconley/quanta"
	"github.com/cockroachdb/apd/v3"
)

// BenchmarkDecimalAdd measures the overhead of quanta.Decimal.Add
// vs raw apd with a reusable context. The difference is the cost of
// immutability and per-operation context creation.
func BenchmarkDecimalAdd(b *testing.B) {
	b.Run("quanta.Decimal", func(b *testing.B) {
		a := quanta.MustNewDecimal("123.456789")
		c := quanta.MustNewDecimal("987.654321")
		for i := 0; i < b.N; i++ {
			a.Add(c)
		}
	})

	b.Run("raw_apd", func(b *testing.B) {
		a := &apd.Decimal{}
		a.SetString("123.456789")
		c := &apd.Decimal{}
		c.SetString("987.654321")
		ctx := &apd.Context{
			Precision:   34,
			Rounding:    apd.RoundHalfEven,
			Traps:       apd.InvalidOperation | apd.DivisionByZero | apd.Overflow,
			MaxExponent: apd.MaxExponent,
			MinExponent: apd.MinExponent,
		}
		var result apd.Decimal
		for i := 0; i < b.N; i++ {
			ctx.Add(&result, a, c)
		}
	})
}

// BenchmarkQuantize measures the cost of the Measure→Quantized boundary crossing.
func BenchmarkQuantize(b *testing.B) {
	usd := quanta.MustNewUnit("USD", "0.01")
	m := quanta.NewMeasureFrom(usd, quanta.MustNewDecimal("123.456789"))
	for i := 0; i < b.N; i++ {
		m.Quantize(quanta.RoundHalfEven)
	}
}

// BenchmarkAllocation measures allocation performance at different scales.
// The LargestRemainderStrategy uses insertion sort, so this reveals its scaling characteristics.
func BenchmarkAllocation(b *testing.B) {
	usd := quanta.MustNewUnit("USD", "0.01")
	total := quanta.NewMeasureFrom(usd, quanta.MustNewDecimal("1000.00"))
	strategy := quanta.LargestRemainderStrategy{}

	for _, n := range []int{3, 10, 100, 1000} {
		weights := make([]int64, n)
		for i := range weights {
			weights[i] = 1
		}
		b.Run(fmt.Sprintf("LargestRemainder/%d_parts", n), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				strategy.Allocate(total, weights, quanta.RoundHalfEven)
			}
		})
	}
}
