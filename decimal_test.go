package quanta

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewDecimal(t *testing.T) {
	t.Run("with valid integer string creates Decimal", func(t *testing.T) {
		d, err := NewDecimal("100")

		require.NoError(t, err)
		assert.Equal(t, "100", fmt.Sprintf("%s", d))
	})

	t.Run("with valid decimal string creates Decimal", func(t *testing.T) {
		d, err := NewDecimal("123.456")

		require.NoError(t, err)
		assert.Equal(t, "123.456", fmt.Sprintf("%s", d))
	})

	t.Run("with negative decimal creates Decimal", func(t *testing.T) {
		d, err := NewDecimal("-50.5")

		require.NoError(t, err)
		assert.Equal(t, "-50.5", fmt.Sprintf("%s", d))
		assert.True(t, d.IsNegative())
	})

	t.Run("with zero creates Decimal", func(t *testing.T) {
		d, err := NewDecimal("0")

		require.NoError(t, err)
		assert.Equal(t, "0", fmt.Sprintf("%s", d))
		assert.True(t, d.IsZero())
	})

	t.Run("with whitespace trims and creates Decimal", func(t *testing.T) {
		d, err := NewDecimal("  100.50  ")

		require.NoError(t, err)
		assert.Equal(t, "100.50", fmt.Sprintf("%s", d))
	})

	t.Run("with scientific notation creates Decimal", func(t *testing.T) {
		d, err := NewDecimal("1.5e6")

		require.NoError(t, err)
		assert.Equal(t, "1.5E+6", fmt.Sprintf("%s", d))
	})

	t.Run("with very large number creates Decimal", func(t *testing.T) {
		d, err := NewDecimal("123456789012345678901234567890.123456789")

		require.NoError(t, err)
		assert.Equal(t, "123456789012345678901234567890.123456789", fmt.Sprintf("%s", d))
	})

	t.Run("with invalid string returns error", func(t *testing.T) {
		_, err := NewDecimal("abc")

		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid decimal string")
	})

	t.Run("with empty string returns error", func(t *testing.T) {
		_, err := NewDecimal("")

		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid decimal string")
	})

	t.Run("with multiple decimal points returns error", func(t *testing.T) {
		_, err := NewDecimal("12.34.56")

		require.Error(t, err)
		assert.Contains(t, err.Error(), "invalid decimal string")
	})
}

func TestMustNewDecimal(t *testing.T) {
	t.Run("with valid string creates Decimal", func(t *testing.T) {
		d := MustNewDecimal("123.456")

		assert.Equal(t, "123.456", fmt.Sprintf("%s", d))
	})

	t.Run("with invalid string panics", func(t *testing.T) {
		assert.Panics(t, func() {
			MustNewDecimal("invalid")
		})
	})
}

func TestNewDecimalFromInt64(t *testing.T) {
	t.Run("creates Decimal from positive int64", func(t *testing.T) {
		d := NewDecimalFromInt64(12345)

		assert.Equal(t, "12345", fmt.Sprintf("%s", d))
	})

	t.Run("creates Decimal from negative int64", func(t *testing.T) {
		d := NewDecimalFromInt64(-12345)

		assert.Equal(t, "-12345", fmt.Sprintf("%s", d))
		assert.True(t, d.IsNegative())
	})

	t.Run("creates Decimal from zero", func(t *testing.T) {
		d := NewDecimalFromInt64(0)

		assert.Equal(t, "0", fmt.Sprintf("%s", d))
		assert.True(t, d.IsZero())
	})

	t.Run("creates Decimal from large int64", func(t *testing.T) {
		d := NewDecimalFromInt64(9223372036854775807) // Max int64

		assert.Equal(t, "9223372036854775807", fmt.Sprintf("%s", d))
	})
}

func TestZero(t *testing.T) {
	t.Run("creates zero Decimal", func(t *testing.T) {
		d := Zero()

		assert.Equal(t, "0", fmt.Sprintf("%s", d))
		assert.True(t, d.IsZero())
		assert.False(t, d.IsNegative())
	})
}

