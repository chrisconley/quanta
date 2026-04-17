package quanta_test

import (
	"sync"
	"testing"

	"github.com/chrisconley/quanta"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// The tests in this file verify that quanta's public types are safe to
// share across goroutines. Run with `go test -race ./...` to catch any
// data races.
//
// Design invariant being tested:
//   - Decimal, Measure, Quantized, Unit, and Quantum are immutable after
//     construction. Every operation returns a new value; nothing mutates
//     the receiver. Concurrent readers cannot observe torn state.

const (
	concurrentGoroutines = 64
	concurrentIterations = 200
)

// TestConcurrent_Decimal_ReadsAndArithmetic shares a Decimal across many
// goroutines. Half perform reads (String, Cmp, Key, …), half perform
// arithmetic that returns new Decimals. The shared instance must never
// mutate and every Mul call must produce identical output.
func TestConcurrent_Decimal_ReadsAndArithmetic(t *testing.T) {
	shared := quanta.MustNewDecimal("19.99")
	factor := quanta.MustNewDecimal("1.08875")

	expectedProduct, err := shared.Mul(factor)
	require.NoError(t, err)
	expectedKey := shared.Key()

	var wg sync.WaitGroup
	wg.Add(concurrentGoroutines)

	for range concurrentGoroutines/2 {
		go func() {
			defer wg.Done()
			for range concurrentIterations {
				assert.Equal(t, expectedKey, shared.Key())
				assert.False(t, shared.IsZero())
				assert.False(t, shared.IsNegative())
				assert.Equal(t, -1, shared.Cmp(quanta.MustNewDecimal("20")))
				_ = shared.String()
			}
		}()
	}

	for range concurrentGoroutines/2 {
		go func() {
			defer wg.Done()
			for range concurrentIterations {
				product, err := shared.Mul(factor)
				assert.NoError(t, err)
				assert.True(t, product.Equal(expectedProduct))

				sum, err := shared.Add(factor)
				assert.NoError(t, err)
				assert.False(t, sum.IsZero())
			}
		}()
	}

	wg.Wait()

	assert.True(t, shared.Equal(quanta.MustNewDecimal("19.99")),
		"shared Decimal must be unchanged after concurrent use")
	assert.True(t, factor.Equal(quanta.MustNewDecimal("1.08875")),
		"shared factor must be unchanged after concurrent use")
}

// TestConcurrent_Measure_Arithmetic shares a Measure across goroutines
// performing Add, Sub, Mul, Div. All operations return new Measures;
// the shared instances must never mutate.
func TestConcurrent_Measure_Arithmetic(t *testing.T) {
	usd := quanta.MustNewUnit("USD", "0.01")
	shared := quanta.NewMeasureFrom(usd, quanta.MustNewDecimal("100.00"))
	addend := quanta.NewMeasureFrom(usd, quanta.MustNewDecimal("25.50"))
	scalar := quanta.MustNewDecimal("1.5")

	expectedSum, err := shared.Add(addend)
	require.NoError(t, err)
	expectedProduct, err := shared.Mul(scalar)
	require.NoError(t, err)

	var wg sync.WaitGroup
	wg.Add(concurrentGoroutines)

	for range concurrentGoroutines {
		go func() {
			defer wg.Done()
			for range concurrentIterations {
				sum, err := shared.Add(addend)
				assert.NoError(t, err)
				assert.True(t, sum.Equal(expectedSum))

				product, err := shared.Mul(scalar)
				assert.NoError(t, err)
				assert.True(t, product.Equal(expectedProduct))
			}
		}()
	}

	wg.Wait()

	assert.True(t,
		shared.Quantity().Equal(quanta.MustNewDecimal("100.00")),
		"shared Measure's quantity must remain unchanged")
	assert.True(t,
		addend.Quantity().Equal(quanta.MustNewDecimal("25.50")),
		"shared addend must remain unchanged")
}

// TestConcurrent_Quantize repeatedly quantizes the same Measure from many
// goroutines. Quantization is deterministic; every call with the same
// rounding mode must produce the same Quantized value and the same dust.
func TestConcurrent_Quantize(t *testing.T) {
	usd := quanta.MustNewUnit("USD", "0.01")
	shared := quanta.NewMeasureFrom(usd, quanta.MustNewDecimal("19.9999"))

	expected := shared.Quantize(quanta.RoundHalfEven)

	var wg sync.WaitGroup
	wg.Add(concurrentGoroutines)

	for range concurrentGoroutines {
		go func() {
			defer wg.Done()
			for range concurrentIterations {
				result := shared.Quantize(quanta.RoundHalfEven)
				assert.True(t, result.Value.Equal(expected.Value))
				assert.True(t, result.Remainder.Equal(expected.Remainder))
			}
		}()
	}

	wg.Wait()

	assert.True(t,
		shared.Quantity().Equal(quanta.MustNewDecimal("19.9999")),
		"Quantize must not mutate the source Measure")
}

// TestConcurrent_Quantized exercises concurrent reads on a shared Quantized.
// A Quantized is a Unit plus an int64 multiplier, so it is trivially
// immutable — this test locks the invariant in place so a future change
// that introduced lazy computation or caching would get caught.
func TestConcurrent_Quantized(t *testing.T) {
	usd := quanta.MustNewUnit("USD", "0.01")
	shared := quanta.DeprecatedNewQuantizedFromMultiplier(usd, 1999)
	twin := quanta.DeprecatedNewQuantizedFromMultiplier(usd, 1999)

	expectedDecimal := shared.Decimal()

	var wg sync.WaitGroup
	wg.Add(concurrentGoroutines)

	for range concurrentGoroutines {
		go func() {
			defer wg.Done()
			for range concurrentIterations {
				assert.True(t, shared.Equal(twin))
				assert.Equal(t, int64(1999), shared.Multiplier())
				assert.True(t, shared.Decimal().Equal(expectedDecimal))
				_ = shared.String()
			}
		}()
	}

	wg.Wait()
}

// TestConcurrent_Pipeline runs the full Measure → calc → Quantize pipeline
// across many goroutines with shared inputs. This is the end-to-end claim:
// a billing calculation that reads from shared rate cards and unit specs
// can be run in parallel without locking.
func TestConcurrent_Pipeline(t *testing.T) {
	usd := quanta.MustNewUnit("USD", "0.01")
	subtotal := quanta.NewMeasureFrom(usd, quanta.MustNewDecimal("19.99"))
	taxRate := quanta.MustNewDecimal("0.08875")

	computeTotal := func() quanta.Quantized {
		tax, err := subtotal.Mul(taxRate)
		require.NoError(t, err)
		total, err := subtotal.Add(tax)
		require.NoError(t, err)
		return total.Quantize(quanta.RoundHalfEven).Value
	}

	expected := computeTotal()

	var wg sync.WaitGroup
	wg.Add(concurrentGoroutines)

	for range concurrentGoroutines {
		go func() {
			defer wg.Done()
			for range concurrentIterations {
				result := computeTotal()
				assert.True(t, result.Equal(expected))
			}
		}()
	}

	wg.Wait()

	assert.True(t, subtotal.Quantity().Equal(quanta.MustNewDecimal("19.99")))
	assert.True(t, taxRate.Equal(quanta.MustNewDecimal("0.08875")))
}

// TestConcurrent_IncompatibleUnitsStillError verifies that error paths are
// also concurrent-safe: attempting USD + EUR from many goroutines must fail
// consistently without racing on unit state.
func TestConcurrent_IncompatibleUnitsStillError(t *testing.T) {
	dollars := quanta.NewMeasureFrom(
		quanta.MustNewUnit("USD", "0.01"),
		quanta.MustNewDecimal("100.00"),
	)
	euros := quanta.NewMeasureFrom(
		quanta.MustNewUnit("EUR", "0.01"),
		quanta.MustNewDecimal("90.00"),
	)

	var wg sync.WaitGroup
	wg.Add(concurrentGoroutines)

	for range concurrentGoroutines {
		go func() {
			defer wg.Done()
			for range concurrentIterations {
				_, err := dollars.Add(euros)
				assert.Error(t, err, "USD + EUR must always fail")
			}
		}()
	}

	wg.Wait()
}
