package main

import (
	"fmt"

	"github.com/maximbilan/mcc"
)

func main() {
	// Example 1: Get a single category description
	code := "5411"
	category, err := mcc.GetCategory(code)
	if err != nil {
		fmt.Printf("Error getting category for code %s: %v\n", code, err)
	} else {
		fmt.Printf("MCC %s: %s\n", code, category)
	}

	// Example 2: Get full category with code and description
	categoryFull, err := mcc.GetCategoryWithCode("5262")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} else {
		fmt.Printf("Full category - Code: %s, Description: %s\n", categoryFull.Code, categoryFull.Description)
	}

	// Example 3: Handle invalid code
	_, err = mcc.GetCategory("9999")
	if err != nil {
		fmt.Printf("Expected error for invalid code: %v\n", err)
	}

	// Example 4: Get all categories (showing first 5)
	fmt.Println("\nFirst 5 available categories:")
	allCategories := mcc.GetAllCategories()
	count := 0
	for code, category := range allCategories {
		if count >= 5 {
			break
		}
		fmt.Printf("%s: %s\n", code, category.Description)
		count++
	}
	fmt.Printf("... and %d more categories\n", len(allCategories)-5)
}