func TestDecimal_Key(t *testing.T) {
	t.Run("returns canonical string for simple decimal", func(t *testing.T) {
		d := MustNewDecimal("123.45")

		assert.Equal(t, "123.45", d.Key())
	})

	t.Run("normalizes trailing zeros", func(t *testing.T) {
		d1 := MustNewDecimal("123")
		d2 := MustNewDecimal("123.0")
		d3 := MustNewDecimal("123.00")

		key1 := d1.Key()
		key2 := d2.Key()
		key3 := d3.Key()

		assert.Equal(t, key1, key2, "123 and 123.0 should have same key")
		assert.Equal(t, key2, key3, "123.0 and 123.00 should have same key")
		assert.Equal(t, "123", key1, "normalized key should be 123")
	})

	t.Run("handles zero with different representations", func(t *testing.T) {
		d1 := MustNewDecimal("0")
		d2 := MustNewDecimal("0.0")
		d3 := MustNewDecimal("0.00")

		key1 := d1.Key()
		key2 := d2.Key()
		key3 := d3.Key()

		assert.Equal(t, key1, key2)
		assert.Equal(t, key2, key3)
		assert.Equal(t, "0", key1)
	})

	t.Run("preserves meaningful decimals", func(t *testing.T) {
		d := MustNewDecimal("123.450")

		// Should normalize to 123.45 (remove trailing zero)
		assert.Equal(t, "123.45", d.Key())
	})

	t.Run("handles negative values", func(t *testing.T) {
		d := MustNewDecimal("-123.45")

		assert.Equal(t, "-123.45", d.Key())
	})
}

func TestDecimal_IsZero(t *testing.T) {
	t.Run("zero value returns true", func(t *testing.T) {
		d := MustNewDecimal("0")

		assert.True(t, d.IsZero())
	})

	t.Run("zero with decimals returns true", func(t *testing.T) {
		d := MustNewDecimal("0.00")

		assert.True(t, d.IsZero())
	})

	t.Run("non-zero value returns false", func(t *testing.T) {
		d := MustNewDecimal("0.01")

		assert.False(t, d.IsZero())
	})

	t.Run("negative zero returns true", func(t *testing.T) {
		d := MustNewDecimal("-0")

		assert.True(t, d.IsZero())
	})
}

func TestDecimal_IsNegative(t *testing.T) {
	t.Run("negative value returns true", func(t *testing.T) {
		d := MustNewDecimal("-100")

		assert.True(t, d.IsNegative())
	})

	t.Run("positive value returns false", func(t *testing.T) {
		d := MustNewDecimal("100")

		assert.False(t, d.IsNegative())
	})

	t.Run("zero returns false", func(t *testing.T) {
		d := MustNewDecimal("0")

		assert.False(t, d.IsNegative())
	})
}

func TestDecimal_Equality(t *testing.T) {
	t.Run("non-comparable", func(t *testing.T) {
		assert.False(t, reflect.TypeOf(Decimal{}).Comparable())
	})
	t.Run("equal values return true", func(t *testing.T) {
		d1 := MustNewDecimal("123.45")
		d2 := MustNewDecimal("123.45")

		assert.True(t, d1.Equal(d2))
	})

	t.Run("different values return false", func(t *testing.T) {
		d1 := MustNewDecimal("123.45")
		d2 := MustNewDecimal("123.46")

		assert.False(t, d1.Equal(d2))
	})

	t.Run("different quanta but equal value return true", func(t *testing.T) {
		d1 := MustNewDecimal("123")
		d2 := MustNewDecimal("123.00")

		assert.True(t, d1.Equal(d2))
	})

	t.Run("zero equals negative zero", func(t *testing.T) {
		d1 := MustNewDecimal("0")
		d2 := MustNewDecimal("-0")

		assert.True(t, d1.Equal(d2))
	})
}

