package quanta

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewMeasure(t *testing.T) {
	usd := UnitSpec{Code: "USD", Quantum: "0.01"}
	tokens := UnitSpec{Code: "tokens", Quantum: "0.001"}
	jpy := UnitSpec{Code: "JPY", Quantum: "1"}

	t.Run("creates measure from spec", func(t *testing.T) {
		measure, err := NewMeasure(MeasureSpec{Unit: usd, Quantity: "123.45"})
		assert.NoError(t, err)
		assert.Equal(t, "USD", measure.Unit().Code())
		assert.True(t, measure.Unit().Quantum().Equal(NewQuantum001()))
		assert.True(t, measure.Quantity().Equal(MustNewDecimal("123.45")))
	})

	t.Run("creates measure with different dimension and quantum", func(t *testing.T) {
		measure, err := NewMeasure(MeasureSpec{Unit: tokens, Quantity: "1234.5678"})
		assert.NoError(t, err)
		assert.Equal(t, "tokens", measure.Unit().Code())
		assert.True(t, measure.Unit().Quantum().Equal(NewQuantum0001()))
		assert.True(t, measure.Quantity().Equal(MustNewDecimal("1234.5678")))
	})

	t.Run("creates measure with whole-number quantum", func(t *testing.T) {
		measure, err := NewMeasure(MeasureSpec{Unit: jpy, Quantity: "500"})
		assert.NoError(t, err)
		assert.Equal(t, "JPY", measure.Unit().Code())
		assert.True(t, measure.Unit().Quantum().Equal(NewQuantum1()))
		assert.True(t, measure.Quantity().Equal(MustNewDecimal("500")))
	})

	t.Run("accepts negative quantity", func(t *testing.T) {
		measure, err := NewMeasure(MeasureSpec{Unit: usd, Quantity: "-50.00"})
		assert.NoError(t, err)
		assert.True(t, measure.Quantity().Equal(MustNewDecimal("-50.00")))
	})

	t.Run("accepts zero quantity", func(t *testing.T) {
		measure, err := NewMeasure(MeasureSpec{Unit: usd, Quantity: "0"})
		assert.NoError(t, err)
		assert.True(t, measure.Quantity().Equal(Zero()))
	})

	t.Run("rejects empty unit code", func(t *testing.T) {
		_, err := NewMeasure(MeasureSpec{Unit: UnitSpec{Quantum: "0.01"}, Quantity: "100"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid unit")
	})

	t.Run("rejects invalid quantum", func(t *testing.T) {
		_, err := NewMeasure(MeasureSpec{Unit: UnitSpec{Code: "USD", Quantum: "bad"}, Quantity: "100"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid unit")
	})

	t.Run("rejects zero quantum", func(t *testing.T) {
		_, err := NewMeasure(MeasureSpec{Unit: UnitSpec{Code: "USD", Quantum: "0"}, Quantity: "100"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid unit")
	})

	t.Run("rejects invalid quantity", func(t *testing.T) {
		_, err := NewMeasure(MeasureSpec{Unit: usd, Quantity: "not-a-number"})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid quantity")
	})

	t.Run("rejects empty quantity", func(t *testing.T) {
		_, err := NewMeasure(MeasureSpec{Unit: usd, Quantity: ""})
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid quantity")
	})
}

func TestMustNewMeasure(t *testing.T) {
	usd := UnitSpec{Code: "USD", Quantum: "0.01"}

	t.Run("creates measure without error", func(t *testing.T) {
		measure := MustNewMeasure(MeasureSpec{Unit: usd, Quantity: "123.45"})
		assert.Equal(t, "USD", measure.Unit().Code())
		assert.True(t, measure.Quantity().Equal(MustNewDecimal("123.45")))
	})

	t.Run("panics on empty unit", func(t *testing.T) {
		assert.Panics(t, func() {
			MustNewMeasure(MeasureSpec{Unit: UnitSpec{Quantum: "0.01"}, Quantity: "100"})
		})
	})

	t.Run("panics on invalid quantum", func(t *testing.T) {
		assert.Panics(t, func() {
			MustNewMeasure(MeasureSpec{Unit: UnitSpec{Code: "USD", Quantum: "bad"}, Quantity: "100"})
		})
	})

	t.Run("panics on invalid quantity", func(t *testing.T) {
		assert.Panics(t, func() {
			MustNewMeasure(MeasureSpec{Unit: usd, Quantity: "bad"})
		})
	})

	t.Run("panic value is an error", func(t *testing.T) {
		defer func() {
			r := recover()
			assert.NotNil(t, r)
			_, ok := r.(error)
			assert.True(t, ok, "expected panic value to be an error, got %T", r)
		}()
		MustNewMeasure(MeasureSpec{Unit: UnitSpec{Quantum: "0.01"}, Quantity: "100"})
	})
}

func TestNewMeasureFrom(t *testing.T) {
	t.Run("creates measure with spec and value", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		value := MustNewDecimal("123.45")
		measure := NewMeasureFrom(spec, value)

		assert.Equal(t, "USD", measure.Unit().Code())
		assert.True(t, measure.Quantity().Equal(value))
	})

	t.Run("creates measure with different dimension", func(t *testing.T) {
		spec := MustNewUnit("tokens", "0.001")
		value := MustNewDecimal("1234.5678")
		measure := NewMeasureFrom(spec, value)

		assert.Equal(t, "tokens", measure.Unit().Code())
		assert.True(t, measure.Quantity().Equal(value))
	})
}

func TestMeasure_Spec(t *testing.T) {
	t.Run("returns measure spec", func(t *testing.T) {
		spec := MustNewUnit("EUR", "0.01")
		measure := NewMeasureFrom(spec, MustNewDecimal("100"))
		assert.True(t, measure.Unit().Equal(spec))
	})
}

func TestMeasure_Value(t *testing.T) {
	t.Run("returns measure value", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		value := MustNewDecimal("99.99")
		measure := NewMeasureFrom(spec, value)
		assert.True(t, measure.Quantity().Equal(value))
	})
}

func TestMeasure_Add(t *testing.T) {
	t.Run("adds two measures with same spec", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		m1 := NewMeasureFrom(spec, MustNewDecimal("10.50"))
		m2 := NewMeasureFrom(spec, MustNewDecimal("5.25"))

		result, err := m1.Add(m2)
		assert.NoError(t, err)
		assert.True(t, result.Quantity().Equal(MustNewDecimal("15.75")))
		assert.True(t, result.Unit().Equal(spec))
	})

	t.Run("adds measures with compatible specs (same code, different quantum)", func(t *testing.T) {
		spec1 := MustNewUnit("USD", "0.01")  // 0.01
		spec2 := MustNewUnit("USD", "0.001") // 0.001
		m1 := NewMeasureFrom(spec1, MustNewDecimal("10.50"))
		m2 := NewMeasureFrom(spec2, MustNewDecimal("5.25"))

		result, err := m1.Add(m2)
		assert.NoError(t, err)
		assert.True(t, result.Quantity().Equal(MustNewDecimal("15.75")))
		// Result uses first measure's spec
		assert.True(t, result.Unit().Equal(spec1))
	})

	t.Run("rejects adding measures with different codes", func(t *testing.T) {
		usdSpec := MustNewUnit("USD", "0.01")
		eurSpec := MustNewUnit("EUR", "0.01")
		m1 := NewMeasureFrom(usdSpec, MustNewDecimal("10.50"))
		m2 := NewMeasureFrom(eurSpec, MustNewDecimal("5.25"))

		_, err := m1.Add(m2)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "incompatible specs")
		assert.Contains(t, err.Error(), "USD")
		assert.Contains(t, err.Error(), "EUR")
	})

	t.Run("handles high-quanta addition", func(t *testing.T) {
		spec := MustNewUnit("tokens", "0.000001")
		m1 := NewMeasureFrom(spec, MustNewDecimal("1.234567890123"))
		m2 := NewMeasureFrom(spec, MustNewDecimal("2.345678901234"))

		result, err := m1.Add(m2)
		assert.NoError(t, err)
		expected := MustNewDecimal("3.580246791357")
		assert.True(t, result.Quantity().Equal(expected))
	})
}

func TestMeasure_Sub(t *testing.T) {
	t.Run("subtracts two measures with same spec", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		m1 := NewMeasureFrom(spec, MustNewDecimal("10.50"))
		m2 := NewMeasureFrom(spec, MustNewDecimal("5.25"))

		result, err := m1.Sub(m2)
		assert.NoError(t, err)
		assert.True(t, result.Quantity().Equal(MustNewDecimal("5.25")))
		assert.True(t, result.Unit().Equal(spec))
	})

	t.Run("subtracts measures with compatible specs (same code, different quantum)", func(t *testing.T) {
		spec1 := MustNewUnit("USD", "0.01")
		spec2 := MustNewUnit("USD", "0.001")
		m1 := NewMeasureFrom(spec1, MustNewDecimal("10.50"))
		m2 := NewMeasureFrom(spec2, MustNewDecimal("3.25"))

		result, err := m1.Sub(m2)
		assert.NoError(t, err)
		assert.True(t, result.Quantity().Equal(MustNewDecimal("7.25")))
	})

	t.Run("rejects subtracting measures with different codes", func(t *testing.T) {
		usdSpec := MustNewUnit("USD", "0.01")
		eurSpec := MustNewUnit("EUR", "0.01")
		m1 := NewMeasureFrom(usdSpec, MustNewDecimal("10.50"))
		m2 := NewMeasureFrom(eurSpec, MustNewDecimal("5.25"))

		_, err := m1.Sub(m2)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "incompatible specs")
	})

	t.Run("handles negative results", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		m1 := NewMeasureFrom(spec, MustNewDecimal("5.00"))
		m2 := NewMeasureFrom(spec, MustNewDecimal("10.00"))

		result, err := m1.Sub(m2)
		assert.NoError(t, err)
		assert.True(t, result.Quantity().Equal(MustNewDecimal("-5.00")))
	})
}

