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

	"github.com/steelshot/go-nullable"
)

type (
	Person struct {
		Name string
		Age  uint8
	}

	Card struct {
		ID     string
		Person nullable.Any[Person]
	}
)

func Example_nullableString() {
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

func Example_nullableStruct() {
	var person nullable.Any[Person]

	if person.Null() {
		fmt.Println("person is null after var declaration")
	}

	person = nullable.Of(Person{})

	if person.Null() {
		fmt.Println("person is null after zero value assignment")
	}

	person = nullable.Of(Person{
		Name: "John Doe",
		Age:  18,
	})

	if person.Null() {
		fmt.Println("person is null after value assignment")
	}

	fmt.Printf(`person value is "%s"`, person)

	// Output:
	// person is null after var declaration
	// person value is "{John Doe 18}"
}

func Example_nullableNestedStruct() {
	var card nullable.Any[Card]

	if card.Null() {
		fmt.Println("card is null after var declaration")
	}

	if card.Value().Person.Null() {
		fmt.Println("card.Person is null after card var declaration")
	}

	card = nullable.Of(Card{})

	if card.Null() {
		fmt.Println("card is null after zero value declaration")
	}

	if card.Value().Person.Null() {
		fmt.Println("card.Person is null after card zero value declaration")
	}

	card = nullable.Of(Card{
		ID: "card-id",
		Person: nullable.Of(Person{
			Name: "John Doe",
			Age:  18,
		}),
	})

	if card.Null() {
		fmt.Println("card is null after value declaration")
	}

	if card.Value().Person.Null() {
		fmt.Println("card.Person is null after card value declaration")
	}

	fmt.Printf(`card value is "%s"`, card)

	// Output:
	// card is null after var declaration
	// card.Person is null after card var declaration
	// card.Person is null after card zero value declaration
	// card value is "{card-id {John Doe 18}}"
}
