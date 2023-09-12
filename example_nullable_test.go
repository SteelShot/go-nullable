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

package nullable_test

import (
	"fmt"

	"github.com/SteelShot/go-nullable"
)

type (
	Person struct {
		Name string
		Age  uint8
	}

	Card struct {
		ID     string
		Person nullable.Nullable[Person]
	}
)

func Example_nullableString() {
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

func Example_nullableStruct() {
	var person nullable.Nullable[Person]

	if person.Null() {
		fmt.Println("person declaration is null")
	}

	person = nullable.Of(Person{
		Name: "John Doe",
		Age:  18,
	})

	if person.Null() {
		fmt.Println("Of assignment is null")
	}

	fmt.Println("value is", person.Value())

	// Output:
	// person declaration is null
	// value is {John Doe 18}
}

func Example_nullableStructField() {
	var card nullable.Nullable[Card]

	if card.Null() {
		fmt.Println("card declaration is null")
	}

	card = nullable.Of(Card{
		ID:     "card-id",
		Person: nullable.Nullable[Person]{},
	})

	if card.Null() {
		fmt.Println("Of assignment is null")
	}

	if card.Value().Person.Null() {
		fmt.Println("card.Person is null")
	}

	fmt.Println("value is", card.Value())

	// Output:
	// card declaration is null
	// card.Person is null
	// value is {card-id nil}
}
