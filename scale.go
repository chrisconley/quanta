package quanta

import "fmt"

// Scale is a linear conversion from one unit to another: it captures the
// expected input unit, a scalar factor, and the output unit.
//
// Apply verifies the incoming Measure is in the expected input unit before
// multiplying and re-tagging. This closes the silent-conversion hole where
// a "multiply and re-tag" helper will accept any Measure and silently
// produce a result in the output unit — regardless of what the caller
// actually passed in.
//
// Example:
//
//	rate := MustNewScale(ScaleSpec{
//	    Input:  UnitSpec{Code: "tokens", Quantum: "1"},
//	    Output: UnitSpec{Code: "USD", Quantum: "0.01"},
//	    Factor: "0.05",
//	})
//
//	usage := MustNewMeasure(MeasureSpec{
//	    Unit:     UnitSpec{Code: "tokens", Quantum: "1"},
//	    Quantity: "1000",
//	})
//	cost, _ := rate.Apply(usage) // 50 USD
//
//	wrong := MustNewMeasure(MeasureSpec{
//	    Unit:     UnitSpec{Code: "credits", Quantum: "1"},
//	    Quantity: "1000",
//	})
//	_, err := rate.Apply(wrong) // error: incompatible input unit
type Scale struct {
	input  Unit
	output Unit
	factor Decimal
}

// ScaleSpec carries primitive construction data for a Scale.
type ScaleSpec struct {
	Input  UnitSpec
	Output UnitSpec
	Factor string
}

// NewScale creates a Scale from a primitive-only spec.
// Returns an error if any field is invalid.
func NewScale(spec ScaleSpec) (Scale, error) {
	input, err := NewUnit(spec.Input.Code, spec.Input.Quantum)
	if err != nil {
		return Scale{}, fmt.Errorf("invalid input unit: %w", err)
	}
	output, err := NewUnit(spec.Output.Code, spec.Output.Quantum)
	if err != nil {
		return Scale{}, fmt.Errorf("invalid output unit: %w", err)
	}
	factor, err := NewDecimal(spec.Factor)
	if err != nil {
		return Scale{}, fmt.Errorf("invalid factor: %w", err)
	}
	return Scale{input: input, output: output, factor: factor}, nil
}

// MustNewScale creates a Scale from a primitive-only spec.
// Panics if the spec is invalid. Use for testing and known-valid constants.
func MustNewScale(spec ScaleSpec) Scale {
	s, err := NewScale(spec)
	if err != nil {
		panic(err)
	}
	return s
}

// NewScaleFrom creates a Scale from already-constructed domain objects.
// No validation — trust the caller that input and output are valid Units.
// Mirrors NewMeasureFrom for composing Scales inside the package.
func NewScaleFrom(input Unit, output Unit, factor Decimal) Scale {
	return Scale{input: input, output: output, factor: factor}
}

// InputUnit returns the unit the Scale expects as input.
func (s Scale) InputUnit() Unit { return s.input }

// OutputUnit returns the unit the Scale tags its output with.
func (s Scale) OutputUnit() Unit { return s.output }

// Factor returns the scalar multiplier.
func (s Scale) Factor() Decimal { return s.factor }

// Apply converts m from the input unit to the output unit by multiplying
// its quantity by the factor and re-tagging. Returns an error if m's unit
// is not CompatibleForCalc with the Scale's input unit (same code; quantum
// may differ).
func (s Scale) Apply(m Measure) (Measure, error) {
	if !m.Unit().CompatibleForCalc(s.input) {
		return Measure{}, fmt.Errorf(
			"scale input mismatch: measure unit %s is not compatible with scale input unit %s",
			m.Unit().String(), s.input.String(),
		)
	}
	product, err := m.Quantity().Mul(s.factor)
	if err != nil {
		return Measure{}, fmt.Errorf("scale multiply error: %w", err)
	}
	return NewMeasureFrom(s.output, product), nil
}

// Equal returns true if two Scales have the same input unit, output unit,
// and factor. Units are compared with Unit.Equal (same code AND quantum).
func (s Scale) Equal(other Scale) bool {
	return s.input.Equal(other.input) &&
		s.output.Equal(other.output) &&
		s.factor.Equal(other.factor)
}
