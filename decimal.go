package quanta

import (
	"fmt"
	"strings"

	"github.com/cockroachdb/apd/v3"
)

// Decimal is an immutable wrapper around apd.Decimal for precise decimal arithmetic
type Decimal struct {
	v apd.Decimal
	_ [0]func() // Makes the struct non-comparable, preventing errors using == or !=
}

// newCalcContext returns a context for high-precision decimal calculations
// with the given rounding mode. Uses IEEE 754 decimal128 precision (34 digits)
// and traps InvalidOperation, DivisionByZero, and Overflow so those surface
// as Go errors rather than silent flagged results.
//
// Built on apd.BaseContext (which sets MaxExponent/MinExponent to the package
// defaults and starts with DefaultTraps); we then override Rounding and
// narrow the trap set to quanta's chosen subset.
func newCalcContext(rounding apd.Rounder) *apd.Context {
	ctx := apd.BaseContext.WithPrecision(34)
	ctx.Rounding = rounding
	ctx.Traps = apd.InvalidOperation | apd.DivisionByZero | apd.Overflow
	return ctx
}

// NewDecimal creates a new Decimal from a string representation
func NewDecimal(s string) (Decimal, error) {
	var v apd.Decimal
	if _, _, err := v.SetString(strings.TrimSpace(s)); err != nil {
		return Decimal{}, fmt.Errorf("invalid decimal string %q: %w", s, err)
	}
	return Decimal{v: v}, nil
}

// MustNewDecimal creates a new Decimal from a string, panicking on error
func MustNewDecimal(s string) Decimal {
	v, err := NewDecimal(s)
	if err != nil {
		panic(err)
	}
	return v
}

// NewDecimalFromInt64 creates a new Decimal from an int64 value
func NewDecimalFromInt64(i int64) Decimal {
	return Decimal{v: *apd.New(i, 0)}
}

// toAPD returns a pointer to the internal apd.Decimal value.
// This is an internal method for use within the package only.
func (d *Decimal) toAPD() *apd.Decimal {
	return &d.v
}

// Zero returns a Decimal representing zero
func Zero() Decimal {
	return Decimal{v: *apd.New(0, 0)}
}

// String returns the string representation of the decimal
func (d Decimal) String() string {
	return d.v.String()
}

// Key returns a canonical string representation suitable for use as a map key
// or for exact equality comparisons. This normalizes the decimal by reducing
// trailing zeros to ensure equal values have identical keys.
func (d Decimal) Key() string {
	var normalized apd.Decimal
	normalized.Set(&d.v)
	apd.BaseContext.Reduce(&normalized, &normalized)
	return normalized.String()
}

// IsZero returns true if the decimal is zero
func (d Decimal) IsZero() bool {
	return d.v.IsZero()
}

// IsNegative returns true if the decimal is negative
func (d Decimal) IsNegative() bool {
	return d.v.Negative
}

// Equal returns true if two decimals are exactly equal in value
func (d Decimal) Equal(other Decimal) bool {
	return d.v.Cmp(&other.v) == 0
}

// Cmp compares two decimals. Returns:
// -1 if d < other
//
//	0 if d == other
//	1 if d > other
func (d Decimal) Cmp(other Decimal) int {
	return d.v.Cmp(&other.v)
}

// Add returns a new Decimal that is the sum of d and other.
func (d Decimal) Add(other Decimal) (Decimal, error) {
	var result apd.Decimal
	ctx := newCalcContext(apd.RoundHalfEven)
	_, err := ctx.Add(&result, &d.v, &other.v)
	if err != nil {
		return Decimal{}, fmt.Errorf("add failed: %w", err)
	}
	return Decimal{v: result}, nil
}

// Sub returns a new Decimal that is the difference of d and other.
func (d Decimal) Sub(other Decimal) (Decimal, error) {
	var result apd.Decimal
	ctx := newCalcContext(apd.RoundHalfEven)
	_, err := ctx.Sub(&result, &d.v, &other.v)
	if err != nil {
		return Decimal{}, fmt.Errorf("sub failed: %w", err)
	}
	return Decimal{v: result}, nil
}