func TestDecimal_Cmp(t *testing.T) {
	t.Run("equal values return 0", func(t *testing.T) {
		d1 := MustNewDecimal("100.50")
		d2 := MustNewDecimal("100.50")

		assert.Equal(t, 0, d1.Cmp(d2))
	})

	t.Run("less than returns -1", func(t *testing.T) {
		d1 := MustNewDecimal("50")
		d2 := MustNewDecimal("100")

		assert.Equal(t, -1, d1.Cmp(d2))
	})

	t.Run("greater than returns 1", func(t *testing.T) {
		d1 := MustNewDecimal("100")
		d2 := MustNewDecimal("50")

		assert.Equal(t, 1, d1.Cmp(d2))
	})

	t.Run("different quanta but equal value returns 0", func(t *testing.T) {
		d1 := MustNewDecimal("100.0")
		d2 := MustNewDecimal("100.00")

		assert.Equal(t, 0, d1.Cmp(d2))
	})

	t.Run("negative comparison works correctly", func(t *testing.T) {
		d1 := MustNewDecimal("-100")
		d2 := MustNewDecimal("-50")

		assert.Equal(t, -1, d1.Cmp(d2))
	})
}

func TestDecimal_String(t *testing.T) {
	t.Run("returns string representation", func(t *testing.T) {
		d := MustNewDecimal("123.456")

		assert.Equal(t, "123.456", fmt.Sprintf("%s", d))
	})

	t.Run("preserves quanta", func(t *testing.T) {
		d := MustNewDecimal("100.00")

		assert.Equal(t, "100.00", fmt.Sprintf("%s", d))
	})
}

func TestDecimal_Add(t *testing.T) {
	t.Run("adds two positive decimals", func(t *testing.T) {
		d1 := MustNewDecimal("100.50")
		d2 := MustNewDecimal("50.25")

		result, err := d1.Add(d2)

		require.NoError(t, err)
		assert.True(t, result.Equal(MustNewDecimal("150.75")))
	})

	t.Run("adds positive and negative decimal", func(t *testing.T) {
		d1 := MustNewDecimal("100")
		d2 := MustNewDecimal("-30")

		result, err := d1.Add(d2)

		require.NoError(t, err)
		assert.True(t, result.Equal(MustNewDecimal("70")))
	})

	t.Run("adds to zero", func(t *testing.T) {
		d1 := Zero()
		d2 := MustNewDecimal("42.5")

		result, err := d1.Add(d2)

		require.NoError(t, err)
		assert.True(t, result.Equal(MustNewDecimal("42.5")))
	})

	t.Run("handles large numbers", func(t *testing.T) {
		d1 := MustNewDecimal("999999999999999999")
		d2 := MustNewDecimal("1")

		result, err := d1.Add(d2)

		require.NoError(t, err)
		assert.True(t, result.Equal(MustNewDecimal("1000000000000000000")))
	})

	t.Run("preserves decimal quanta", func(t *testing.T) {
		d1 := MustNewDecimal("0.001")
		d2 := MustNewDecimal("0.002")

		result, err := d1.Add(d2)

		require.NoError(t, err)
		assert.True(t, result.Equal(MustNewDecimal("0.003")))
	})

	t.Run("original decimals unchanged", func(t *testing.T) {
		d1 := MustNewDecimal("100")
		d2 := MustNewDecimal("50")

		_, err := d1.Add(d2)

		require.NoError(t, err)
		assert.Equal(t, "100", fmt.Sprintf("%s", d1))
		assert.Equal(t, "50", fmt.Sprintf("%s", d2))
	})
}

func TestDecimal_Sub(t *testing.T) {
	t.Run("subtracts two positive decimals", func(t *testing.T) {
		d1 := MustNewDecimal("100.50")
		d2 := MustNewDecimal("50.25")

		result, err := d1.Sub(d2)

		require.NoError(t, err)
		assert.True(t, result.Equal(MustNewDecimal("50.25")))
	})

	t.Run("subtracts negative decimal", func(t *testing.T) {
		d1 := MustNewDecimal("100")
		d2 := MustNewDecimal("-30")

		result, err := d1.Sub(d2)

		require.NoError(t, err)
		assert.True(t, result.Equal(MustNewDecimal("130")))
	})

	t.Run("subtracts from zero", func(t *testing.T) {
		d1 := Zero()
		d2 := MustNewDecimal("42.5")

		result, err := d1.Sub(d2)

		require.NoError(t, err)
		assert.True(t, result.Equal(MustNewDecimal("-42.5")))
	})

	t.Run("subtracts to negative result", func(t *testing.T) {
		d1 := MustNewDecimal("50")
		d2 := MustNewDecimal("100")

		result, err := d1.Sub(d2)

		require.NoError(t, err)
		assert.True(t, result.Equal(MustNewDecimal("-50")))
		assert.True(t, result.IsNegative())
	})

	t.Run("original decimals unchanged", func(t *testing.T) {
		d1 := MustNewDecimal("100")
		d2 := MustNewDecimal("50")

		_, err := d1.Sub(d2)

		require.NoError(t, err)
		assert.Equal(t, "100", fmt.Sprintf("%s", d1))
		assert.Equal(t, "50", fmt.Sprintf("%s", d2))
	})
}

