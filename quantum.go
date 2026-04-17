package quanta

import "fmt"

// Quantum represents the smallest meaningful increment for a measure or money value.
// Design principles:
// - Must be > 0 and >= 1e-15 (prevents absurdly tiny quanta)
// - Used by Quantized to snap values to exact multiples
//
// Examples:
//   - Money quantum: 0.01 USD (cents), 1 JPY (yen), 0.001 KWD (fils)
//   - Usage quantum: 0.001 tokens, 1 API call, 0.000001 bytes
type Quantum struct {
	value Decimal
}

// minQuantum is the minimum allowed quantum value (1e-15).
// This prevents performance issues from absurdly tiny quanta while still
// being far more precise than any real-world currency or instrumentation.
var minQuantum = MustNewDecimal("0.000000000000001") // 1e-15

// DeprecatedNewQuantum creates a new Quantum from a Decimal value.
// Returns error if value <= 0 or < 1e-15.
func DeprecatedNewQuantum(value Decimal) (Quantum, error) {
	if value.Cmp(Zero()) <= 0 {
		return Quantum{}, fmt.Errorf("quantum must be positive, got %s", value.String())
	}
	if value.Cmp(minQuantum) < 0 {
		return Quantum{}, fmt.Errorf("quantum must be >= 1e-15, got %s", value.String())
	}
	return Quantum{value: value}, nil
}

// DeprecatedMustNewQuantum creates a new Quantum from a Decimal value.
// Panics if value is invalid. Use for testing and known-valid constants.
func DeprecatedMustNewQuantum(value Decimal) Quantum {
	q, err := DeprecatedNewQuantum(value)
	if err != nil {
		panic(err)
	}
	return q
}

// NewQuantum creates a new Quantum from a string representation.
// Returns error if the string is not a valid decimal or if quantum constraints are violated.
func NewQuantum(s string) (Quantum, error) {
	dec, err := NewDecimal(s)
	if err != nil {
		return Quantum{}, fmt.Errorf("invalid decimal: %w", err)
	}
	return DeprecatedNewQuantum(dec)
}

// MustNewQuantum creates a new Quantum from a string representation.
// Panics if the string is invalid. Use for testing and known-valid constants.
func MustNewQuantum(s string) Quantum {
	q, err := NewQuantum(s)
	if err != nil {
		panic(err)
	}
	return q
}

// Decimal returns the quantum value as a Decimal.
func (q Quantum) Decimal() Decimal {
	return q.value
}

// String returns the string representation of the quantum.
func (q Quantum) String() string {
	return q.value.String()
}

// Equal returns true if two quanta have the same value.
// Note: This only compares values, not dimension types (handled by Go's type system).
func (q Quantum) Equal(other Quantum) bool {
	return q.value.Equal(other.value)
}

// Common quantum constructors for convenience

// NewQuantum001 creates a quantum of 0.01 (e.g., USD cents, EUR cents).
func NewQuantum001() Quantum {
	return MustNewQuantum("0.01")
}

// NewQuantum1 creates a quantum of 1 (e.g., JPY yen, whole units).
func NewQuantum1() Quantum {
	return MustNewQuantum("1")
}

// NewQuantum0001 creates a quantum of 0.001 (e.g., KWD fils).
func NewQuantum0001() Quantum {
	return MustNewQuantum("0.001")
}
