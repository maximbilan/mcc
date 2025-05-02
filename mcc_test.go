package mcc

import (
	"testing"
)

func TestGetCategory(t *testing.T) {
	tests := []struct {
		code     string
		expected string
		err      bool
	}{
		{"0742", "Veterinary services", false},
		{"0743", "Wine producers", false},
		{"0744", "Champagne producers", false},
		{"0763", "Agricultural co-operatives", false},
		{"9999", "", true}, // Non-existent code
		{"", "", true},     // Empty code
		{"abc", "", true},  // Invalid code
	}

	for _, test := range tests {
		result, err := GetCategory(test.code)
		if test.err {
			if err == nil {
				t.Errorf("GetCategory(%q) expected error, got nil", test.code)
			}
		} else {
			if err != nil {
				t.Errorf("GetCategory(%q) unexpected error: %v", test.code, err)
			}
			if result != test.expected {
				t.Errorf("GetCategory(%q) = %q, want %q", test.code, result, test.expected)
			}
		}
	}
}

func TestGetAllCategories(t *testing.T) {
	categories := GetAllCategories()

	// Test that we have some categories
	if len(categories) == 0 {
		t.Error("GetAllCategories() returned empty map")
	}

	// Test a few known categories
	knownCodes := []string{"0742", "0743", "0744", "0763"}
	for _, code := range knownCodes {
		if _, exists := categories[code]; !exists {
			t.Errorf("GetAllCategories() missing known code %q", code)
		}
	}
}
