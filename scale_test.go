package quanta

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewScale(t *testing.T) {
	tokens := UnitSpec{Code: "tokens", Quantum: "1"}
	usd := UnitSpec{Code: "USD", Quantum: "0.01"}

	t.Run("creates scale from valid spec", func(t *testing.T) {
		s, err := NewScale(ScaleSpec{Input: tokens, Output: usd, Factor: "0.05"})
		require.NoError(t, err)
		assert.Equal(t, "tokens", s.InputUnit().Code())
		assert.Equal(t, "USD", s.OutputUnit().Code())
		assert.True(t, s.Factor().Equal(MustNewDecimal("0.05")))
	})

	t.Run("accepts zero factor", func(t *testing.T) {
		s, err := NewScale(ScaleSpec{Input: tokens, Output: usd, Factor: "0"})
		require.NoError(t, err)
		assert.True(t, s.Factor().IsZero())
	})

	t.Run("accepts negative factor", func(t *testing.T) {
		s, err := NewScale(ScaleSpec{Input: tokens, Output: usd, Factor: "-0.05"})
		require.NoError(t, err)
		assert.True(t, s.Factor().Equal(MustNewDecimal("-0.05")))
	})

	t.Run("rejects empty input unit code", func(t *testing.T) {
		_, err := NewScale(ScaleSpec{Input: UnitSpec{Quantum: "1"}, Output: usd, Factor: "0.05"})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid input unit")
	})

	t.Run("rejects invalid input quantum", func(t *testing.T) {
		_, err := NewScale(ScaleSpec{Input: UnitSpec{Code: "tokens", Quantum: "bad"}, Output: usd, Factor: "0.05"})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid input unit")
	})

	t.Run("rejects empty output unit code", func(t *testing.T) {
		_, err := NewScale(ScaleSpec{Input: tokens, Output: UnitSpec{Quantum: "0.01"}, Factor: "0.05"})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid output unit")
	})

	t.Run("rejects invalid output quantum", func(t *testing.T) {
		_, err := NewScale(ScaleSpec{Input: tokens, Output: UnitSpec{Code: "USD", Quantum: "bad"}, Factor: "0.05"})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid output unit")
	})

	t.Run("rejects invalid factor", func(t *testing.T) {
		_, err := NewScale(ScaleSpec{Input: tokens, Output: usd, Factor: "not-a-number"})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid factor")
	})

	t.Run("rejects empty factor", func(t *testing.T) {
		_, err := NewScale(ScaleSpec{Input: tokens, Output: usd, Factor: ""})
		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid factor")
	})
}

func TestMustNewScale(t *testing.T) {
	tokens := UnitSpec{Code: "tokens", Quantum: "1"}
	usd := UnitSpec{Code: "USD", Quantum: "0.01"}

	t.Run("returns scale on valid spec", func(t *testing.T) {
		s := MustNewScale(ScaleSpec{Input: tokens, Output: usd, Factor: "0.05"})
		assert.Equal(t, "USD", s.OutputUnit().Code())
	})

	t.Run("panics on empty input code", func(t *testing.T) {
		assert.Panics(t, func() {
			MustNewScale(ScaleSpec{Input: UnitSpec{Quantum: "1"}, Output: usd, Factor: "0.05"})
		})
	})

	t.Run("panics on invalid factor", func(t *testing.T) {
		assert.Panics(t, func() {
			MustNewScale(ScaleSpec{Input: tokens, Output: usd, Factor: "bad"})
		})
	})
}

func TestNewScaleFrom(t *testing.T) {
	tokens := MustNewUnit("tokens", "1")
	usd := MustNewUnit("USD", "0.01")

	t.Run("constructs without validation", func(t *testing.T) {
		s := NewScaleFrom(tokens, usd, MustNewDecimal("0.05"))
		assert.True(t, s.InputUnit().Equal(tokens))
		assert.True(t, s.OutputUnit().Equal(usd))
		assert.True(t, s.Factor().Equal(MustNewDecimal("0.05")))
	})
}

