package quanta

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMeasure_Quantize(t *testing.T) {
	t.Run("quantizes to nearest with RoundHalfEven", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		measure := NewMeasureFrom(spec, MustNewDecimal("12.3456"))

		quantized := measure.Quantize(RoundHalfEven).Value

		assert.Equal(t, int64(1235), quantized.Multiplier())
		assert.True(t, quantized.Decimal().Equal(MustNewDecimal("12.35")))
	})

	t.Run("quantizes with RoundHalfUp", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		measure := NewMeasureFrom(spec, MustNewDecimal("12.345"))

		quantized := measure.Quantize(RoundHalfUp).Value

		assert.Equal(t, int64(1235), quantized.Multiplier())
		assert.True(t, quantized.Decimal().Equal(MustNewDecimal("12.35")))
	})

	t.Run("quantizes with RoundHalfDown", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		measure := NewMeasureFrom(spec, MustNewDecimal("12.345"))

		quantized := measure.Quantize(RoundHalfDown).Value

		assert.Equal(t, int64(1234), quantized.Multiplier())
		assert.True(t, quantized.Decimal().Equal(MustNewDecimal("12.34")))
	})

	t.Run("quantizes with RoundUp", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		measure := NewMeasureFrom(spec, MustNewDecimal("12.341"))

		quantized := measure.Quantize(RoundUp).Value

		assert.Equal(t, int64(1235), quantized.Multiplier())
		assert.True(t, quantized.Decimal().Equal(MustNewDecimal("12.35")))
	})

	t.Run("quantizes with RoundDown", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		measure := NewMeasureFrom(spec, MustNewDecimal("12.349"))

		quantized := measure.Quantize(RoundDown).Value

		assert.Equal(t, int64(1234), quantized.Multiplier())
		assert.True(t, quantized.Decimal().Equal(MustNewDecimal("12.34")))
	})

	t.Run("quantizes with RoundCeiling", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		measure := NewMeasureFrom(spec, MustNewDecimal("12.341"))

		quantized := measure.Quantize(RoundCeiling).Value

		assert.Equal(t, int64(1235), quantized.Multiplier())
		assert.True(t, quantized.Decimal().Equal(MustNewDecimal("12.35")))
	})

	t.Run("quantizes with RoundFloor", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		measure := NewMeasureFrom(spec, MustNewDecimal("12.349"))

		quantized := measure.Quantize(RoundFloor).Value

		assert.Equal(t, int64(1234), quantized.Multiplier())
		assert.True(t, quantized.Decimal().Equal(MustNewDecimal("12.34")))
	})

	t.Run("quantizes value already on quantum boundary", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		measure := NewMeasureFrom(spec, MustNewDecimal("12.35"))

		quantized := measure.Quantize(RoundHalfEven).Value

		assert.Equal(t, int64(1235), quantized.Multiplier())
		assert.True(t, quantized.Decimal().Equal(MustNewDecimal("12.35")))
	})

	t.Run("quantizes negative value", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		measure := NewMeasureFrom(spec, MustNewDecimal("-12.346"))

		quantized := measure.Quantize(RoundHalfEven).Value

		assert.Equal(t, int64(-1235), quantized.Multiplier())
		assert.True(t, quantized.Decimal().Equal(MustNewDecimal("-12.35")))
	})

	t.Run("quantizes with different quantum (0.001)", func(t *testing.T) {
		spec := MustNewUnit("KWD", "0.001")
		measure := NewMeasureFrom(spec, MustNewDecimal("12.3456"))

		quantized := measure.Quantize(RoundHalfEven).Value

		assert.Equal(t, int64(12346), quantized.Multiplier())
		assert.True(t, quantized.Decimal().Equal(MustNewDecimal("12.346")))
	})

	t.Run("quantizes with quantum of 1 (whole units)", func(t *testing.T) {
		spec := MustNewUnit("JPY", "1")
		measure := NewMeasureFrom(spec, MustNewDecimal("123.456"))

		quantized := measure.Quantize(RoundHalfEven).Value

		assert.Equal(t, int64(123), quantized.Multiplier())
		assert.True(t, quantized.Decimal().Equal(MustNewDecimal("123")))
	})

	t.Run("quantizes zero", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		measure := NewMeasureFrom(spec, Zero())

		quantized := measure.Quantize(RoundHalfEven).Value

		assert.Equal(t, int64(0), quantized.Multiplier())
		assert.True(t, quantized.Decimal().Equal(Zero()))
	})
}

