# Go Nullable

Nullable is specifically designed for json marshalling and unmarshalling but not limited to, where zero values have business logic meaning.

## Examples

```go
package main

import (
	"fmt"

	"github.com/steelshot/go-nullable"
)

func main() {
	var greet nullable.Any[string]

	if greet.Null() {
		fmt.Println("greet is null after var declaration")
	}

	greet = nullable.Of("")

	if greet.Null() {
		fmt.Println("greet is null after zero value assignment")
	}

	greet = nullable.Of("Hello, World!")

	if greet.Null() {
		fmt.Println("greet is null after value assignment")
	}

	fmt.Printf(`greet value is "%s"`, greet)

	// Output:
	// greet is null after var declaration
	// greet value is "Hello, World!"
}
```

## Roadmap

- [x] Aliases for builtin types
- [ ] Text & Binary Marshalling
- [ ] Support YAML, TOML and others without 3rd party dependencies
- [ ] API freeze & release 1.0.0

## FAQ

#### Which GoLang version do I need?

As of the current release, the minimum required Go version is [1.18](https://go.dev/doc/go1.18) which introduced Go generics

#### What is planned for 1.0.0?

As of the current release Go generics are in their infancy age.
As Go generics grow and become more streamlined, features will be added, modified or removed if there is benefit to the module.
Breaking changes before v1.0.0 are expected.

## License

[MIT](https://choosealicense.com/licenses/mit/)