func TestDecimal_Mul(t *testing.T) {
	t.Run("multiplies two positive decimals", func(t *testing.T) {
		d1 := MustNewDecimal("10.5")
		d2 := MustNewDecimal("2.0")

		result, err := d1.Mul(d2)

		require.NoError(t, err)
		assert.True(t, result.Equal(MustNewDecimal("21")))
	})

	t.Run("multiplies by zero", func(t *testing.T) {
		d1 := MustNewDecimal("100.50")
		d2 := Zero()

		result, err := d1.Mul(d2)

		require.NoError(t, err)
		assert.True(t, result.IsZero())
	})

	t.Run("multiplies negative by positive", func(t *testing.T) {
		d1 := MustNewDecimal("-10")
		d2 := MustNewDecimal("5")

		result, err := d1.Mul(d2)

		require.NoError(t, err)
		assert.True(t, result.Equal(MustNewDecimal("-50")))
	})

	t.Run("multiplies two negative decimals", func(t *testing.T) {
		d1 := MustNewDecimal("-10")
		d2 := MustNewDecimal("-5")

		result, err := d1.Mul(d2)

		require.NoError(t, err)
		assert.True(t, result.Equal(MustNewDecimal("50")))
	})

	t.Run("preserves high quanta", func(t *testing.T) {
		d1 := MustNewDecimal("1.23456789")
		d2 := MustNewDecimal("9.87654321")

		result, err := d1.Mul(d2)

		require.NoError(t, err)
		// 1.23456789 * 9.87654321 = 12.19326311126352690869
		assert.Contains(t, fmt.Sprintf("%s", result), "12.193263111")
	})

	t.Run("original decimals unchanged", func(t *testing.T) {
		d1 := MustNewDecimal("10")
		d2 := MustNewDecimal("5")

		_, err := d1.Mul(d2)

		require.NoError(t, err)
		assert.Equal(t, "10", fmt.Sprintf("%s", d1))
		assert.Equal(t, "5", fmt.Sprintf("%s", d2))
	})
}

func TestDecimal_Div(t *testing.T) {
	t.Run("divides two positive decimals", func(t *testing.T) {
		d1 := MustNewDecimal("100")
		d2 := MustNewDecimal("4")

		result, err := d1.Div(d2)

		require.NoError(t, err)
		assert.True(t, result.Equal(MustNewDecimal("25")))
	})

	t.Run("divides with decimal result", func(t *testing.T) {
		d1 := MustNewDecimal("10")
		d2 := MustNewDecimal("3")

		result, err := d1.Div(d2)

		require.NoError(t, err)
		// Should have high precision
		assert.Contains(t, fmt.Sprintf("%s", result), "3.333333")
	})

	t.Run("divides by zero returns error", func(t *testing.T) {
		d1 := MustNewDecimal("100")
		d2 := Zero()

		_, err := d1.Div(d2)

		require.Error(t, err)
		assert.Contains(t, err.Error(), "div failed")
	})

	t.Run("divides negative by positive", func(t *testing.T) {
		d1 := MustNewDecimal("-100")
		d2 := MustNewDecimal("4")

		result, err := d1.Div(d2)

		require.NoError(t, err)
		assert.True(t, result.Equal(MustNewDecimal("-25")))
	})

	t.Run("divides zero by non-zero", func(t *testing.T) {
		d1 := Zero()
		d2 := MustNewDecimal("100")

		result, err := d1.Div(d2)

		require.NoError(t, err)
		assert.True(t, result.IsZero())
	})

	t.Run("original decimals unchanged", func(t *testing.T) {
		d1 := MustNewDecimal("100")
		d2 := MustNewDecimal("4")

		_, err := d1.Div(d2)

		require.NoError(t, err)
		assert.Equal(t, "100", fmt.Sprintf("%s", d1))
		assert.Equal(t, "4", fmt.Sprintf("%s", d2))
	})
}

