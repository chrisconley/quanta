package quanta_test

import (
	"fmt"
	"testing"

	"github.com/chrisconley/quanta"
	"github.com/cockroachdb/apd/v3"
	"github.com/stretchr/testify/assert"
)

// These tests demonstrate common gotchas with decimal arithmetic and the
// tradeoffs precision makes to eliminate them. Each test shows the raw apd
// behavior first (which is correct — apd is doing what it's designed to do),
// then shows what precision trades away to remove the footgun.

// TestGotcha_InPlaceMutation demonstrates apd's destination-pointer pattern.
// apd operations write results into a caller-provided *Decimal, which is
// efficient (no allocation) but means using an input as the destination
// silently overwrites the original value.
//
// Tradeoff: precision allocates a new Decimal on every operation.
func TestGotcha_InPlaceMutation(t *testing.T) {
	// --- apd: destination-pointer pattern ---
	price := new(apd.Decimal)
	price.SetString("19.99")

	taxRate := new(apd.Decimal)
	taxRate.SetString("0.08875")

	ctx := &apd.Context{
		Precision: 34,
		Rounding:  apd.RoundHalfEven,
		Traps:     apd.InvalidOperation | apd.DivisionByZero | apd.Overflow,
	}

	// Compute tax using price as the destination — original price is gone
	expectedProduct := new(apd.Decimal)
	expectedProduct.SetString("1.7741125")
	ctx.Mul(price, price, taxRate)
	assert.Equal(t, 0, price.Cmp(expectedProduct),
		"apd: price variable now holds the product, not the original")

	// --- precision: every operation returns a new value ---
	pPrice := quanta.MustNewDecimal("19.99")
	pTaxRate := quanta.MustNewDecimal("0.08875")
	product, _ := pPrice.Mul(pTaxRate)

	assert.True(t, pPrice.Equal(quanta.MustNewDecimal("19.99")),
		"quanta: original value is preserved")
	assert.True(t, product.Equal(quanta.MustNewDecimal("1.7741125")),
		"quanta: result is a separate value")
}

// TestGotcha_StringRepresentation demonstrates that decimal values can have
// multiple string representations. "1.0" and "1.00" are mathematically equal
// but produce different strings because apd preserves the original scale.
// This is correct behavior (scale can be meaningful), but it's a gotcha
// if you use string output for equality checks or map keys.
//
// Tradeoff: precision adds Key() which normalizes trailing zeros, and Equal()
// which uses mathematical comparison.
func TestGotcha_StringRepresentation(t *testing.T) {
	// --- apd: string output preserves original scale ---
	a := new(apd.Decimal)
	a.SetString("1.0")
	b := new(apd.Decimal)
	b.SetString("1.00")

	assert.NotEqual(t, fmt.Sprintf("%v", a), fmt.Sprintf("%v", b),
		"apd: same value, different string output")
	assert.Equal(t, 0, a.Cmp(b),
		"apd: Cmp correctly identifies them as equal")

	// --- precision: Key() normalizes, Equal() uses Cmp ---
	pa := quanta.MustNewDecimal("1.0")
	pb := quanta.MustNewDecimal("1.00")

	assert.NotEqual(t, fmt.Sprintf("%v", pa), fmt.Sprintf("%v", pb),
		"quanta: Sprintf still preserves scale (same as apd)")
	assert.Equal(t, pa.Key(), pb.Key(),
		"quanta: Key() normalizes trailing zeros for map keys")
	assert.True(t, pa.Equal(pb),
		"quanta: Equal() uses mathematical comparison")
}

// TestGotcha_SilentPrecisionLoss demonstrates apd's configurable precision.
// apd lets you set any precision, and operations that lose digits return a
// Condition flag (Inexact/Rounded) — but error is nil unless that condition
// is explicitly trapped. This is flexible and correct, but if you configure
// a low precision without trapping Inexact, results are silently truncated.
//
// Tradeoff: precision hardcodes decimal128 (34 significant digits) and traps
// InvalidOperation, DivisionByZero, and Overflow. You can't configure a
// different precision, but you also can't accidentally use the wrong one.
func TestGotcha_SilentPrecisionLoss(t *testing.T) {
	// --- apd: configurable precision, condition-based signaling ---
	lowPrec := &apd.Context{
		Precision:   4,
		Rounding:    apd.RoundHalfEven,
		MaxExponent: apd.MaxExponent,
		MinExponent: apd.MinExponent,
		// Traps not set — Inexact/Rounded won't produce an error
	}

	hundred := new(apd.Decimal)
	hundred.SetString("100")
	three := new(apd.Decimal)
	three.SetString("3")

	var result apd.Decimal
	cond, err := lowPrec.Quo(&result, hundred, three)

	expected := new(apd.Decimal)
	expected.SetString("33.33")

	assert.NoError(t, err,
		"apd: no error — Inexact is a condition flag, not a trapped error")
	assert.True(t, cond.Inexact(),
		"apd: Inexact condition is set (easy to miss if you discard the return)")
	assert.Equal(t, 0, result.Cmp(expected),
		"apd: result truncated to 4 significant digits")

	// --- precision: fixed decimal128, 34 significant digits ---
	pHundred := quanta.MustNewDecimal("100")
	pThree := quanta.MustNewDecimal("3")
	pResult, _ := pHundred.Div(pThree)

	assert.True(t,
		pResult.Equal(quanta.MustNewDecimal("33.33333333333333333333333333333333")),
		"quanta: 34 digits — enough for any financial calculation")
}
