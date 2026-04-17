package quanta

import "fmt"

// Unit represents the unit of measure for a quantity.
// It combines a Code (e.g., "USD", "tokens") with a Quantum to define
// how values should be quantized and what operations are compatible.
//
// The code and quantum are inseparable - "USD with 0.01 quantum" is one concept,
// not two separate things. You cannot have "USD" with a different quantum.
//
// Design principles:
// - Prevents mixing incompatible codes (USD vs EUR, tokens vs API calls)
// - Defines the quantum for snapping values to exact multiples
// - Code and quantum together form the complete unit definition
// - Two compatibility modes:
//   - CompatibleForCalc: Same code, quantum can differ (for Measure operations)
//   - Equal: Same code AND quantum (for Quantized aggregations)
type Unit struct {
	code    string
	quantum Quantum
}

// NewUnit creates a new Unit from primitive arguments: code and quantum string.
// Returns error if code is empty or quantum is invalid.
func NewUnit(code string, quantum string) (Unit, error) {
	if code == "" {
		return Unit{}, fmt.Errorf("unit code cannot be empty")
	}
	q, err := NewQuantum(quantum)
	if err != nil {
		return Unit{}, fmt.Errorf("invalid quantum: %w", err)
	}
	return Unit{code: code, quantum: q}, nil
}

// MustNewUnit creates a new Unit from primitive arguments.
// Panics if any argument is invalid. Use for testing and known-valid constants.
func MustNewUnit(code string, quantum string) Unit {
	u, err := NewUnit(code, quantum)
	if err != nil {
		panic(err)
	}
	return u
}

// Code returns the unit's code (e.g., "USD", "tokens", "api-credits").
func (u Unit) Code() string {
	return u.code
}

// Quantum returns the unit's quantum (precision level).
func (u Unit) Quantum() Quantum {
	return u.quantum
}

// CompatibleForCalc returns true if two units can be used together in calculations.
// Compatible units have the same code; quantum values may differ.
//
// Use this for Measure operations (Add, Sub, Mul, Div) where we want to allow
// high-precision calculations even if the final quantum targets differ.
//
// Example:
//
//	unitA := MustNewUnit("USD", "0.01")
//	unitB := MustNewUnit("USD", "0.001")
//	unitA.CompatibleForCalc(unitB) → true (same currency, different precision)
//
//	unitC := MustNewUnit("EUR", "0.01")
//	unitA.CompatibleForCalc(unitC) → false (different currency)
func (u Unit) CompatibleForCalc(other Unit) bool {
	return u.code == other.code
}

// Equal returns true if two units are exactly equal (same code AND quantum).
//
// Use this for Quantized aggregations where we want to ensure all values
// are snapped to the same quantum before combining them.
//
// Example:
//
//	unitA := MustNewUnit("USD", "0.01")
//	unitB := MustNewUnit("USD", "0.01")
//	unitA.Equal(unitB) → true
//
//	unitC := MustNewUnit("USD", "0.001")
//	unitA.Equal(unitC) → false (same currency, different quantum)
func (u Unit) Equal(other Unit) bool {
	return u.code == other.code && u.quantum.Equal(other.quantum)
}

// String returns a string representation of the unit.
func (u Unit) String() string {
	return fmt.Sprintf("%s[q=%s]", u.code, u.quantum.String())
}
