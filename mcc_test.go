package mcc

import (
	"errors"
	"strings"
	"testing"
)

func TestGetCategory(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		expected string
		err      bool
	}{
		{"Valid code - 0742", "0742", "Veterinary services", false},
		{"Valid code - 0743", "0743", "Wine producers", false},
		{"Valid code - 0744", "0744", "Champagne producers", false},
		{"Valid code - 0763", "0763", "Agricultural co-operatives", false},
		{"Valid code - 5411", "5411", "Groceries and supermarkets", false},
		{"Valid code - 5262", "5262", "Garden supply stores", false},
		{"Non-existent code", "9999", "", true},
		{"Empty code", "", "", true},
		{"Invalid format - letters", "abc", "", true},
		{"Invalid format - too short", "123", "", true},
		{"Invalid format - too long", "12345", "", true},
		{"Invalid format - alphanumeric", "12ab", "", true},
		{"Whitespace padding", "  5411  ", "Groceries and supermarkets", false},
		{"Leading zeros", "0742", "Veterinary services", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetCategory(tt.code)
			if tt.err {
				if err == nil {
					t.Errorf("GetCategory(%q) expected error, got nil", tt.code)
				}
				if !errors.Is(err, ErrNotFound) {
					t.Errorf("GetCategory(%q) expected ErrNotFound, got %v", tt.code, err)
				}
			} else {
				if err != nil {
					t.Errorf("GetCategory(%q) unexpected error: %v", tt.code, err)
				}
				if result != tt.expected {
					t.Errorf("GetCategory(%q) = %q, want %q", tt.code, result, tt.expected)
				}
			}
		})
	}
}

func TestGetCategoryWithCode(t *testing.T) {
	tests := []struct {
		name     string
		code     string
		expected Category
		err      bool
	}{
		{"Valid code", "5411", Category{Code: "5411", Description: "Groceries and supermarkets"}, false},
		{"Non-existent code", "9999", Category{}, true},
		{"Empty code", "", Category{}, true},
		{"Invalid format", "abc", Category{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := GetCategoryWithCode(tt.code)
			if tt.err {
				if err == nil {
					t.Errorf("GetCategoryWithCode(%q) expected error, got nil", tt.code)
				}
			} else {
				if err != nil {
					t.Errorf("GetCategoryWithCode(%q) unexpected error: %v", tt.code, err)
				}
				if result.Code != tt.expected.Code || result.Description != tt.expected.Description {
					t.Errorf("GetCategoryWithCode(%q) = %+v, want %+v", tt.code, result, tt.expected)
				}
			}
		})
	}
}

func TestGetAllCategories(t *testing.T) {
	categories := GetAllCategories()

	// Test that we have some categories
	if len(categories) == 0 {
		t.Error("GetAllCategories() returned empty map")
	}

	// Test a few known categories
	knownCodes := []string{"0742", "0743", "0744", "0763", "5411", "5262"}
	for _, code := range knownCodes {
		if _, exists := categories[code]; !exists {
			t.Errorf("GetAllCategories() missing known code %q", code)
		}
	}

	// Test that the returned map is a copy (modifications shouldn't affect internal data)
	categories["TEST"] = Category{Code: "TEST", Description: "Test"}
	allCategoriesAgain := GetAllCategories()
	if _, exists := allCategoriesAgain["TEST"]; exists {
		t.Error("GetAllCategories() returned map is not a copy - modification affected internal data")
	}
}

func TestGetAllCategoriesIsolation(t *testing.T) {
	// Verify that modifying the returned map doesn't affect the original
	categories1 := GetAllCategories()
	categories1["9999"] = Category{Code: "9999", Description: "Test"}

	categories2 := GetAllCategories()
	if _, exists := categories2["9999"]; exists {
		t.Error("Modifying returned map affected internal data - map is not isolated")
	}
}

func TestErrorMessages(t *testing.T) {
	tests := []struct {
		name           string
		code           string
		expectedInErr  string
	}{
		{"Empty code error", "", "empty code"},
		{"Invalid format error", "abc", "invalid format"},
		{"Not found error", "9999", "9999"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetCategory(tt.code)
			if err == nil {
				t.Errorf("Expected error for code %q", tt.code)
				return
			}
			if !strings.Contains(err.Error(), tt.expectedInErr) {
				t.Errorf("Error message should contain %q, got %q", tt.expectedInErr, err.Error())
			}
		})
	}
}
