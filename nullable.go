/*
 * Copyright (C) 2023 Džiugas Eiva
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all
 * copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
 * SOFTWARE.
 */

// Package nullable is specifically designed for json marshalling and unmarshalling but not limited to, where zero values have business logic meaning.
package nullable

import (
	"encoding/json"
	"fmt"

	"github.com/steelshot/go-nullable/internal"
)

type (
	// Any holds a single value of type T
	// BUG(golang) Golang currently (if ever?) has no type constraint that indicates that a type is json (un)marshalling compatible.
	// If you use a json incompatible type like complex64, the marshal and unmarshal functions will always fail.
	// BUG(steelshot) Currently the Any type does not implement text/binary (un)marshalling.
	// YAML, TOML and other's (un)marshalling will not work, for the time being, You should implement your own types that embed Any with custom (un)marshalling.
	Any[T any] struct {
		value *T
	}
)

// Of returns a new non-null Any[typeof value]
func Of[T any](value T) Any[T] {
	return Any[T]{&value}
}

// Value returns the value of Any; unless it is null, then the zero value of T will be returned
func (r Any[T]) Value() (value T) {
	if r.value != nil {
		value = *r.value
	}

	return
}

// Null will indicate whether Any is null (without value).
func (r Any[T]) Null() bool {
	return r.value == nil
}

// GoString implements fmt.GoStringer
func (r Any[T]) GoString() string {
	if r.value == nil {
		return "<nil>"
	}

	return fmt.Sprintf("%#v", *r.value)
}

// String implements fmt.Stringer
func (r Any[T]) String() string {
	if r.value == nil {
		return "<nil>"
	}

	return fmt.Sprintf("%v", *r.value)
}

// Format implements fmt.Formatter
func (r Any[T]) Format(f fmt.State, verb rune) {
	if r.value == nil {
		fmt.Fprint(f, "<nil>")
	} else {
		switch verb {
		case 's': // replace s with v since most types don't work with %s
			verb = 'v'
		}

		fmt.Fprintf(f, internal.FormatString(f, verb), *r.value)
	}
}

// (Un)Marshaller

// MarshalJSON implements json.Marshaler
func (r Any[T]) MarshalJSON() ([]byte, error) {
	if r.value == nil {
		return []byte("null"), nil
	}

	return json.Marshal(*r.value)
}

// UnmarshalJSON implements json.Unmarshaler
func (r *Any[T]) UnmarshalJSON(bytes []byte) (err error) {
	return json.Unmarshal(bytes, &r.value)
}
