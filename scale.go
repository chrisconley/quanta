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
//	tokens := MustNewUnit("tokens", "1")
//	usd := MustNewUnit("USD", "0.01")
//	rate := MustNewScale(tokens, usd, MustNewDecimal("0.05"))
//
//	usage := NewMeasureFrom(tokens, MustNewDecimal("1000"))
//	cost, _ := rate.Apply(usage) // 50 USD
//
//	wrongUnit := NewMeasureFrom(MustNewUnit("credits", "1"), MustNewDecimal("1000"))
//	_, err := rate.Apply(wrongUnit) // error: incompatible input unit
type Scale struct {
	input  Unit
	output Unit
	factor Decimal
}

// NewScale creates a Scale from the input unit, output unit, and factor.
// Returns an error if either unit has an empty code (the zero value).
// The factor may be any Decimal, including zero or negative.
func NewScale(input Unit, output Unit, factor Decimal) (Scale, error) {
	if input.Code() == "" {
		return Scale{}, fmt.Errorf("scale input unit must not be empty")
	}
	if output.Code() == "" {
		return Scale{}, fmt.Errorf("scale output unit must not be empty")
	}
	return Scale{input: input, output: output, factor: factor}, nil
}

// MustNewScale creates a Scale and panics if construction fails.
// Use for tests and known-valid constants.
func MustNewScale(input Unit, output Unit, factor Decimal) Scale {
	s, err := NewScale(input, output, factor)
	if err != nil {
		panic(err)
	}
	return s
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