func TestDecimal_Floor(t *testing.T) {
	t.Run("positive decimal rounds down", func(t *testing.T) {
		d := MustNewDecimal("12.8")
		assert.Equal(t, int64(12), d.Floor())
	})

	t.Run("positive decimal already integer", func(t *testing.T) {
		d := MustNewDecimal("12.0")
		assert.Equal(t, int64(12), d.Floor())
	})

	t.Run("negative decimal rounds toward negative infinity", func(t *testing.T) {
		d := MustNewDecimal("-12.2")
		assert.Equal(t, int64(-13), d.Floor())
	})

	t.Run("negative decimal more negative", func(t *testing.T) {
		d := MustNewDecimal("-12.8")
		assert.Equal(t, int64(-13), d.Floor())
	})

	t.Run("zero", func(t *testing.T) {
		d := Zero()
		assert.Equal(t, int64(0), d.Floor())
	})

	t.Run("small positive fraction", func(t *testing.T) {
		d := MustNewDecimal("0.999")
		assert.Equal(t, int64(0), d.Floor())
	})

	t.Run("small negative fraction", func(t *testing.T) {
		d := MustNewDecimal("-0.5")
		assert.Equal(t, int64(-1), d.Floor())
	})
}

func TestDecimal_Ceiling(t *testing.T) {
	t.Run("positive decimal rounds up", func(t *testing.T) {
		d := MustNewDecimal("12.2")
		assert.Equal(t, int64(13), d.Ceiling())
	})

	t.Run("positive decimal already integer", func(t *testing.T) {
		d := MustNewDecimal("12.0")
		assert.Equal(t, int64(12), d.Ceiling())
	})

	t.Run("negative decimal rounds toward positive infinity", func(t *testing.T) {
		d := MustNewDecimal("-12.8")
		assert.Equal(t, int64(-12), d.Ceiling())
	})

	t.Run("negative decimal less negative", func(t *testing.T) {
		d := MustNewDecimal("-12.2")
		assert.Equal(t, int64(-12), d.Ceiling())
	})

	t.Run("zero", func(t *testing.T) {
		d := Zero()
		assert.Equal(t, int64(0), d.Ceiling())
	})

	t.Run("small positive fraction", func(t *testing.T) {
		d := MustNewDecimal("0.5")
		assert.Equal(t, int64(1), d.Ceiling())
	})

	t.Run("small negative fraction", func(t *testing.T) {
		d := MustNewDecimal("-0.5")
		assert.Equal(t, int64(0), d.Ceiling())
	})
}

func TestDecimal_Frac(t *testing.T) {
	t.Run("positive decimal extracts fractional part", func(t *testing.T) {
		d := MustNewDecimal("12.34")
		frac := d.Frac()
		expected := MustNewDecimal("0.34")
		assert.True(t, frac.Equal(expected))
	})

	t.Run("integer has zero fractional part", func(t *testing.T) {
		d := MustNewDecimal("12.00")
		frac := d.Frac()
		assert.True(t, frac.IsZero())
	})

	t.Run("negative decimal fractional part", func(t *testing.T) {
		// -12.34 has floor(-12.34) = -13
		// frac = -12.34 - (-13) = 0.66
		d := MustNewDecimal("-12.34")
		frac := d.Frac()
		expected := MustNewDecimal("0.66")
		assert.True(t, frac.Equal(expected))
	})

	t.Run("zero has zero fractional part", func(t *testing.T) {
		d := Zero()
		frac := d.Frac()
		assert.True(t, frac.IsZero())
	})

	t.Run("fractional part always in [0, 1) for positive", func(t *testing.T) {
		d := MustNewDecimal("123.999")
		frac := d.Frac()

		assert.False(t, frac.IsNegative())
		assert.True(t, frac.Cmp(MustNewDecimal("1")) < 0)
		assert.True(t, frac.Equal(MustNewDecimal("0.999")))
	})

	t.Run("fractional part in (0, 1) for negative", func(t *testing.T) {
		// -123.001 has floor = -124
		// frac = -123.001 - (-124) = 0.999
		d := MustNewDecimal("-123.001")
		frac := d.Frac()

		assert.False(t, frac.IsNegative())
		assert.True(t, frac.Cmp(MustNewDecimal("1")) < 0)
		expected := MustNewDecimal("0.999")
		assert.True(t, frac.Equal(expected))
	})
}