func TestMeasure_Mul(t *testing.T) {
	t.Run("multiplies measure by scalar", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		measure := NewMeasureFrom(spec, MustNewDecimal("10.00"))
		scalar := MustNewDecimal("2.5")

		result, err := measure.Mul(scalar)
		assert.NoError(t, err)
		assert.True(t, result.Quantity().Equal(MustNewDecimal("25.00")))
		assert.True(t, result.Unit().Equal(spec))
	})

	t.Run("multiplies with high quanta", func(t *testing.T) {
		spec := MustNewUnit("tokens", "0.000001")
		measure := NewMeasureFrom(spec, MustNewDecimal("1.234567"))
		scalar := MustNewDecimal("2.345678")

		result, err := measure.Mul(scalar)
		assert.NoError(t, err)
		expected := MustNewDecimal("2.895896651426")
		assert.True(t, result.Quantity().Equal(expected))
	})

	t.Run("multiplies by zero", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		measure := NewMeasureFrom(spec, MustNewDecimal("100.00"))
		scalar := Zero()

		result, err := measure.Mul(scalar)
		assert.NoError(t, err)
		assert.True(t, result.Quantity().Equal(Zero()))
	})

	t.Run("multiplies by negative scalar", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		measure := NewMeasureFrom(spec, MustNewDecimal("10.00"))
		scalar := MustNewDecimal("-2")

		result, err := measure.Mul(scalar)
		assert.NoError(t, err)
		assert.True(t, result.Quantity().Equal(MustNewDecimal("-20.00")))
	})
}

