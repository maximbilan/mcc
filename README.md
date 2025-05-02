# MCC Categorizer

A lightweight Go framework for categorizing Merchant Category Codes (MCC) into human-readable English categories.

## Features

- Simple and efficient MCC code categorization
- Comprehensive coverage of standard MCC codes
- Easy-to-use API
- Well-documented code
- MIT licensed

## Installation

```bash
go get github.com/maximbilan/mcc
```

## Usage

```go
package main

import (
    "fmt"
    "github.com/maximbilan/mcc"
)

func main() {
    // Get category for a specific MCC code
    category, err := mcc.GetCategory("0742") // Veterinary Services
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    fmt.Println(category)
}
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the MIT License - see the LICENSE file for details. 