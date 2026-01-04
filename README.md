# MCC Categorizer

A lightweight Go library for categorizing Merchant Category Codes (MCC) into human-readable English categories based on the ISO 18245 standard.

## Features

- ✅ Simple and efficient MCC code categorization
- ✅ Comprehensive coverage of standard MCC codes from ISO 18245 standard database
- ✅ Input validation (ensures 4-digit numeric codes)
- ✅ Easy-to-use API with comprehensive error handling
- ✅ Well-documented code with examples
- ✅ 100% test coverage
- ✅ MIT licensed

## Installation

```bash
go get github.com/maximbilan/mcc
```

## Usage

### Basic Usage

Get the category description for a specific MCC code:

```go
package main

import (
    "fmt"
    "github.com/maximbilan/mcc"
)

func main() {
    category, err := mcc.GetCategory("5411")
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Println(category) // Output: Groceries and supermarkets
}
```

### Get Full Category Information

Get both the code and description using `GetCategoryWithCode()`:

```go
category, err := mcc.GetCategoryWithCode("5262")
if err != nil {
    fmt.Printf("Error: %v\n", err)
    return
}
fmt.Printf("Code: %s, Description: %s\n", category.Code, category.Description)
// Output: Code: 5262, Description: Garden supply stores
```

### Error Handling

The library provides clear error messages and uses `ErrNotFound` for consistent error checking:

```go
import (
    "errors"
    "fmt"
    "github.com/maximbilan/mcc"
)

category, err := mcc.GetCategory("9999")
if err != nil {
    if errors.Is(err, mcc.ErrNotFound) {
        fmt.Println("MCC code not found")
    } else {
        fmt.Printf("Error: %v\n", err)
    }
}
```

### Get All Categories

Retrieve all available MCC codes and their descriptions:

```go
allCategories := mcc.GetAllCategories()
for code, category := range allCategories {
    fmt.Printf("%s: %s\n", code, category.Description)
}
```

### Input Validation

The library automatically validates MCC codes:
- Must be exactly 4 digits
- Whitespace is automatically trimmed
- Invalid formats return descriptive error messages

```go
// These will return errors:
mcc.GetCategory("")        // Empty code
mcc.GetCategory("abc")     // Invalid format
mcc.GetCategory("123")     // Too short
mcc.GetCategory("12345")   // Too long

// This works (whitespace is trimmed):
mcc.GetCategory("  5411  ") // Valid
```

## API Reference

### Functions

#### `GetCategory(code string) (string, error)`

Returns the category description for a given MCC code. The code is normalized by removing whitespace and validated before lookup.

**Parameters:**
- `code`: A 4-digit MCC code string

**Returns:**
- `string`: The category description
- `error`: Error if code is invalid or not found

#### `GetCategoryWithCode(code string) (Category, error)`

Returns the full `Category` struct containing both the code and description.

**Parameters:**
- `code`: A 4-digit MCC code string

**Returns:**
- `Category`: The category with code and description
- `error`: Error if code is invalid or not found

#### `GetAllCategories() map[string]Category`

Returns a copy of all available MCC codes and their descriptions. The returned map is a copy to prevent external modification.

**Returns:**
- `map[string]Category`: A map of all MCC codes to their categories

### Types

#### `Category`

```go
type Category struct {
    Code        string
    Description string
}
```

Represents a merchant category with its code and description.

### Errors

#### `ErrNotFound`

A sentinel error returned when an MCC code is not found in the database. Use `errors.Is()` to check for this error:

```go
if errors.Is(err, mcc.ErrNotFound) {
    // Handle not found error
}
```

## Examples

See the [example directory](./example/main.go) for more complete examples.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details. 