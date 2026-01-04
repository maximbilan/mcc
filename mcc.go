// Package mcc provides functionality for categorizing Merchant Category Codes (MCC)
// into human-readable English categories based on the ISO 18245 standard.
//
// MCC codes are 4-digit numeric codes used to classify businesses by the type
// of goods or services they provide. This package provides a comprehensive
// mapping of MCC codes to their descriptions.
package mcc

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

var (
	// mccCodePattern validates that an MCC code is exactly 4 digits
	mccCodePattern = regexp.MustCompile(`^\d{4}$`)

	// ErrNotFound is returned when an MCC code is not found in the database
	ErrNotFound = errors.New("MCC code not found")
)

// Category represents a merchant category with its code and description
type Category struct {
	Code        string
	Description string
}

// GetCategory returns the category description for a given MCC code.
// The code is normalized by removing whitespace before lookup.
//
// Example:
//
//	category, err := mcc.GetCategory("5411")
//	if err != nil {
//	    // handle error
//	}
//	fmt.Println(category) // "Groceries and supermarkets"
func GetCategory(code string) (string, error) {
	category, err := GetCategoryWithCode(code)
	if err != nil {
		return "", err
	}
	return category.Description, nil
}

// GetCategoryWithCode returns the full Category struct for a given MCC code.
// This is useful when you need both the code and description.
//
// The code is normalized by removing whitespace and validated to ensure
// it's a 4-digit numeric string before lookup.
//
// Example:
//
//	category, err := mcc.GetCategoryWithCode("5411")
//	if err != nil {
//	    // handle error
//	}
//	fmt.Printf("Code: %s, Description: %s\n", category.Code, category.Description)
func GetCategoryWithCode(code string) (Category, error) {
	// Normalize the code by removing any whitespace
	code = strings.TrimSpace(code)

	// Validate the code format
	if code == "" {
		return Category{}, fmt.Errorf("%w: empty code", ErrNotFound)
	}

	if !mccCodePattern.MatchString(code) {
		return Category{}, fmt.Errorf("%w: invalid format (expected 4 digits, got %q)", ErrNotFound, code)
	}

	// Check if the code exists in our map
	if category, exists := mccData[code]; exists {
		return category, nil
	}

	return Category{}, fmt.Errorf("%w: %s", ErrNotFound, code)
}

// GetAllCategories returns a copy of the map containing all available MCC codes
// and their descriptions. The returned map is a copy to prevent external
// modification of the internal data.
//
// Example:
//
//	allCategories := mcc.GetAllCategories()
//	for code, category := range allCategories {
//	    fmt.Printf("%s: %s\n", code, category.Description)
//	}
func GetAllCategories() map[string]Category {
	// Return a copy to prevent external modification
	result := make(map[string]Category, len(mccData))
	for k, v := range mccData {
		result[k] = v
	}
	return result
}
