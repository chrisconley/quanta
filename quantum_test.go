package quanta

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewQuantum(t *testing.T) {
	t.Run("creates valid quantum", func(t *testing.T) {
		q, err := NewQuantum("0.01")
		assert.NoError(t, err)
		assert.Equal(t, "0.01", fmt.Sprintf("%s", q))
	})

	t.Run("accepts whole-number quantum", func(t *testing.T) {
		q, err := NewQuantum("1")
		assert.NoError(t, err)
		assert.Equal(t, "1", fmt.Sprintf("%s", q))
	})

	t.Run("accepts quantum exactly 1e-15", func(t *testing.T) {
		q, err := NewQuantum("0.000000000000001")
		assert.NoError(t, err)
		assert.True(t, q.Decimal().Equal(MustNewDecimal("0.000000000000001")))
	})

	t.Run("rejects empty string", func(t *testing.T) {
		_, err := NewQuantum("")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid decimal")
	})

	t.Run("rejects invalid decimal string", func(t *testing.T) {
		_, err := NewQuantum("not-a-number")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid decimal")
	})

	t.Run("rejects zero", func(t *testing.T) {
		_, err := NewQuantum("0")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "must be positive")
	})

	t.Run("rejects negative", func(t *testing.T) {
		_, err := NewQuantum("-0.01")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "must be positive")
	})

	t.Run("rejects quantum smaller than 1e-15", func(t *testing.T) {
		_, err := NewQuantum("0.0000000000000001") // 1e-16
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "must be >= 1e-15")
	})
}

func TestMustNewQuantum(t *testing.T) {
	t.Run("creates valid quantum without error", func(t *testing.T) {
		q := MustNewQuantum("0.01")
		assert.Equal(t, "0.01", fmt.Sprintf("%s", q))
	})

	t.Run("panics on invalid string", func(t *testing.T) {
		assert.Panics(t, func() {
			MustNewQuantum("not-a-number")
		})
	})

	t.Run("panics on zero", func(t *testing.T) {
		assert.Panics(t, func() {
			MustNewQuantum("0")
		})
	})

	t.Run("panic value is an error", func(t *testing.T) {
		defer func() {
			r := recover()
			assert.NotNil(t, r)
			_, ok := r.(error)
			assert.True(t, ok, "expected panic value to be an error, got %T", r)
		}()
		MustNewQuantum("not-a-number")
	})
}

func TestDeprecatedNewQuantum(t *testing.T) {
	t.Run("creates valid quantum", func(t *testing.T) {
		value := MustNewDecimal("0.01")
		q, err := DeprecatedNewQuantum(value)
		assert.NoError(t, err)
		assert.True(t, q.Decimal().Equal(value))
	})

	t.Run("rejects zero quantum", func(t *testing.T) {
		_, err := DeprecatedNewQuantum(Zero())
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "must be positive")
	})

	t.Run("rejects negative quantum", func(t *testing.T) {
		value := MustNewDecimal("-0.01")
		_, err := DeprecatedNewQuantum(value)
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "must be positive")
	})
}

func TestDeprecatedMustNewQuantum(t *testing.T) {
	t.Run("creates valid quantum without error", func(t *testing.T) {
		value := MustNewDecimal("0.01")
		q := DeprecatedMustNewQuantum(value)
		assert.True(t, q.Decimal().Equal(value))
	})

	t.Run("panics on invalid quantum", func(t *testing.T) {
		assert.Panics(t, func() {
			DeprecatedMustNewQuantum(Zero())
		})
	})

	t.Run("panic value is an error", func(t *testing.T) {
		defer func() {
			r := recover()
			assert.NotNil(t, r)
			_, ok := r.(error)
			assert.True(t, ok, "expected panic value to be an error, got %T", r)
		}()
		DeprecatedMustNewQuantum(Zero())
	})
}

func TestQuantum_Decimal(t *testing.T) {
	t.Run("returns underlying decimal value", func(t *testing.T) {
		expected := MustNewDecimal("0.001")
		q := DeprecatedMustNewQuantum(expected)
		assert.True(t, q.Decimal().Equal(expected))
	})
}

func TestQuantum_String(t *testing.T) {
	t.Run("returns string representation", func(t *testing.T) {
		q := MustNewQuantum("0.01")
		assert.Equal(t, "0.01", fmt.Sprintf("%s", q))
	})
}

func TestQuantum_Equality(t *testing.T) {
	t.Run("non-comparable", func(t *testing.T) {
		assert.False(t, reflect.TypeOf(Quantum{}).Comparable())
	})
	t.Run("returns true for equal quanta", func(t *testing.T) {
		q1 := MustNewQuantum("0.01")
		q2 := MustNewQuantum("0.01")
		assert.True(t, q1.Equal(q2))
	})

	t.Run("returns false for different quanta", func(t *testing.T) {
		q1 := MustNewQuantum("0.01")
		q2 := MustNewQuantum("0.001")
		assert.False(t, q1.Equal(q2))
	})

	t.Run("handles equivalent decimal representations", func(t *testing.T) {
		q1 := MustNewQuantum("0.01")
		q2 := MustNewQuantum("0.010")
		assert.True(t, q1.Equal(q2))
	})
}

func TestQuantum_ConvenienceConstructors(t *testing.T) {
	t.Run("NewQuantum001 creates 0.01 quantum", func(t *testing.T) {
		q := NewQuantum001()
		assert.Equal(t, "0.01", fmt.Sprintf("%s", q))
	})

	t.Run("NewQuantum1 creates 1 quantum", func(t *testing.T) {
		q := NewQuantum1()
		assert.Equal(t, "1", fmt.Sprintf("%s", q))
	})

	t.Run("NewQuantum0001 creates 0.001 quantum", func(t *testing.T) {
		q := NewQuantum0001()
		assert.Equal(t, "0.001", fmt.Sprintf("%s", q))
	})
}

func TestQuantum_EdgeCases(t *testing.T) {
	t.Run("very large quantum is accepted", func(t *testing.T) {
		q, err := NewQuantum("1000000")
		assert.NoError(t, err)
		assert.Equal(t, "1000000", fmt.Sprintf("%s", q))
	})

	t.Run("quantum with many decimal places is accepted", func(t *testing.T) {
		q, err := NewQuantum("0.00000000000001") // 1e-14
		assert.NoError(t, err)
		expected := MustNewDecimal("0.00000000000001")
		assert.True(t, q.Decimal().Equal(expected))
	})

	t.Run("quantum just below minimum is rejected", func(t *testing.T) {
		// Slightly less than 1e-15
		_, err := NewQuantum("0.0000000000000009")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "must be >= 1e-15")
	})
}