// Mul returns a new Decimal that is the product of d and other.
func (d Decimal) Mul(other Decimal) (Decimal, error) {
	var result apd.Decimal
	ctx := newCalcContext(apd.RoundHalfEven)
	_, err := ctx.Mul(&result, &d.v, &other.v)
	if err != nil {
		return Decimal{}, fmt.Errorf("mul failed: %w", err)
	}
	return Decimal{v: result}, nil
}

// Div returns a new Decimal that is the quotient of d divided by other.
func (d Decimal) Div(other Decimal) (Decimal, error) {
	var result apd.Decimal
	ctx := newCalcContext(apd.RoundHalfEven)
	_, err := ctx.Quo(&result, &d.v, &other.v)
	if err != nil {
		return Decimal{}, fmt.Errorf("div failed: %w", err)
	}
	return Decimal{v: result}, nil
}

// Floor returns the largest integer value less than or equal to d.
// This rounds toward negative infinity.
//
// Examples:
//
//	MustNewDecimal("12.8").Floor()   → 12
//	MustNewDecimal("12.2").Floor()   → 12
//	MustNewDecimal("-12.2").Floor()  → -13
//	MustNewDecimal("-12.8").Floor()  → -13
func (d Decimal) Floor() int64 {
	ctx := newCalcContext(apd.RoundFloor)

	var rounded apd.Decimal
	_, err := ctx.RoundToIntegralValue(&rounded, &d.v)
	if err != nil {
		panic(fmt.Sprintf("error computing floor: %v", err))
	}

	result, err := rounded.Int64()
	if err != nil {
		panic(fmt.Sprintf("floor result too large for int64: %v", err))
	}

	return result
}

// Ceiling returns the smallest integer value greater than or equal to d.
// This rounds toward positive infinity.
//
// Examples:
//
//	MustNewDecimal("12.2").Ceiling()   → 13
//	MustNewDecimal("12.8").Ceiling()   → 13
//	MustNewDecimal("-12.8").Ceiling()  → -12
//	MustNewDecimal("-12.2").Ceiling()  → -12
func (d Decimal) Ceiling() int64 {
	ctx := newCalcContext(apd.RoundCeiling)

	var rounded apd.Decimal
	_, err := ctx.RoundToIntegralValue(&rounded, &d.v)
	if err != nil {
		panic(fmt.Sprintf("error computing ceiling: %v", err))
	}

	result, err := rounded.Int64()
	if err != nil {
		panic(fmt.Sprintf("ceiling result too large for int64: %v", err))
	}

	return result
}

// Frac returns the fractional part of d.
// The result is always in the range [0, 1) for positive numbers and (-1, 0] for negative numbers.
// Calculated as: frac(d) = d - floor(d)
//
// Examples:
//
//	MustNewDecimal("12.34").Frac()   → 0.34
//	MustNewDecimal("12.00").Frac()   → 0
//	MustNewDecimal("-12.34").Frac()  → 0.66 (because -12.34 - (-13) = 0.66)
//
// Use case: Allocation algorithms need to track fractional parts to distribute dust.
func (d Decimal) Frac() Decimal {
	floor := d.Floor()
	floorDec := NewDecimalFromInt64(floor)

	frac, err := d.Sub(floorDec)
	if err != nil {
		// This should not happen - subtracting an integer from a decimal
		panic(fmt.Sprintf("error computing fractional part: %v", err))
	}

	return frac
}

// Abs returns the absolute value of d.
// Examples:
//
//	MustNewDecimal("12.34").Abs()   → 12.34
//	MustNewDecimal("-12.34").Abs()  → 12.34
func (d Decimal) Abs() Decimal {
	if d.IsNegative() {
		return d.Negate()
	}
	return d
}

// Negate returns the negation of d.
// Examples:
//
//	MustNewDecimal("12.34").Negate()   → -12.34
//	MustNewDecimal("-12.34").Negate()  → 12.34
//	MustNewDecimal("0").Negate()       → 0
func (d Decimal) Negate() Decimal {
	var result apd.Decimal
	result.Neg(&d.v)
	return Decimal{v: result}
}
