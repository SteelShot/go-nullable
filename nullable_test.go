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
	"fmt"
	"testing"
)

const (
	Null = "null"
	Nil  = "nil"
)

func TestNullable_Null(t *testing.T) {
	n := Nullable[struct{}]{}

	if !n.Null() {
		t.Fail()
	}

	n = Of(struct{}{})

	if n.Null() {
		t.Fail()
	}
}

func TestNullable_Value(t *testing.T) {
	n := Nullable[string]{}

	if n.Value() != "" {
		t.Fail()
	}

	n = Of("")

	if n.Value() != "" {
		t.Fail()
	}

	n = Of(Null)

	if n.Value() != Null {
		t.Fail()
	}
}

func TestNullable_ValuePointer(t *testing.T) {
	n := Nullable[*string]{}

	if n.Value() != nil {
		t.Fail()
	}

	n = Of(new(string))

	if n.Value() == nil {
		t.Fail()
	}

	ptr := new(string)

	n = Of(ptr)

	if n.Value() != ptr {
		t.Fail()
	}
}

func TestNullable_ValuePointerOfPointer(t *testing.T) {
	n := Nullable[**string]{}

	if n.Value() != nil {
		t.Fail()
	}

	n = Of(new(*string))

	if n.Value() == nil {
		t.Fail()
	}

	ptr := new(*string)

	n = Of(ptr)

	if n.Value() != ptr {
		t.Fail()
	}
}

func TestNullable_GoString(t *testing.T) {
	nullableFmtStringersTest(t, Null, Nullable[string].GoString)
}

func TestNullable_GoStringPointer(t *testing.T) {
	s := Null

	nullableFmtStringersTest(t, &s, Nullable[*string].GoString)
}

func TestNullable_String(t *testing.T) {
	nullableFmtStringersTest(t, Null, Nullable[string].String)
}

func TestNullable_StringPointer(t *testing.T) {
	s := Null

	nullableFmtStringersTest(t, &s, Nullable[*string].String)
}

// reusable tests
func nullableFmtStringersTest[T any](t *testing.T, expected T, f func(nullable Nullable[T]) string) {
	n := Nullable[T]{}

	if f(n) != Nil {
		t.Fail()
	}

	n = Of(*new(T))

	if f(n) != fmt.Sprint(*new(T)) {
		t.Fail()
	}

	n = Of(expected)

	if f(n) != fmt.Sprint(expected) {
		t.Fail()
	}
}
