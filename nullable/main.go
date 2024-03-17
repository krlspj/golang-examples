package main

import (
	"encoding/json"
	"fmt"
)

const simpleJsonData = `{
  "name": "John Doe",
  "age": 0,
  "height": null,
  "married": false,
  "pet": ""
}`

type NullInt struct {
	Value   int
	IsNull  bool
	IsEmpty bool
}

type NullString struct {
	Value   string
	IsNull  bool
	IsEmpty bool
}

type NullBool struct {
	Value   bool
	IsNull  bool
	IsEmpty bool
}

func NewEmptyNullInt() NullInt {
	return NullInt{IsNull: false, IsEmpty: true}
}

func NewEmptyNullString() NullString {
	return NullString{IsNull: false, IsEmpty: true}
}

func NewEmptyNullBool() NullBool {
	return NullBool{IsNull: false, IsEmpty: true}
}

type Person struct {
	Name    NullString `json:"name"`
	Age     NullInt    `json:"age"`
	Height  NullInt    `json:"height"`
	Married NullBool   `json:"married"`
	Pet     NullString `json:"pet"`
	Job     NullString `json:"job"`
}

func NewPerson() Person {
	return Person{
		Name:    NewEmptyNullString(),
		Age:     NewEmptyNullInt(),
		Height:  NewEmptyNullInt(),
		Married: NewEmptyNullBool(),
		Pet:     NewEmptyNullString(),
		Job:     NewEmptyNullString(),
	}
}

func (ni *NullInt) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		ni.IsNull = true
		ni.IsEmpty = false
		return nil
	}
	ni.IsNull = false
	ni.IsEmpty = false
	return json.Unmarshal(data, &ni.Value)
}

func (ns *NullString) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		ns.IsNull = true
		ns.IsEmpty = false
		return nil
	}
	ns.IsNull = false
	ns.IsEmpty = false
	return json.Unmarshal(data, &ns.Value)
}

func (nb *NullBool) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		nb.IsNull = true
		nb.IsEmpty = false
		return nil
	}
	nb.IsNull = false
	nb.IsEmpty = false
	return json.Unmarshal(data, &nb.Value)
}

func main() {
	person := NewPerson()
	err := json.Unmarshal([]byte(simpleJsonData), &person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Person: %+v\n", person)
}