func TestDecimal_Negate(t *testing.T) {
	t.Run("negates positive decimal", func(t *testing.T) {
		d := MustNewDecimal("12.34")
		result := d.Negate()

		assert.True(t, result.Equal(MustNewDecimal("-12.34")))
		assert.True(t, result.IsNegative())
	})

	t.Run("negates negative decimal", func(t *testing.T) {
		d := MustNewDecimal("-12.34")
		result := d.Negate()

		assert.True(t, result.Equal(MustNewDecimal("12.34")))
		assert.False(t, result.IsNegative())
	})

	t.Run("negates zero", func(t *testing.T) {
		d := Zero()
		result := d.Negate()

		assert.True(t, result.IsZero())
	})

	t.Run("negates large positive number", func(t *testing.T) {
		d := MustNewDecimal("999999999999.999")
		result := d.Negate()

		assert.True(t, result.Equal(MustNewDecimal("-999999999999.999")))
		assert.True(t, result.IsNegative())
	})

	t.Run("negates large negative number", func(t *testing.T) {
		d := MustNewDecimal("-999999999999.999")
		result := d.Negate()

		assert.True(t, result.Equal(MustNewDecimal("999999999999.999")))
		assert.False(t, result.IsNegative())
	})

	t.Run("negates small positive fraction", func(t *testing.T) {
		d := MustNewDecimal("0.001")
		result := d.Negate()

		assert.True(t, result.Equal(MustNewDecimal("-0.001")))
		assert.True(t, result.IsNegative())
	})

	t.Run("negates small negative fraction", func(t *testing.T) {
		d := MustNewDecimal("-0.001")
		result := d.Negate()

		assert.True(t, result.Equal(MustNewDecimal("0.001")))
		assert.False(t, result.IsNegative())
	})

	t.Run("double negate returns original value", func(t *testing.T) {
		d := MustNewDecimal("42.5")
		result := d.Negate().Negate()

		assert.True(t, result.Equal(d))
	})

	t.Run("original decimal unchanged", func(t *testing.T) {
		d := MustNewDecimal("100.50")
		original := fmt.Sprintf("%s", d)

		_ = d.Negate()

		assert.Equal(t, original, fmt.Sprintf("%s", d))
	})

	t.Run("equivalent to Zero().Sub()", func(t *testing.T) {
		d := MustNewDecimal("123.45")

		negateResult := d.Negate()
		subResult, err := Zero().Sub(d)

		require.NoError(t, err)
		assert.True(t, negateResult.Equal(subResult))
	})
}

