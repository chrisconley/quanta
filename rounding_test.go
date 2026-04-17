package quanta

import (
	"testing"

	"github.com/cockroachdb/apd/v3"
	"github.com/stretchr/testify/assert"
)

func TestRounding_toAPDRounder(t *testing.T) {
	t.Run("converts RoundHalfEven", func(t *testing.T) {
		r := RoundHalfEven
		apdr := r.toAPDRounder()
		assert.Equal(t, apd.RoundHalfEven, apdr)
	})

	t.Run("converts RoundHalfUp", func(t *testing.T) {
		r := RoundHalfUp
		apdr := r.toAPDRounder()
		assert.Equal(t, apd.RoundHalfUp, apdr)
	})

	t.Run("converts RoundHalfDown", func(t *testing.T) {
		r := RoundHalfDown
		apdr := r.toAPDRounder()
		assert.Equal(t, apd.RoundHalfDown, apdr)
	})

	t.Run("converts RoundUp", func(t *testing.T) {
		r := RoundUp
		apdr := r.toAPDRounder()
		assert.Equal(t, apd.RoundUp, apdr)
	})

	t.Run("converts RoundDown", func(t *testing.T) {
		r := RoundDown
		apdr := r.toAPDRounder()
		assert.Equal(t, apd.RoundDown, apdr)
	})

	t.Run("converts RoundCeiling", func(t *testing.T) {
		r := RoundCeiling
		apdr := r.toAPDRounder()
		assert.Equal(t, apd.RoundCeiling, apdr)
	})

	t.Run("converts RoundFloor", func(t *testing.T) {
		r := RoundFloor
		apdr := r.toAPDRounder()
		assert.Equal(t, apd.RoundFloor, apdr)
	})

	t.Run("defaults to RoundHalfEven for unknown value", func(t *testing.T) {
		r := Rounding(999) // Invalid value
		apdr := r.toAPDRounder()
		assert.Equal(t, apd.RoundHalfEven, apdr)
	})
}

func TestRounding_String(t *testing.T) {
	tests := []struct {
		rounding Rounding
		expected string
	}{
		{RoundHalfEven, "half-even"},
		{RoundHalfUp, "half-up"},
		{RoundHalfDown, "half-down"},
		{RoundUp, "up"},
		{RoundDown, "down"},
		{RoundCeiling, "ceiling"},
		{RoundFloor, "floor"},
	}

	for _, tt := range tests {
		t.Run(tt.expected, func(t *testing.T) {
			assert.Equal(t, tt.expected, tt.rounding.String())
		})
	}

	t.Run("returns unknown for invalid value", func(t *testing.T) {
		r := Rounding(999)
		assert.Equal(t, "unknown", r.String())
	})
}
