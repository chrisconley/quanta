package quanta

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewScale(t *testing.T) {
	tokens := MustNewUnit("tokens", "1")
	usd := MustNewUnit("USD", "0.01")
	factor := MustNewDecimal("0.05")

	t.Run("creates scale from valid units and factor", func(t *testing.T) {
		s, err := NewScale(tokens, usd, factor)
		require.NoError(t, err)
		assert.True(t, s.InputUnit().Equal(tokens))
		assert.True(t, s.OutputUnit().Equal(usd))
		assert.True(t, s.Factor().Equal(factor))
	})

	t.Run("accepts zero factor", func(t *testing.T) {
		s, err := NewScale(tokens, usd, Zero())
		require.NoError(t, err)
		assert.True(t, s.Factor().IsZero())
	})

	t.Run("accepts negative factor", func(t *testing.T) {
		s, err := NewScale(tokens, usd, MustNewDecimal("-0.05"))
		require.NoError(t, err)
		assert.True(t, s.Factor().Equal(MustNewDecimal("-0.05")))
	})

	t.Run("rejects zero-valued input unit", func(t *testing.T) {
		_, err := NewScale(Unit{}, usd, factor)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "input unit")
	})

	t.Run("rejects zero-valued output unit", func(t *testing.T) {
		_, err := NewScale(tokens, Unit{}, factor)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "output unit")
	})
}

func TestMustNewScale(t *testing.T) {
	tokens := MustNewUnit("tokens", "1")
	usd := MustNewUnit("USD", "0.01")

	t.Run("returns scale on valid inputs", func(t *testing.T) {
		s := MustNewScale(tokens, usd, MustNewDecimal("0.05"))
		assert.True(t, s.OutputUnit().Equal(usd))
	})

	t.Run("panics on zero-valued unit", func(t *testing.T) {
		assert.Panics(t, func() {
			MustNewScale(Unit{}, usd, MustNewDecimal("0.05"))
		})
	})
}

func TestScale_Apply(t *testing.T) {
	tokens := MustNewUnit("tokens", "1")
	usd := MustNewUnit("USD", "0.01")
	credits := MustNewUnit("credits", "1")

	t.Run("scales a matching-unit measure", func(t *testing.T) {
		s := MustNewScale(tokens, usd, MustNewDecimal("0.05"))
		input := NewMeasureFrom(tokens, MustNewDecimal("1000"))
		out, err := s.Apply(input)
		require.NoError(t, err)
		assert.True(t, out.Unit().Equal(usd))
		assert.True(t, out.Quantity().Equal(MustNewDecimal("50.00")))
	})

	t.Run("accepts input with same code but different quantum", func(t *testing.T) {
		tokensFine := MustNewUnit("tokens", "0.001")
		s := MustNewScale(tokens, usd, MustNewDecimal("0.05"))
		input := NewMeasureFrom(tokensFine, MustNewDecimal("1000"))
		out, err := s.Apply(input)
		require.NoError(t, err)
		assert.True(t, out.Unit().Equal(usd))
	})

	t.Run("rejects input with mismatched unit code", func(t *testing.T) {
		s := MustNewScale(tokens, usd, MustNewDecimal("0.05"))
		wrong := NewMeasureFrom(credits, MustNewDecimal("1000"))
		_, err := s.Apply(wrong)
		require.Error(t, err)
		assert.Contains(t, err.Error(), "scale input mismatch")
	})

	t.Run("preserves precision (no rounding)", func(t *testing.T) {
		s := MustNewScale(tokens, usd, MustNewDecimal("0.033333"))
		input := NewMeasureFrom(tokens, MustNewDecimal("100"))
		out, err := s.Apply(input)
		require.NoError(t, err)
		assert.True(t, out.Quantity().Equal(MustNewDecimal("3.3333")))
	})

	t.Run("zero factor yields zero result", func(t *testing.T) {
		s := MustNewScale(tokens, usd, Zero())
		input := NewMeasureFrom(tokens, MustNewDecimal("100"))
		out, err := s.Apply(input)
		require.NoError(t, err)
		assert.True(t, out.Quantity().IsZero())
		assert.True(t, out.Unit().Equal(usd))
	})

	t.Run("zero input yields zero result", func(t *testing.T) {
		s := MustNewScale(tokens, usd, MustNewDecimal("0.05"))
		input := NewMeasureFrom(tokens, Zero())
		out, err := s.Apply(input)
		require.NoError(t, err)
		assert.True(t, out.Quantity().IsZero())
	})
}

func TestScale_Equal(t *testing.T) {
	tokens := MustNewUnit("tokens", "1")
	usd := MustNewUnit("USD", "0.01")

	t.Run("equal scales", func(t *testing.T) {
		a := MustNewScale(tokens, usd, MustNewDecimal("0.05"))
		b := MustNewScale(tokens, usd, MustNewDecimal("0.05"))
		assert.True(t, a.Equal(b))
	})

	t.Run("different factor", func(t *testing.T) {
		a := MustNewScale(tokens, usd, MustNewDecimal("0.05"))
		b := MustNewScale(tokens, usd, MustNewDecimal("0.10"))
		assert.False(t, a.Equal(b))
	})

	t.Run("different input unit", func(t *testing.T) {
		credits := MustNewUnit("credits", "1")
		a := MustNewScale(tokens, usd, MustNewDecimal("0.05"))
		b := MustNewScale(credits, usd, MustNewDecimal("0.05"))
		assert.False(t, a.Equal(b))
	})

	t.Run("different output unit", func(t *testing.T) {
		eur := MustNewUnit("EUR", "0.01")
		a := MustNewScale(tokens, usd, MustNewDecimal("0.05"))
		b := MustNewScale(tokens, eur, MustNewDecimal("0.05"))
		assert.False(t, a.Equal(b))
	})

	t.Run("different quantum on same code not equal", func(t *testing.T) {
		usdFine := MustNewUnit("USD", "0.001")
		a := MustNewScale(tokens, usd, MustNewDecimal("0.05"))
		b := MustNewScale(tokens, usdFine, MustNewDecimal("0.05"))
		assert.False(t, a.Equal(b))
	})
}
