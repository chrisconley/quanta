package quanta

import "fmt"

// Measure represents a high-precision value during calculations.
//
// Design principles:
// - High-precision working value - no snapping to quantum during calculations
// - Used in pipelines: rate calculations, aggregations, transformations
// - Never stored or sent across boundaries (use Quantized for that)
// - Immutable - all operations return new Measure instances
// - Operations check unit compatibility (same code, quantum can differ)
//
// To post a value across a boundary (storage, API, invoice):
//
//	measure.Quantize(rounding) → Quantized
//
// Examples:
//
//	usd := UnitSpec{Code: "USD", Quantum: "0.01"}
//	price := MustNewMeasure(MeasureSpec{Unit: usd, Quantity: "19.99"})
//	tax := MustNewMeasure(MeasureSpec{Unit: usd, Quantity: "1.60"})
//	total, _ := price.Add(tax) // 21.59 (high precision, not quantized yet)
//	quantized := total.Quantize(RoundHalfEven) // Snap to 0.01 quantum
type Measure struct {
	unit     Unit
	quantity Decimal
}

// UnitSpec carries primitive construction data for a Unit.
type UnitSpec struct {
	Code    string
	Quantum string
}

// MeasureSpec carries primitive construction data for a Measure.
type MeasureSpec struct {
	Unit     UnitSpec
	Quantity string
}

// NewMeasure creates a new Measure from a primitive-only spec.
// Returns error if any field is invalid.
func NewMeasure(spec MeasureSpec) (Measure, error) {
	u, err := NewUnit(spec.Unit.Code, spec.Unit.Quantum)
	if err != nil {
		return Measure{}, fmt.Errorf("invalid unit: %w", err)
	}
	d, err := NewDecimal(spec.Quantity)
	if err != nil {
		return Measure{}, fmt.Errorf("invalid quantity: %w", err)
	}
	return Measure{unit: u, quantity: d}, nil
}

// MustNewMeasure creates a new Measure from a primitive-only spec.
// Panics if any field is invalid. Use for testing and known-valid constants.
func MustNewMeasure(spec MeasureSpec) Measure {
	m, err := NewMeasure(spec)
	if err != nil {
		panic(err)
	}
	return m
}

// NewMeasureFrom creates a new Measure with the given unit and quantity.
func NewMeasureFrom(unit Unit, quantity Decimal) Measure {
	return Measure{unit: unit, quantity: quantity}
}

// Unit returns the measure's unit (code + quantum).
func (m Measure) Unit() Unit {
	return m.unit
}

// Quantity returns the measure's numeric quantity.
func (m Measure) Quantity() Decimal {
	return m.quantity
}

// Add adds two measures and returns a new measure with the sum.
// Returns error if units are not compatible for calculation (different codes).
func (m Measure) Add(other Measure) (Measure, error) {
	if !m.unit.CompatibleForCalc(other.unit) {
		return Measure{}, fmt.Errorf(
			"cannot add measures with incompatible specs: %s vs %s",
			m.unit.String(), other.unit.String(),
		)
	}

	sum, err := m.quantity.Add(other.quantity)
	if err != nil {
		return Measure{}, fmt.Errorf("error adding values: %w", err)
	}

	return NewMeasureFrom(m.unit, sum), nil
}

// Sub subtracts another measure from this one and returns a new measure with the difference.
// Returns error if units are not compatible for calculation (different codes).
func (m Measure) Sub(other Measure) (Measure, error) {
	if !m.unit.CompatibleForCalc(other.unit) {
		return Measure{}, fmt.Errorf(
			"cannot subtract measures with incompatible specs: %s vs %s",
			m.unit.String(), other.unit.String(),
		)
	}

	diff, err := m.quantity.Sub(other.quantity)
	if err != nil {
		return Measure{}, fmt.Errorf("error subtracting values: %w", err)
	}

	return NewMeasureFrom(m.unit, diff), nil
}

// Mul multiplies this measure by a scalar Decimal and returns a new measure with the product.
// The result retains the same spec as the original measure.
func (m Measure) Mul(scalar Decimal) (Measure, error) {
	product, err := m.quantity.Mul(scalar)
	if err != nil {
		return Measure{}, fmt.Errorf("error multiplying value: %w", err)
	}

	return NewMeasureFrom(m.unit, product), nil
}

// Div divides this measure by a scalar Decimal and returns a new measure with the quotient.
// The result retains the same spec as the original measure.
// Returns error if scalar is zero.
func (m Measure) Div(scalar Decimal) (Measure, error) {
	quotient, err := m.quantity.Div(scalar)
	if err != nil {
		return Measure{}, fmt.Errorf("error dividing value: %w", err)
	}

	return NewMeasureFrom(m.unit, quotient), nil
}

// Quantize snaps this measure to its quantum and returns a QuantizationResult.
// This is the boundary operation that converts a working Measure into a
// posted fact Quantized that can be stored, sent in APIs, or shown on invoices.
//
// Returns both the quantized value and the remainder that was rounded off.
// The rounding parameter determines how to round when the value doesn't fall
// exactly on a quantum boundary.
//
// Example:
//
//	usdSpec := MustNewUnit("USD", "0.01")
//	measure := NewMeasureFrom(usdSpec, MustNewDecimal("12.3456"))
//	result := measure.Quantize(RoundHalfEven)
//	posted := result.Value      // $12.35 (snapped to nearest $0.01)
//	dust := result.Remainder    // $-0.0044 (rounded up)
func (m Measure) Quantize(rounding Rounding) QuantizationResult {
	return deprecatedNewQuantized(m.unit, m.quantity, rounding)
}

// Equal returns true if two measures represent the same quantity in the same unit.
func (m Measure) Equal(other Measure) bool {
	return m.unit.Equal(other.unit) && m.quantity.Equal(other.quantity)
}

// String returns a string representation of the measure.
func (m Measure) String() string {
	return fmt.Sprintf("Measure[%s: %s]", m.unit.String(), m.quantity.String())
}
