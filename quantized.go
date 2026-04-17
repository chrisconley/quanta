package quanta

import (
	"fmt"

	"github.com/cockroachdb/apd/v3"
)

// QuantizationResult represents the outcome of quantizing a Measure.
// It contains both the quantized value and the remainder that was rounded off.
//
// Design principles:
// - Ephemeral operation result - NOT a persistent domain object
// - Remainder is metadata about the quantization event
// - Caller decides what to do with remainder (track for dust, create audit record, discard)
//
// Example:
//
//	measure := NewMeasureFrom(usdSpec, MustNewDecimal("12.3456"))
//	result := measure.Quantize(RoundHalfEven)
//	posted := result.Value      // $12.35 (store this)
//	dust := result.Remainder    // $-0.0044 (rounded up, track if needed)
type QuantizationResult struct {
	Value     Quantized // The posted value (snapped to quantum)
	Remainder Measure   // What was rounded off (original - quantized)
}

// Quantized represents a value snapped to an exact multiple of its quantum.
//
// Design principles:
// - Read-only "posted fact" - no math operations (Add, Sub, Mul, Div)
// - Stored in database, sent in APIs, shown in invoices
// - Safe for exact equality checks (uses integer multiplier internally)
// - Created via Measure.Quantize(rounding) which returns QuantizationResult
// - No ToMeasure() - enforce boundary: calculate in Measure, quantize only at the end
//
// Examples:
//
//	usdSpec := MustNewUnit("USD", "0.01")
//	measure := NewMeasureFrom(usdSpec, MustNewDecimal("12.3456"))
//	result := measure.Quantize(RoundHalfEven)
//	posted := result.Value      // Quantized = $12.35
//	dust := result.Remainder    // Measure = $-0.0044 (rounded up)
//
// Internal representation:
// - Stores integer multiplier of quantum to avoid floating-point errors
// - Example: $12.35 with quantum 0.01 = multiplier 1235
type Quantized struct {
	unit       Unit
	multiplier int64 // Number of quanta (quantity = multiplier × quantum)
}

// deprecatedNewQuantized creates a QuantizationResult by snapping a value to the spec's quantum.
// This is an internal constructor called by Measure.Quantize().
// Returns both the quantized value and the remainder that was rounded off.
func deprecatedNewQuantized(unit Unit, quantity Decimal, rounding Rounding) QuantizationResult {
	quantum := unit.Quantum().Decimal()

	// Divide value by quantum to get multiplier
	ctx := newCalcContext(rounding.toAPDRounder())

	var quotient apd.Decimal
	_, err := ctx.Quo(&quotient, quantity.toAPD(), quantum.toAPD())
	if err != nil {
		// This should not happen with valid inputs
		panic(fmt.Sprintf("error quantizing value: %v", err))
	}

	// Round to integer
	var rounded apd.Decimal
	_, err = ctx.RoundToIntegralValue(&rounded, &quotient)
	if err != nil {
		panic(fmt.Sprintf("error rounding to integer: %v", err))
	}

	// Convert to int64
	multiplier, err := rounded.Int64()
	if err != nil {
		panic(fmt.Sprintf("multiplier too large for int64: %v", err))
	}

	// Calculate remainder: original - quantized
	// Reconstruct quantized value: multiplier × quantum
	multiplierDec := NewDecimalFromInt64(multiplier)
	quantizedValue, err := multiplierDec.Mul(quantum)
	if err != nil {
		panic(fmt.Sprintf("error calculating quantized value: %v", err))
	}

	// remainder = original - quantized
	remainderDec, err := quantity.Sub(quantizedValue)
	if err != nil {
		panic(fmt.Sprintf("error calculating remainder: %v", err))
	}

	return QuantizationResult{
		Value: Quantized{
			unit:       unit,
			multiplier: multiplier,
		},
		Remainder: NewMeasureFrom(unit, remainderDec),
	}
}

// DeprecatedNewQuantizedFromMultiplier creates a Quantized directly from a multiplier.
// This is useful when reading from storage where the multiplier is already known.
// No remainder is tracked since this creates an exact quantum-aligned quantity.
//
// Example:
//
//	usdSpec := MustNewUnit("USD", "0.01")
//	amount := DeprecatedNewQuantizedFromMultiplier(usdSpec, 1235) // $12.35 (1235 × 0.01)
func DeprecatedNewQuantizedFromMultiplier(unit Unit, multiplier int64) Quantized {
	return Quantized{
		unit:       unit,
		multiplier: multiplier,
	}
}

// Unit returns the quantized value's unit (code + quantum).
func (q Quantized) Unit() Unit {
	return q.unit
}

// Multiplier returns the integer multiplier (number of quanta).
// This is useful for storage and for exact arithmetic on quantized values.
func (q Quantized) Multiplier() int64 {
	return q.multiplier
}

// Decimal returns the quantized value as a Decimal.
// This reconstructs the decimal value from multiplier × quantum.
func (q Quantized) Decimal() Decimal {
	quantum := q.unit.Quantum().Decimal()
	multiplierDec := NewDecimalFromInt64(q.multiplier)

	result, err := multiplierDec.Mul(quantum)
	if err != nil {
		// This should not happen with valid quantized values
		panic(fmt.Sprintf("error converting quantized to decimal: %v", err))
	}

	return result
}

// Equal returns true if two quantized values are exactly equal.
// Two quantized values are equal if they have equal specs and the same multiplier.
//
// Example:
//
//	usdSpec := MustNewUnit("USD", "0.01")
//	q1 := DeprecatedNewQuantizedFromMultiplier(usdSpec, 1235) // $12.35 from storage
//	result := measure.Quantize(RoundHalfEven)
//	q2 := result.Value                               // $12.35 from calculation
//	q1.Equal(q2) → true (same posted value)
func (q Quantized) Equal(other Quantized) bool {
	return q.unit.Equal(other.unit) && q.multiplier == other.multiplier
}

// String returns a string representation of the quantized quantity.
func (q Quantized) String() string {
	return fmt.Sprintf("Quantized[%s: %s]", q.unit.String(), q.Decimal().String())
}
