package quanta

import (
	"testing"
)

// TestValidateWeights tests the weight validation helper.
func TestValidateWeights(t *testing.T) {
	tests := []struct {
		name    string
		weights []int64
		wantErr bool
		errMsg  string
	}{
		{
			name:    "empty weights",
			weights: []int64{},
			wantErr: true,
			errMsg:  "weights cannot be empty",
		},
		{
			name:    "negative weight",
			weights: []int64{5, -3, 2},
			wantErr: true,
			errMsg:  "weight[1] is negative",
		},
		{
			name:    "all zero weights",
			weights: []int64{0, 0, 0},
			wantErr: true,
			errMsg:  "all weights are zero",
		},
		{
			name:    "single positive weight",
			weights: []int64{10},
			wantErr: false,
		},
		{
			name:    "multiple positive weights",
			weights: []int64{1, 1, 1},
			wantErr: false,
		},
		{
			name:    "unequal positive weights",
			weights: []int64{5, 3, 2},
			wantErr: false,
		},
		{
			name:    "some zero weights mixed with positive",
			weights: []int64{5, 0, 3},
			wantErr: false,
		},
		{
			name:    "large weights",
			weights: []int64{1000000, 2000000, 3000000},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateWeights(tt.weights)
			if tt.wantErr {
				if err == nil {
					t.Errorf("validateWeights() expected error containing %q, got nil", tt.errMsg)
					return
				}
				if tt.errMsg != "" && !contains(err.Error(), tt.errMsg) {
					t.Errorf("validateWeights() error = %v, want substring %q", err, tt.errMsg)
				}
			} else {
				if err != nil {
					t.Errorf("validateWeights() unexpected error: %v", err)
				}
			}
		})
	}
}

// TestAllocationResult_PartsLength tests that AllocationResult.Parts has correct length.
// (Actual strategy implementations will be tested in their own files)
func TestAllocationResult_Structure(t *testing.T) {
	// This is a structural test to ensure AllocationResult can be instantiated
	spec := MustNewUnit("USD", "0.01")
	parts := []Quantized{
		DeprecatedNewQuantizedFromMultiplier(spec, 333),
		DeprecatedNewQuantizedFromMultiplier(spec, 333),
		DeprecatedNewQuantizedFromMultiplier(spec, 334),
	}
	dust := NewMeasureFrom(spec, MustNewDecimal("0"))

	result := AllocationResult{
		Parts: parts,
		Dust:  dust,
	}

	if len(result.Parts) != 3 {
		t.Errorf("AllocationResult.Parts length = %d, want 3", len(result.Parts))
	}

	// Verify dust is zero value
	if !result.Dust.Quantity().IsZero() {
		t.Errorf("AllocationResult.Dust = %v, want zero", result.Dust.Quantity())
	}
}

// contains checks if a string contains a substring.
func contains(s, substr string) bool {
	return len(s) >= len(substr) && (s == substr || len(substr) == 0 ||
		(len(s) > 0 && len(substr) > 0 && hasSubstring(s, substr)))
}

func hasSubstring(s, substr string) bool {
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return true
		}
	}
	return false
}
