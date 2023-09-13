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

package nullables

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
)

type (
	// Json is *essentially* an alias to nullable.Nullable with added json marshalling/unmarshalling
	// NOTE: Since golang has no type constraints for json compatible types, if you use a json
	// incompatible type i.e. complex64, json marshalling will always throw an error.
	// This is due to internally using json.Marshal / json.Unmarshal on the underlying types
	Json[T any] struct {
		internal[T]
	}

	// Yaml is *essentially* an alias to nullable.Nullable with added yaml marshalling/unmarshalling
	// NOTE: Since golang has no type constraints for yaml compatible types, if you use a yaml
	// incompatible type i.e. complex64, yaml marshalling will always throw an error.
	// This is due to internally using yaml.Marshal / yaml.Unmarshal on the underlying types
	Yaml[T any] struct {
		internal[T]
	}
)

// JsonOf returns a default Json[T] based on the type of value param.
func JsonOf[T any](value T) Json[T] {
	return Json[T]{of(value)}
}

// YamlOf returns a default YamlOf[T] based on the type of value param.
func YamlOf[T any](value T) Yaml[T] {
	return Yaml[T]{of(value)}
}

// MarshalJSON implements json.Marshaler
func (r Json[T]) MarshalJSON() ([]byte, error) {
	if r.Null() {
		return []byte("null"), nil
	}

	return json.Marshal(r.Value())
}

// UnmarshalJSON implements json.Unmarshaler
func (r *Json[T]) UnmarshalJSON(bytes []byte) (err error) {
	var newT *T

	if err = json.Unmarshal(bytes, &newT); err != nil {
		return
	}

	if newT != nil {
		*r = JsonOf(*newT)
	}

	return
}

// IsZero implements yaml.IsZeroer
func (r Yaml[T]) IsZero() bool {
	return r.Null()
}

// MarshalYAML implements yaml.Marshaler
func (r Yaml[T]) MarshalYAML() (interface{}, error) {
	if r.Null() {
		return nil, nil
	}

	return r.Value(), nil
}

// UnmarshalYAML implements yaml.Unmarshaler
func (r *Yaml[T]) UnmarshalYAML(node *yaml.Node) (err error) {
	var newT *T

	if err = node.Decode(&newT); err != nil {
		return
	}

	if newT != nil {
		*r = YamlOf(*newT)
	}

	return
}