func TestDecimal_FloorCeilingFrac_AllocationScenario(t *testing.T) {
	t.Run("divide $10 among 3 parties - track fractional parts", func(t *testing.T) {
		total := MustNewDecimal("1000") // $10.00 as multiplier
		parties := MustNewDecimal("3")

		// Each share: 1000 / 3 = 333.333...
		shareDecimal, err := total.Div(parties)
		require.NoError(t, err)

		// Integer part (floor)
		intPart := shareDecimal.Floor()
		assert.Equal(t, int64(333), intPart)

		// Fractional part
		fracPart := shareDecimal.Frac()
		// Should be 0.333... (non-zero, between 0 and 1)
		assert.False(t, fracPart.IsZero())
		assert.True(t, fracPart.Cmp(MustNewDecimal("0.3")) > 0)
		assert.True(t, fracPart.Cmp(MustNewDecimal("0.4")) < 0)

		// Dust: 1000 - (333 × 3) = 1
		distributed := intPart * 3
		dust := int64(1000) - distributed
		assert.Equal(t, int64(1), dust)

		// In allocation: distribute this 1 unit to party with largest frac
		// (All have same frac, so pick first one)
	})

	t.Run("weighted allocation - different fractional parts", func(t *testing.T) {
		total := MustNewDecimal("1000") // $10.00

		// Weights: [0.5, 0.3, 0.2]
		share1, _ := total.Mul(MustNewDecimal("0.5")) // 500
		share2, _ := total.Mul(MustNewDecimal("0.3")) // 300
		share3, _ := total.Mul(MustNewDecimal("0.2")) // 200

		floor1 := share1.Floor()
		floor2 := share2.Floor()
		floor3 := share3.Floor()

		assert.Equal(t, int64(500), floor1)
		assert.Equal(t, int64(300), floor2)
		assert.Equal(t, int64(200), floor3)

		frac1 := share1.Frac()
		frac2 := share2.Frac()
		frac3 := share3.Frac()

		// All exact, so zero fracs
		assert.True(t, frac1.IsZero())
		assert.True(t, frac2.IsZero())
		assert.True(t, frac3.IsZero())

		// No dust
		dust := int64(1000) - (floor1 + floor2 + floor3)
		assert.Equal(t, int64(0), dust)
	})

	t.Run("uneven split with dust - sort by fractional part", func(t *testing.T) {
		total := MustNewDecimal("100")

		// Weights: [0.33, 0.33, 0.34]
		share1, _ := total.Mul(MustNewDecimal("0.33")) // 33
		share2, _ := total.Mul(MustNewDecimal("0.33")) // 33
		share3, _ := total.Mul(MustNewDecimal("0.34")) // 34

		floor1 := share1.Floor()
		floor2 := share2.Floor()
		floor3 := share3.Floor()

		frac1 := share1.Frac()
		frac2 := share2.Frac()
		frac3 := share3.Frac()

		// All fracs are zero (exact values)
		assert.Equal(t, int64(33), floor1)
		assert.Equal(t, int64(33), floor2)
		assert.Equal(t, int64(34), floor3)

		assert.True(t, frac1.IsZero())
		assert.True(t, frac2.IsZero())
		assert.True(t, frac3.IsZero())

		// No dust
		dust := int64(100) - (floor1 + floor2 + floor3)
		assert.Equal(t, int64(0), dust)
	})
}

