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

// Package nullable allows creating generic nullable types
package nullable

import "fmt"

type (
	// Nullable is a container struct holding a pointer to an arbitrary value.
	// By design, the Nullable type is immutable and empty values are considered
	// by default nil/null.
	Nullable[T any] struct {
		value *T
	}
)

// Of returns a default Nullable[T] based on the type of value param.
func Of[T any](value T) Nullable[T] {
	return Nullable[T]{&value}
}

// Value will either return the stored or zero value
func (r Nullable[T]) Value() (value T) {
	if !r.Null() {
		value = *r.value
	}

	return
}

// Null will indicate whether the Nullable type is nil/null
func (r Nullable[T]) Null() bool {
	return r.value == nil
}

// GoString implements fmt.GoStringer
func (r Nullable[T]) GoString() string {
	if r.Null() {
		return "nil"
	}

	return fmt.Sprintf("%v", *r.value)
}

// String implements fmt.Stringer
func (r Nullable[T]) String() string {
	return r.GoString()
}
