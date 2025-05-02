package mcc

import (
	"errors"
	"strings"
)

// Category represents a merchant category with its code and description
type Category struct {
	Code        string
	Description string
}

// Returns the category description for a given MCC code
func GetCategory(code string) (string, error) {
	// Normalize the code by removing any whitespace
	code = strings.TrimSpace(code)

	// Check if the code exists in our map
	if category, exists := mccData[code]; exists {
		return category.Description, nil
	}

	return "", errors.New("MCC code not found")
}

// Returns a map of all available MCC codes and their descriptions
func GetAllCategories() map[string]Category {
	return mccData
}
