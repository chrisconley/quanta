package quanta

import "github.com/cockroachdb/apd/v3"

// Rounding represents different rounding modes for decimal operations.
// These modes control how values are rounded when quantizing to a specific precision.
type Rounding int

const (
	// RoundHalfEven rounds to nearest, ties to even (banker's rounding).
	// This is the default rounding mode as it's statistically unbiased.
	// Examples: 1.5 → 2, 2.5 → 2, 1.235 → 1.24, 1.225 → 1.22
	RoundHalfEven Rounding = iota

	// RoundHalfUp rounds to nearest, ties away from zero.
	// Examples: 1.5 → 2, 2.5 → 3, -1.5 → -2
	RoundHalfUp

	// RoundHalfDown rounds to nearest, ties toward zero.
	// Examples: 1.5 → 1, 2.5 → 2, -1.5 → -1
	RoundHalfDown

	// RoundUp rounds away from zero (ceiling for positive, floor for negative).
	// Examples: 1.1 → 2, -1.1 → -2
	RoundUp

	// RoundDown rounds toward zero (floor for positive, ceiling for negative).
	// Examples: 1.9 → 1, -1.9 → -1
	RoundDown

	// RoundCeiling rounds toward positive infinity.
	// Examples: 1.1 → 2, -1.9 → -1
	RoundCeiling

	// RoundFloor rounds toward negative infinity.
	// Examples: 1.9 → 1, -1.1 → -2
	RoundFloor
)

// toAPDRounder converts the Rounding enum to apd.Rounder.
// This is an internal conversion - apd types are not exposed in the public API.
func (r Rounding) toAPDRounder() apd.Rounder {
	switch r {
	case RoundHalfUp:
		return apd.RoundHalfUp
	case RoundHalfDown:
		return apd.RoundHalfDown
	case RoundUp:
		return apd.RoundUp
	case RoundDown:
		return apd.RoundDown
	case RoundCeiling:
		return apd.RoundCeiling
	case RoundFloor:
		return apd.RoundFloor
	case RoundHalfEven:
		return apd.RoundHalfEven
	default:
		// Default to banker's rounding
		return apd.RoundHalfEven
	}
}

// String returns the string representation of the rounding mode.
func (r Rounding) String() string {
	switch r {
	case RoundHalfEven:
		return "half-even"
	case RoundHalfUp:
		return "half-up"
	case RoundHalfDown:
		return "half-down"
	case RoundUp:
		return "up"
	case RoundDown:
		return "down"
	case RoundCeiling:
		return "ceiling"
	case RoundFloor:
		return "floor"
	default:
		return "unknown"
	}
}