func TestMeasure_Div(t *testing.T) {
	t.Run("divides measure by scalar", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		measure := NewMeasureFrom(spec, MustNewDecimal("10.00"))
		scalar := MustNewDecimal("2")

		result, err := measure.Div(scalar)
		assert.NoError(t, err)
		assert.True(t, result.Quantity().Equal(MustNewDecimal("5.00")))
		assert.True(t, result.Unit().Equal(spec))
	})

	t.Run("divides with high quanta", func(t *testing.T) {
		spec := MustNewUnit("tokens", "0.000001")
		measure := NewMeasureFrom(spec, MustNewDecimal("10"))
		scalar := MustNewDecimal("3")

		result, err := measure.Div(scalar)
		assert.NoError(t, err)
		// Division by 3 produces repeating decimal (precision is 34 total digits, so 33 after decimal)
		expected := MustNewDecimal("3.333333333333333333333333333333333")
		assert.True(t, result.Quantity().Equal(expected))
	})

	t.Run("rejects division by zero", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		measure := NewMeasureFrom(spec, MustNewDecimal("10.00"))
		scalar := Zero()

		_, err := measure.Div(scalar)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "error dividing value")
	})

	t.Run("divides by negative scalar", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		measure := NewMeasureFrom(spec, MustNewDecimal("10.00"))
		scalar := MustNewDecimal("-2")

		result, err := measure.Div(scalar)
		assert.NoError(t, err)
		assert.True(t, result.Quantity().Equal(MustNewDecimal("-5.00")))
	})
}

func TestMeasure_String(t *testing.T) {
	t.Run("returns formatted string representation", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		measure := NewMeasureFrom(spec, MustNewDecimal("123.45"))
		assert.Equal(t, "Measure[USD[q=0.01]: 123.45]", measure.String())
	})

	t.Run("includes spec and value in string", func(t *testing.T) {
		spec := MustNewUnit("tokens", "0.001")
		measure := NewMeasureFrom(spec, MustNewDecimal("1234.567"))
		assert.Equal(t, "Measure[tokens[q=0.001]: 1234.567]", measure.String())
	})
}

