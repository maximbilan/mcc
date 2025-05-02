package main

import (
	"fmt"

	"github.com/maximbilan/mcc"
)

func main() {
	// Example 1: Get a single category
	code := "5411"
	category, err := mcc.GetCategory(code)
	if err != nil {
		fmt.Printf("Error getting category for code %s: %v\n", code, err)
	} else {
		fmt.Printf("MCC %s: %s\n", code, category)
	}

	// Example 2: Get all categories
	fmt.Println("\nAll available categories:")
	allCategories := mcc.GetAllCategories()
	for code, category := range allCategories {
		fmt.Printf("%s: %s\n", code, category.Description)
	}
}
