# Go Nullable

Nullable generic go types using the power of golang generics

## Examples

```go
package main

// https://go.dev/play/p/Clck0xmCprK

import (
	"fmt"

	"github.com/SteelShot/go-nullable"
)

func main() {
	var greet nullable.Nullable[string]

	if greet.Null() {
		fmt.Println("greet declaration is null")
	}

	greet = nullable.Of("Hello, World!")

	if greet.Null() {
		fmt.Println("greet assignment is null")
	}

	fmt.Println("value is", greet.Value())

	// Output:
	// greet declaration is null
	// value is Hello, World!
}
```

## FAQ

#### Which GoLang version do I need?

As of the current release, the minimum required Go version is [1.18](https://go.dev/doc/go1.18) which introduced Go generics

#### What is planned for 1.0.0?

As of the current release Go generics are in their infancy age.
As Go generics grow and become more streamlined, features will be added, modified or removed if there is benefit to the module.
Breaking changes before 1.0.0 are expected.
Currently, there is no roadmap.

## License

[MIT](https://choosealicense.com/licenses/mit/)