func TestNewQuantizedFromMultiplier(t *testing.T) {
	t.Run("creates quantized from multiplier", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		quantized := DeprecatedNewQuantizedFromMultiplier(spec, 1235)

		assert.Equal(t, int64(1235), quantized.Multiplier())
		assert.True(t, quantized.Decimal().Equal(MustNewDecimal("12.35")))
	})

	t.Run("creates quantized with zero multiplier", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		quantized := DeprecatedNewQuantizedFromMultiplier(spec, 0)

		assert.Equal(t, int64(0), quantized.Multiplier())
		assert.True(t, quantized.Decimal().Equal(Zero()))
	})

	t.Run("creates quantized with negative multiplier", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		quantized := DeprecatedNewQuantizedFromMultiplier(spec, -1235)

		assert.Equal(t, int64(-1235), quantized.Multiplier())
		assert.True(t, quantized.Decimal().Equal(MustNewDecimal("-12.35")))
	})

	t.Run("creates quantized with different quantum", func(t *testing.T) {
		spec := MustNewUnit("KWD", "0.001")
		quantized := DeprecatedNewQuantizedFromMultiplier(spec, 12345)

		assert.Equal(t, int64(12345), quantized.Multiplier())
		assert.True(t, quantized.Decimal().Equal(MustNewDecimal("12.345")))
	})
}

func TestQuantized_Spec(t *testing.T) {
	t.Run("returns quantized spec", func(t *testing.T) {
		spec := MustNewUnit("EUR", "0.01")
		quantized := DeprecatedNewQuantizedFromMultiplier(spec, 100)
		assert.True(t, quantized.Unit().Equal(spec))
	})
}

func TestQuantized_Multiplier(t *testing.T) {
	t.Run("returns quantized multiplier", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		quantized := DeprecatedNewQuantizedFromMultiplier(spec, 9999)
		assert.Equal(t, int64(9999), quantized.Multiplier())
	})
}

func TestQuantized_Decimal(t *testing.T) {
	t.Run("converts to decimal correctly", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		quantized := DeprecatedNewQuantizedFromMultiplier(spec, 1235)
		expected := MustNewDecimal("12.35")
		assert.True(t, quantized.Decimal().Equal(expected))
	})

	t.Run("converts with different quantum", func(t *testing.T) {
		spec := MustNewUnit("KWD", "0.001")
		quantized := DeprecatedNewQuantizedFromMultiplier(spec, 12345)
		expected := MustNewDecimal("12.345")
		assert.True(t, quantized.Decimal().Equal(expected))
	})

	t.Run("converts zero", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		quantized := DeprecatedNewQuantizedFromMultiplier(spec, 0)
		assert.True(t, quantized.Decimal().Equal(Zero()))
	})
}

func TestQuantized_Equality(t *testing.T) {
	t.Run("non-comparable", func(t *testing.T) {
		assert.False(t, reflect.TypeOf(Quantized{}).Comparable())
	})
	t.Run("equal quantized values with same spec and multiplier", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		q1 := DeprecatedNewQuantizedFromMultiplier(spec, 1235)
		q2 := DeprecatedNewQuantizedFromMultiplier(spec, 1235)
		assert.True(t, q1.Equal(q2))
	})

	t.Run("not equal with different multipliers", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		q1 := DeprecatedNewQuantizedFromMultiplier(spec, 1235)
		q2 := DeprecatedNewQuantizedFromMultiplier(spec, 1234)
		assert.False(t, q1.Equal(q2))
	})

	t.Run("not equal with different specs (different code)", func(t *testing.T) {
		usdSpec := MustNewUnit("USD", "0.01")
		eurSpec := MustNewUnit("EUR", "0.01")
		q1 := DeprecatedNewQuantizedFromMultiplier(usdSpec, 1235)
		q2 := DeprecatedNewQuantizedFromMultiplier(eurSpec, 1235)
		assert.False(t, q1.Equal(q2))
	})

	t.Run("not equal with different specs (different quantum)", func(t *testing.T) {
		spec1 := MustNewUnit("USD", "0.01")  // 0.01
		spec2 := MustNewUnit("USD", "0.001") // 0.001
		q1 := DeprecatedNewQuantizedFromMultiplier(spec1, 1235)
		q2 := DeprecatedNewQuantizedFromMultiplier(spec2, 1235)
		assert.False(t, q1.Equal(q2))
	})
}

