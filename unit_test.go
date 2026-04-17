package quanta

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUnit(t *testing.T) {
	t.Run("creates unit with code and quantum", func(t *testing.T) {
		unit, err := NewUnit("USD", "0.01")
		assert.NoError(t, err)
		assert.Equal(t, "USD", unit.Code())
		assert.True(t, unit.Quantum().Equal(NewQuantum001()))
	})

	t.Run("accepts different quantum values", func(t *testing.T) {
		unit, err := NewUnit("KWD", "0.001")
		assert.NoError(t, err)
		assert.Equal(t, "KWD", unit.Code())
		assert.Equal(t, "0.001", fmt.Sprintf("%s", unit.Quantum()))
	})

	t.Run("accepts whole-number quantum", func(t *testing.T) {
		unit, err := NewUnit("JPY", "1")
		assert.NoError(t, err)
		assert.Equal(t, "JPY", unit.Code())
		assert.True(t, unit.Quantum().Equal(NewQuantum1()))
	})

	t.Run("rejects empty code", func(t *testing.T) {
		_, err := NewUnit("", "0.01")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "unit code cannot be empty")
	})

	t.Run("rejects invalid quantum string", func(t *testing.T) {
		_, err := NewUnit("USD", "not-a-number")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid quantum")
	})

	t.Run("rejects empty quantum", func(t *testing.T) {
		_, err := NewUnit("USD", "")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid quantum")
	})

	t.Run("rejects zero quantum", func(t *testing.T) {
		_, err := NewUnit("USD", "0")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid quantum")
	})

	t.Run("rejects negative quantum", func(t *testing.T) {
		_, err := NewUnit("USD", "-0.01")
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "invalid quantum")
	})
}

func TestMustNewUnit(t *testing.T) {
	t.Run("creates unit without error", func(t *testing.T) {
		unit := MustNewUnit("USD", "0.01")
		assert.Equal(t, "USD", unit.Code())
		assert.True(t, unit.Quantum().Equal(NewQuantum001()))
	})

	t.Run("panics on empty code", func(t *testing.T) {
		assert.Panics(t, func() {
			MustNewUnit("", "0.01")
		})
	})

	t.Run("panics on invalid quantum", func(t *testing.T) {
		assert.Panics(t, func() {
			MustNewUnit("USD", "bad")
		})
	})

	t.Run("panic value is an error", func(t *testing.T) {
		defer func() {
			r := recover()
			assert.NotNil(t, r)
			_, ok := r.(error)
			assert.True(t, ok, "expected panic value to be an error, got %T", r)
		}()
		MustNewUnit("", "0.01")
	})
}

func TestSpec_Code(t *testing.T) {
	t.Run("returns spec code", func(t *testing.T) {
		spec := MustNewUnit("EUR", "0.01")
		assert.Equal(t, "EUR", spec.Code())
	})
}

func TestSpec_Quantum(t *testing.T) {
	t.Run("returns spec quantum", func(t *testing.T) {
		spec := MustNewUnit("KWD", "0.001")
		assert.True(t, spec.Quantum().Equal(MustNewQuantum("0.001")))
		assert.Equal(t, "0.001", fmt.Sprintf("%s", spec.Quantum()))
	})
}

func TestSpec_CompatibleForCalc(t *testing.T) {
	t.Run("same code with same quantum is compatible", func(t *testing.T) {
		spec1 := MustNewUnit("USD", "0.01")
		spec2 := MustNewUnit("USD", "0.01")
		assert.True(t, spec1.CompatibleForCalc(spec2))
	})

	t.Run("same code with different quantum is compatible", func(t *testing.T) {
		spec1 := MustNewUnit("USD", "0.01")
		spec2 := MustNewUnit("USD", "0.001")
		assert.True(t, spec1.CompatibleForCalc(spec2))
	})

	t.Run("different code with same quantum is not compatible", func(t *testing.T) {
		spec1 := MustNewUnit("USD", "0.01")
		spec2 := MustNewUnit("EUR", "0.01")
		assert.False(t, spec1.CompatibleForCalc(spec2))
	})

	t.Run("different code with different quantum is not compatible", func(t *testing.T) {
		spec1 := MustNewUnit("USD", "0.01")
		spec2 := MustNewUnit("EUR", "0.001")
		assert.False(t, spec1.CompatibleForCalc(spec2))
	})
}