// TestDecimal_Traps pins the behavior of every apd Condition reachable
// through quanta's public API. The current trap set is
//
//	InvalidOperation | DivisionByZero | Overflow
//
// but this test characterizes *every* condition, not just the trapped ones.
// Silent-today subtests will flip to errors if the trap set is widened (e.g.
// adopting apd.DefaultTraps), making any such change a visible, reviewable
// test diff rather than an invisible runtime behavior shift.
//
// Coverage is at the Decimal layer only — Measure, Floor, Ceiling, and
// Quantize route through the same newCalcContext helper, so trap behavior
// is characterized here transitively. See TestMeasure_TrapsPropagate for
// the wrapper error-propagation smoke test.
//
// Condition taxonomy (apd v3):
//
//   - Trapped (errors today):
//     DivisionByZero, Overflow, InvalidOperation
//   - Hard-coded (errors regardless of traps, per apd Condition.GoError):
//     SystemOverflow, SystemUnderflow
//   - Signalled but not trapped (silent today):
//     DivisionUndefined, Inexact, Rounded, (quiet NaN propagation)
//   - Not reachable through quanta's public API (not tested):
//     DivisionImpossible (only via QuoInteger, which quanta does not expose),
//     standalone Subnormal (absorbed into SystemUnderflow at our magnitudes),
//     Clamped (apd uses this internally for IEEE Emax clamping we do not hit).
//
// Note that quanta.NewDecimal accepts the strings "NaN", "sNaN", "Infinity",
// and "-Infinity", so those inputs are part of the reachable surface. Whether
// they *should* be accepted is a separate design question; today they are.
func TestDecimal_Traps(t *testing.T) {
	// --- Errors today (trapped or hard-coded) ---

	t.Run("DivisionByZero errors on x/0", func(t *testing.T) {
		_, err := MustNewDecimal("5").Div(Zero())

		require.Error(t, err)
		assert.ErrorContains(t, err, "div failed",
			"quanta wraps the apd error with operation context")
		assert.ErrorContains(t, err, "division by zero",
			"the DivisionByZero condition should surface verbatim")
	})

	t.Run("Overflow errors when result exponent exceeds MaxExponent", func(t *testing.T) {
		// 1E99999 * 1E99999 = 1E199998; apd.MaxExponent is 100000.
		// Both SystemOverflow (hard-coded) and Overflow (trapped) fire.
		huge := MustNewDecimal("1E99999")

		_, err := huge.Mul(huge)

		require.Error(t, err)
		assert.ErrorContains(t, err, "mul failed")
		assert.ErrorContains(t, err, "exponent out of range",
			"apd uses this text for both System and regular Overflow")
	})

	t.Run("SystemUnderflow errors when result exponent falls below MinExponent", func(t *testing.T) {
		// 1E-99999 * 1E-99999 = 1E-199998; apd.MinExponent is -100000.
		// SystemUnderflow fires and is hard-coded to error even though
		// Underflow is NOT in our trap set. Confirms that "very tiny results
		// error" behavior survives any future trap-set narrowing.
		tiny := MustNewDecimal("1E-99999")

		_, err := tiny.Mul(tiny)

		require.Error(t, err)
		assert.ErrorContains(t, err, "mul failed")
		assert.ErrorContains(t, err, "exponent out of range")
	})

	t.Run("InvalidOperation errors on signaling NaN input", func(t *testing.T) {
		// sNaN in any arithmetic operand triggers InvalidOperation.
		// A quiet NaN ("NaN") does NOT — see the "quiet NaN propagates" case.
		_, err := MustNewDecimal("sNaN").Add(MustNewDecimal("1"))

		require.Error(t, err)
		assert.ErrorContains(t, err, "add failed")
		assert.ErrorContains(t, err, "invalid operation")
	})

	t.Run("InvalidOperation errors on Infinity - Infinity", func(t *testing.T) {
		// Inf - Inf is mathematically undefined and triggers InvalidOperation.
		// Inf + Inf = Inf (no signal); only the undefined cases trap.
		inf := MustNewDecimal("Infinity")

		_, err := inf.Sub(inf)

		require.Error(t, err)
		assert.ErrorContains(t, err, "sub failed")
		assert.ErrorContains(t, err, "invalid operation")
	})

	// --- Silent today (would flip to errors under a widened trap set) ---

	t.Run("DivisionUndefined silently returns NaN on 0/0", func(t *testing.T) {
		// 0/0 is DivisionUndefined in apd, which is NOT in our narrow trap
		// set but IS in apd.DefaultTraps. Adopting DefaultTraps flips this
		// subtest to an error case — that's the intended visibility.
		result, err := Zero().Div(Zero())

		require.NoError(t, err,
			"DivisionUndefined is not in the current narrow trap set")
		assert.Equal(t, "NaN", result.String(),
			"untrapped DivisionUndefined produces a NaN result")
	})

	t.Run("quiet NaN propagates silently through arithmetic", func(t *testing.T) {
		// IEEE 754 behavior: quiet NaN as an operand yields NaN with NO
		// condition flag raised. Distinct from sNaN which DOES signal.
		// This is normal/correct behavior even under apd.DefaultTraps —
		// the way to guard against this is validating inputs, not traps.
		result, err := MustNewDecimal("NaN").Add(MustNewDecimal("1"))

		require.NoError(t, err)
		assert.Equal(t, "NaN", result.String())
	})

	t.Run("Inexact silently rounds to 34 significant digits", func(t *testing.T) {
		// Inexact and Rounded fire on 10/3. Neither is in apd.DefaultTraps
		// — trapping them would make basic arithmetic unusable. This subtest
		// pins that 10/3 produces a 34-digit approximation with no error,
		// even after any reasonable trap-set change.
		result, err := MustNewDecimal("10").Div(MustNewDecimal("3"))

		require.NoError(t, err)
		assert.Equal(t, "3.333333333333333333333333333333333", result.String(),
			"34 significant digits per decimal128 precision")
	})
}