func TestScale_Apply(t *testing.T) {
	tokens := MustNewUnit("tokens", "1")
	usd := MustNewUnit("USD", "0.01")
	credits := MustNewUnit("credits", "1")

	t.Run("scales a matching-unit measure", func(t *testing.T) {
		s := NewScaleFrom(tokens, usd, MustNewDecimal("0.05"))
		input := NewMeasureFrom(tokens, MustNewDecimal("1000"))
		out, err := s.Apply(input)
		require.NoError(t, err)
		assert.True(t, out.Unit().Equal(usd))
		assert.True(t, out.Quantity().Equal(MustNewDecimal("50.00")))
	})

	t.Run("accepts input with same code but different quantum", func(t *testing.T) {
		tokensFine := MustNewUnit("tokens", "0.001")
		s := NewScaleFrom(tokens, usd, MustNewDecimal("0.05"))
		input := NewMeasureFrom(tokensFine, MustNewDecimal("1000"))
		out, err := s.Apply(input)
		require.NoError(t, err)
		assert.True(t, out.Unit().Equal(usd))
	})

	t.Run("rejects input with mismatched unit code", func(t *testing.T) {
		s := NewScaleFrom(tokens, usd, MustNewDecimal("0.05"))
		wrong := NewMeasureFrom(credits, MustNewDecimal("1000"))
		_, err := s.Apply(wrong)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "scale input mismatch")
	})

	t.Run("preserves precision (no rounding)", func(t *testing.T) {
		s := NewScaleFrom(tokens, usd, MustNewDecimal("0.033333"))
		input := NewMeasureFrom(tokens, MustNewDecimal("100"))
		out, err := s.Apply(input)
		require.NoError(t, err)
		assert.True(t, out.Quantity().Equal(MustNewDecimal("3.3333")))
	})

	t.Run("zero factor yields zero result", func(t *testing.T) {
		s := NewScaleFrom(tokens, usd, Zero())
		input := NewMeasureFrom(tokens, MustNewDecimal("100"))
		out, err := s.Apply(input)
		require.NoError(t, err)
		assert.True(t, out.Quantity().IsZero())
		assert.True(t, out.Unit().Equal(usd))
	})

	t.Run("zero input yields zero result", func(t *testing.T) {
		s := NewScaleFrom(tokens, usd, MustNewDecimal("0.05"))
		input := NewMeasureFrom(tokens, Zero())
		out, err := s.Apply(input)
		require.NoError(t, err)
		assert.True(t, out.Quantity().IsZero())
	})
}

func TestScale_Equality(t *testing.T) {
	t.Run("non-comparable", func(t *testing.T) {
		assert.False(t, reflect.TypeOf(Scale{}).Comparable())
	})

	tokens := MustNewUnit("tokens", "1")
	usd := MustNewUnit("USD", "0.01")

	t.Run("equal scales", func(t *testing.T) {
		a := NewScaleFrom(tokens, usd, MustNewDecimal("0.05"))
		b := NewScaleFrom(tokens, usd, MustNewDecimal("0.05"))
		assert.True(t, a.Equal(b))
	})

	t.Run("different factor", func(t *testing.T) {
		a := NewScaleFrom(tokens, usd, MustNewDecimal("0.05"))
		b := NewScaleFrom(tokens, usd, MustNewDecimal("0.10"))
		assert.False(t, a.Equal(b))
	})

	t.Run("different input unit", func(t *testing.T) {
		credits := MustNewUnit("credits", "1")
		a := NewScaleFrom(tokens, usd, MustNewDecimal("0.05"))
		b := NewScaleFrom(credits, usd, MustNewDecimal("0.05"))
		assert.False(t, a.Equal(b))
	})

	t.Run("different output unit", func(t *testing.T) {
		eur := MustNewUnit("EUR", "0.01")
		a := NewScaleFrom(tokens, usd, MustNewDecimal("0.05"))
		b := NewScaleFrom(tokens, eur, MustNewDecimal("0.05"))
		assert.False(t, a.Equal(b))
	})

	t.Run("different quantum on same code not equal", func(t *testing.T) {
		usdFine := MustNewUnit("USD", "0.001")
		a := NewScaleFrom(tokens, usd, MustNewDecimal("0.05"))
		b := NewScaleFrom(tokens, usdFine, MustNewDecimal("0.05"))
		assert.False(t, a.Equal(b))
	})
}
