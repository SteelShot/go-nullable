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

package nullable

import (
	"encoding/json"
	"fmt"
	"testing"
)

const (
	Null = "null"
	Nil  = "<nil>"

	PersonName = "John Doe"
	PersonAge  = uint8(18)
)

type Person struct {
	Name String `json:"name,omitempty"`
	Age  Uint8  `json:"age,omitempty"`
}

func TestNullable_Null(t *testing.T) {
	n := Any[struct{}]{}

	assertEquals(t, true, n.Null())

	n = Of(struct{}{})

	assertEquals(t, false, n.Null())
}

func TestNullable_Value(t *testing.T) {
	n := Any[string]{}

	assertEquals(t, true, n.Null())

	n = Of("")

	assertEquals(t, false, n.Null())

	n = Of(Null)

	assertEquals(t, Null, n.Value())
}

func TestNullable_ValuePointer(t *testing.T) {
	n := Any[*string]{}

	assertEquals(t, true, n.Null())

	n = Of(new(string))

	assertEquals(t, false, n.Null())

	ptr := new(string)

	n = Of(ptr)

	assertEquals(t, ptr, n.Value())

	*ptr = Null

	assertEquals(t, *ptr, *n.Value())
}

func TestNullable_ValuePointerOfPointer(t *testing.T) {
	n := Any[**string]{}

	assertEquals(t, true, n.Null())

	n = Of(new(*string))

	assertEquals(t, false, n.Null())

	ptr := new(*string)

	n = Of(ptr)

	assertEquals(t, ptr, n.Value())

	*ptr = new(string)

	assertEquals(t, *ptr, *n.Value())

	**ptr = Null

	assertEquals(t, Null, **n.Value())
}

func TestNullable_Stringer(t *testing.T) {
	nullableFmtStringersTest(t, Null, Any[string].String)
}

func TestNullable_StringerPointer(t *testing.T) {
	s := Null

	nullableFmtStringersTest(t, &s, Any[*string].String)
}

func Test_SimpleJsonUnmarshalMarshal(t *testing.T) {
	var (
		person        Person
		RawJsonPerson = fmt.Sprintf(`{"name":"%s","age":%d}`, PersonName, PersonAge)
	)

	if err := json.Unmarshal([]byte(RawJsonPerson), &person); err != nil {
		t.Fatal(err)
	}

	if person.Name.Null() || person.Name.Value() != PersonName {
		t.FailNow()
	}

	if person.Age.Null() || person.Age.Value() != PersonAge {
		t.FailNow()
	}

	if bytes, err := json.Marshal(person); err != nil {
		t.Fatal(err)
	} else if result := string(bytes); result != RawJsonPerson {
		t.Fatal(result, "!=", RawJsonPerson)
	}
}

// reusable tests & utilities
func nullableFmtStringersTest[T any](t *testing.T, expected T, f func(nullable Any[T]) string) {
	n := Any[T]{}

	assertEquals(t, Nil, f(n))

	n = Of(*new(T))

	assertEquals(t, fmt.Sprint(*new(T)), f(n))

	n = Of(expected)

	assertEquals(t, fmt.Sprint(expected), f(n))
}

func assertEquals[T comparable](t *testing.T, expected, value T) {
	assertEqualsMsg(t, expected, value, "")
}

func assertEqualsMsg[T comparable](t *testing.T, expected, value T, message string) {
	if expected != value {
		if message == "" {
			t.Fatalf("expected: %v, got: %v", expected, value)
		}

		t.Fatalf("%s: expected: %v, got: %v", message, expected, value)
	}
}