func TestQuantized_String(t *testing.T) {
	t.Run("returns formatted string representation", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		quantized := DeprecatedNewQuantizedFromMultiplier(spec, 1235)
		assert.Equal(t, "Quantized[USD[q=0.01]: 12.35]", quantized.String())
	})

	t.Run("includes spec and decimal value in string", func(t *testing.T) {
		spec := MustNewUnit("tokens", "0.001")
		quantized := DeprecatedNewQuantizedFromMultiplier(spec, 12345)
		assert.Equal(t, "Quantized[tokens[q=0.001]: 12.345]", quantized.String())
	})
}

func TestQuantized_UsageCases(t *testing.T) {
	t.Run("billing invoice line item", func(t *testing.T) {
		// Calculate total with tax, then quantize for invoice
		usdSpec := MustNewUnit("USD", "0.01")

		basePrice := NewMeasureFrom(usdSpec, MustNewDecimal("19.99"))
		taxRate := MustNewDecimal("0.10")
		tax, _ := basePrice.Mul(taxRate)
		total, _ := basePrice.Add(tax)

		// Quantize for invoice (tax rounds up)
		quantizedTotal := total.Quantize(RoundHalfUp).Value

		assert.Equal(t, int64(2199), quantizedTotal.Multiplier())
		assert.True(t, quantizedTotal.Decimal().Equal(MustNewDecimal("21.99")))
	})

	t.Run("meter reading posted to storage", func(t *testing.T) {
		// Aggregate usage, then quantize for storage
		tokensSpec := MustNewUnit("tokens", "0.001")

		usage1 := NewMeasureFrom(tokensSpec, MustNewDecimal("1.2345"))
		usage2 := NewMeasureFrom(tokensSpec, MustNewDecimal("2.3456"))
		totalUsage, _ := usage1.Add(usage2)

		// Quantize for storage (neutral rounding)
		quantizedUsage := totalUsage.Quantize(RoundHalfEven).Value

		assert.Equal(t, int64(3580), quantizedUsage.Multiplier())
		assert.True(t, quantizedUsage.Decimal().Equal(MustNewDecimal("3.580")))
	})

	t.Run("round-trip: measure -> quantize -> storage -> reconstruct", func(t *testing.T) {
		// Original calculation
		usdSpec := MustNewUnit("USD", "0.01")
		original := NewMeasureFrom(usdSpec, MustNewDecimal("123.456"))
		quantized := original.Quantize(RoundHalfEven).Value

		// Store multiplier in database
		storedMultiplier := quantized.Multiplier()
		assert.Equal(t, int64(12346), storedMultiplier)

		// Reconstruct from database
		reconstructed := DeprecatedNewQuantizedFromMultiplier(usdSpec, storedMultiplier)

		assert.True(t, reconstructed.Equal(quantized))
		assert.True(t, reconstructed.Decimal().Equal(MustNewDecimal("123.46")))
	})

	t.Run("aggregating quantized values by summing multipliers", func(t *testing.T) {
		usdSpec := MustNewUnit("USD", "0.01")

		// Three invoice line items
		q1 := DeprecatedNewQuantizedFromMultiplier(usdSpec, 1999) // $19.99
		q2 := DeprecatedNewQuantizedFromMultiplier(usdSpec, 2500) // $25.00
		q3 := DeprecatedNewQuantizedFromMultiplier(usdSpec, 750)  // $7.50

		// Sum multipliers directly (exact arithmetic)
		totalMultiplier := q1.Multiplier() + q2.Multiplier() + q3.Multiplier()
		total := DeprecatedNewQuantizedFromMultiplier(usdSpec, totalMultiplier)

		assert.Equal(t, int64(5249), total.Multiplier())
		assert.True(t, total.Decimal().Equal(MustNewDecimal("52.49")))
	})
}