func TestUnit_Equality(t *testing.T) {
	t.Run("non-comparable", func(t *testing.T) {
		assert.False(t, reflect.TypeOf(Unit{}).Comparable())
	})
	t.Run("same code and quantum are equal", func(t *testing.T) {
		spec1 := MustNewUnit("USD", "0.01")
		spec2 := MustNewUnit("USD", "0.01")
		assert.True(t, spec1.Equal(spec2))
	})

	t.Run("same code but different quantum are not equal", func(t *testing.T) {
		spec1 := MustNewUnit("USD", "0.01")
		spec2 := MustNewUnit("USD", "0.001")
		assert.False(t, spec1.Equal(spec2))
	})

	t.Run("different code with same quantum are not equal", func(t *testing.T) {
		spec1 := MustNewUnit("USD", "0.01")
		spec2 := MustNewUnit("EUR", "0.01")
		assert.False(t, spec1.Equal(spec2))
	})

	t.Run("equivalent quantum values are equal", func(t *testing.T) {
		spec1 := MustNewUnit("USD", "0.01")
		spec2 := MustNewUnit("USD", "0.010")
		assert.True(t, spec1.Equal(spec2))
	})
}

func TestSpec_String(t *testing.T) {
	t.Run("returns formatted string representation", func(t *testing.T) {
		spec := MustNewUnit("USD", "0.01")
		assert.Equal(t, "USD[q=0.01]", spec.String())
	})

	t.Run("includes quantum in string", func(t *testing.T) {
		spec := MustNewUnit("KWD", "0.001")
		assert.Equal(t, "KWD[q=0.001]", spec.String())
	})
}

func TestSpec_UsageCases(t *testing.T) {
	t.Run("currency specs", func(t *testing.T) {
		usdCents := MustNewUnit("USD", "0.01")
		jpyYen := MustNewUnit("JPY", "1")
		kwdFils := MustNewUnit("KWD", "0.001")

		assert.Equal(t, "USD[q=0.01]", fmt.Sprintf("%s", usdCents))
		assert.Equal(t, "JPY[q=1]", fmt.Sprintf("%s", jpyYen))
		assert.Equal(t, "KWD[q=0.001]", fmt.Sprintf("%s", kwdFils))
	})

	t.Run("usage specs", func(t *testing.T) {
		tokens := MustNewUnit("tokens", "0.001")
		apiCalls := MustNewUnit("api_calls", "1")
		bytes := MustNewUnit("bytes", "0.000001")

		assert.Equal(t, "tokens[q=0.001]", fmt.Sprintf("%s", tokens))
		assert.Equal(t, "api_calls[q=1]", fmt.Sprintf("%s", apiCalls))
		assert.Equal(t, "bytes[q=0.000001]", fmt.Sprintf("%s", bytes))
	})

	t.Run("calculating with compatible specs", func(t *testing.T) {
		// Two USD amounts with different precision can be calculated together
		highPrecision := MustNewUnit("USD", "0.0001")
		lowPrecision := MustNewUnit("USD", "0.01")

		assert.True(t, highPrecision.CompatibleForCalc(lowPrecision))
		assert.False(t, highPrecision.Equal(lowPrecision))
	})

	t.Run("aggregating quantized values requires equal specs", func(t *testing.T) {
		spec1 := MustNewUnit("USD", "0.01")
		spec2 := MustNewUnit("USD", "0.01")

		// For aggregating quantized values, specs must be exactly equal
		assert.True(t, spec1.Equal(spec2))
	})
}
