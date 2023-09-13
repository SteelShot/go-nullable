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
	"strings"
	"testing"
)

const (
	PersonName    = "John Doe"
	RawJsonPerson = `{"name":"John Doe","age":null}`
	RawYamlPerson = `name: John Doe`
)

type (
	JsonPerson struct {
		Name Json[string] `json:"name,omitempty"`
		Age  Json[uint8]  `json:"age,omitempty"`
	}

	YamlPerson struct {
		Name Yaml[string] `yaml:"name,omitempty"`
		Age  Yaml[uint8]  `yaml:"age,omitempty"`
	}
)

func TestJson_MarshalAndUnmarshal(t *testing.T) {
	var jsonPerson JsonPerson

	if err := json.Unmarshal([]byte(RawJsonPerson), &jsonPerson); err != nil {
		t.Fatal(err)
	}

	if jsonPerson.Name.Null() || jsonPerson.Name.Value() != PersonName {
		t.FailNow()
	}

	if bytes, err := json.Marshal(jsonPerson); err != nil {
		t.Fatal(err)
	} else if string(bytes) != RawJsonPerson {
		t.Fatal(string(bytes), "!=", RawJsonPerson)
	}
}

func TestYaml_MarshalAndUnmarshal(t *testing.T) {
	var yamlPerson YamlPerson

	if err := yaml.Unmarshal([]byte(RawYamlPerson), &yamlPerson); err != nil {
		t.Fatal(err)
	}

	if yamlPerson.Name.Null() || yamlPerson.Name.Value() != PersonName {
		t.FailNow()
	}

	if bytes, err := yaml.Marshal(yamlPerson); err != nil {
		t.Fatal(err)
	} else if result := strings.TrimSpace(string(bytes)); result != RawYamlPerson {
		t.Fatal(result, "!=", RawYamlPerson)
	}
}