func TestQuantized_EdgeCases(t *testing.T) {
	t.Run("large multiplier", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		quantized := DeprecatedNewQuantizedFromMultiplier(spec, 999999999)

		assert.Equal(t, int64(999999999), quantized.Multiplier())
		assert.True(t, quantized.Decimal().Equal(MustNewDecimal("9999999.99")))
	})

	t.Run("negative multiplier", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		quantized := DeprecatedNewQuantizedFromMultiplier(spec, -5000)

		assert.Equal(t, int64(-5000), quantized.Multiplier())
		assert.True(t, quantized.Decimal().Equal(MustNewDecimal("-50.00")))
	})

	t.Run("very small quantum", func(t *testing.T) {
		spec := MustNewUnit("bytes", "0.000001")
		measure := NewMeasureFrom(spec, MustNewDecimal("1.2345678"))
		quantized := measure.Quantize(RoundHalfEven).Value

		assert.Equal(t, int64(1234568), quantized.Multiplier())
		assert.True(t, quantized.Decimal().Equal(MustNewDecimal("1.234568")))
	})
}

func TestQuantizationResult_Remainder(t *testing.T) {
	t.Run("positive remainder when rounded down", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		measure := NewMeasureFrom(spec, MustNewDecimal("12.3426"))
		result := measure.Quantize(RoundHalfEven)

		remainder := result.Remainder
		// 12.3426 -> 12.34 (rounded down), so remainder is positive
		assert.False(t, remainder.Quantity().IsNegative())
		assert.False(t, remainder.Quantity().IsZero())
	})

	t.Run("negative remainder when rounded up", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		measure := NewMeasureFrom(spec, MustNewDecimal("12.346"))
		result := measure.Quantize(RoundUp)

		remainder := result.Remainder
		expected := MustNewDecimal("-0.004") // 12.346 - 12.35 = -0.004
		assert.True(t, remainder.Quantity().Equal(expected))
	})

	t.Run("zero remainder for exact fit", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		measure := NewMeasureFrom(spec, MustNewDecimal("12.35"))
		result := measure.Quantize(RoundHalfEven)

		remainder := result.Remainder
		assert.True(t, remainder.Quantity().IsZero())
	})

	t.Run("remainder with negative value", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		measure := NewMeasureFrom(spec, MustNewDecimal("-12.346"))
		result := measure.Quantize(RoundHalfEven)

		// -12.346 -> -12.35 (round to nearest even)
		// remainder = -12.346 - (-12.35) = +0.004
		remainder := result.Remainder
		expected := MustNewDecimal("0.004")
		assert.True(t, remainder.Quantity().Equal(expected))
	})

	t.Run("remainder varies by rounding mode", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		value := MustNewDecimal("12.345")

		// RoundUp: 12.345 -> 12.35, remainder = -0.005
		resUp := NewMeasureFrom(spec, value).Quantize(RoundUp)
		assert.True(t, resUp.Remainder.Quantity().Equal(MustNewDecimal("-0.005")))

		// RoundDown: 12.345 -> 12.34, remainder = +0.005
		resDown := NewMeasureFrom(spec, value).Quantize(RoundDown)
		assert.True(t, resDown.Remainder.Quantity().Equal(MustNewDecimal("0.005")))

		// RoundHalfEven: 12.345 -> 12.34 (round to even), remainder = +0.005
		resHalfEven := NewMeasureFrom(spec, value).Quantize(RoundHalfEven)
		assert.True(t, resHalfEven.Remainder.Quantity().Equal(MustNewDecimal("0.005")))
	})

	t.Run("remainder preserves dimension", func(t *testing.T) {
		spec := MustNewUnit("tokens", "0.001")
		measure := NewMeasureFrom(spec, MustNewDecimal("12.3456"))
		result := measure.Quantize(RoundHalfEven)

		remainder := result.Remainder
		assert.True(t, remainder.Unit().Equal(spec))
	})

	t.Run("accumulating remainders across line items", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")

		// Three line items with remainders
		r1 := NewMeasureFrom(spec, MustNewDecimal("10.334")).Quantize(RoundHalfEven) // ->10.33, rem=+0.004
		r2 := NewMeasureFrom(spec, MustNewDecimal("20.666")).Quantize(RoundHalfEven) // ->20.67, rem=-0.004
		r3 := NewMeasureFrom(spec, MustNewDecimal("30.001")).Quantize(RoundHalfEven) // ->30.00, rem=+0.001

		// Sum remainders
		sum1, _ := r1.Remainder.Add(r2.Remainder)
		totalRemainder, _ := sum1.Add(r3.Remainder)

		// 0.004 + (-0.004) + 0.001 = 0.001
		expected := MustNewDecimal("0.001")
		assert.True(t, totalRemainder.Quantity().Equal(expected))
	})
}