func TestMeasure_UsageCases(t *testing.T) {
	t.Run("billing calculation pipeline", func(t *testing.T) {
		// Simulate a billing calculation: usage x rate + tax
		usdSpec := MustNewUnit("USD", "0.01")

		// Base price
		basePrice := NewMeasureFrom(usdSpec, MustNewDecimal("19.99"))

		// Add tax (10%)
		taxRate := MustNewDecimal("0.10")
		tax, err := basePrice.Mul(taxRate)
		assert.NoError(t, err)

		// Total
		total, err := basePrice.Add(tax)
		assert.NoError(t, err)

		expected := MustNewDecimal("21.989") // 19.99 + 1.999
		assert.True(t, total.Quantity().Equal(expected))

		// Note: In real usage, you would call total.Quantize(RoundHalfUp)
		// to snap to 0.01 quantum -> $21.99 for the invoice
		// (Quantize will be added in next commit)
	})

	t.Run("usage aggregation pipeline", func(t *testing.T) {
		// Simulate aggregating token usage across multiple operations
		tokensSpec := MustNewUnit("tokens", "0.001")

		operation1 := NewMeasureFrom(tokensSpec, MustNewDecimal("1.234"))
		operation2 := NewMeasureFrom(tokensSpec, MustNewDecimal("2.345"))
		operation3 := NewMeasureFrom(tokensSpec, MustNewDecimal("3.456"))

		sum1, err := operation1.Add(operation2)
		assert.NoError(t, err)
		total, err := sum1.Add(operation3)
		assert.NoError(t, err)

		expected := MustNewDecimal("7.035")
		assert.True(t, total.Quantity().Equal(expected))
	})

	t.Run("rate calculation: usage x unit price", func(t *testing.T) {
		// Usage: 1,234.5678 tokens
		tokensSpec := MustNewUnit("tokens", "0.0001")
		usage := NewMeasureFrom(tokensSpec, MustNewDecimal("1234.5678"))

		// Rate: $0.0025 per token
		// Note: We can't multiply Measure x Measure
		// So we use scalar multiplication for the rate
		rate := MustNewDecimal("0.0025")
		cost, err := usage.Mul(rate)
		assert.NoError(t, err)

		expected := MustNewDecimal("3.08641950")
		assert.True(t, cost.Quantity().Equal(expected))
	})
}

func TestMeasure_Equality(t *testing.T) {
	t.Run("non-comparable", func(t *testing.T) {
		assert.False(t, reflect.TypeOf(Measure{}).Comparable())
	})

	usd := MustNewUnit("USD", "0.01")
	eur := MustNewUnit("EUR", "0.01")
	usdFine := MustNewUnit("USD", "0.001")

	t.Run("same quantity and unit are equal", func(t *testing.T) {
		a := NewMeasureFrom(usd, MustNewDecimal("100.00"))
		b := NewMeasureFrom(usd, MustNewDecimal("100.00"))
		assert.True(t, a.Equal(b))
	})

	t.Run("different decimal representation of same quantity is equal", func(t *testing.T) {
		a := NewMeasureFrom(usd, MustNewDecimal("1.0"))
		b := NewMeasureFrom(usd, MustNewDecimal("1.00"))
		assert.True(t, a.Equal(b))
	})

	t.Run("different quantity is not equal", func(t *testing.T) {
		a := NewMeasureFrom(usd, MustNewDecimal("100.00"))
		b := NewMeasureFrom(usd, MustNewDecimal("200.00"))
		assert.False(t, a.Equal(b))
	})

	t.Run("different unit code is not equal", func(t *testing.T) {
		a := NewMeasureFrom(usd, MustNewDecimal("100.00"))
		b := NewMeasureFrom(eur, MustNewDecimal("100.00"))
		assert.False(t, a.Equal(b))
	})

	t.Run("different quantum is not equal", func(t *testing.T) {
		a := NewMeasureFrom(usd, MustNewDecimal("100.00"))
		b := NewMeasureFrom(usdFine, MustNewDecimal("100.00"))
		assert.False(t, a.Equal(b))
	})
}

// TestMeasure_TrapsPropagate is a smoke test that Decimal-layer trap errors
// surface through Measure's wrapper methods with both the quanta-level and
// apd-level context intact. Trap behavior itself is exhaustively characterized
// at the Decimal layer in TestDecimal_Traps — this test defends the
// error-propagation contract, not the trap set.
func TestMeasure_TrapsPropagate(t *testing.T) {
	usd := MustNewUnit("USD", "0.01")
	m := NewMeasureFrom(usd, MustNewDecimal("100.00"))

	_, err := m.Div(Zero())

	require.Error(t, err)
	assert.ErrorContains(t, err, "error dividing value",
		"Measure should wrap Decimal errors with its own context")
	assert.ErrorContains(t, err, "division by zero",
		"the underlying apd trap condition should remain visible in the chain")
}
