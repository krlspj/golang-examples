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
	Value   string `json:",omitempty"`
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
	Name    NullString `json:"name,omitempty"`
	Age     NullInt    `json:"age,omitempty"`
	Height  NullInt    `json:"height,omitempty"`
	Married NullBool   `json:"married,omitempty"`
	Pet     NullString `json:"pet,omitempty"`
	Job     NullString `json:"job,omitempty"`
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

func (ns NullString) MarshalJSON() ([]byte, error) {
	if ns.IsNull {
		return []byte("null"), nil
	}
	if ns.IsEmpty {
		return []byte("{}"), nil
	}
	return json.Marshal(ns.Value)
}

func (ns NullBool) MarshalJSON() ([]byte, error) {
	if ns.IsNull {
		return []byte("null"), nil
	}
	if ns.IsEmpty {
		return []byte("{}"), nil
	}
	return json.Marshal(ns.Value)
}

func (ns NullInt) MarshalJSON() ([]byte, error) {
	if ns.IsNull {
		return []byte("null"), nil
	}
	if ns.IsEmpty {
		return []byte("{}"), nil
	}
	return json.Marshal(ns.Value)
}

func (p Person) MarshalJSON() ([]byte, error) {
	var name []byte
	if p.Name.IsEmpty {
		//passs
	} else if p.Name.IsNull {
		name = []byte("null")
	} else {
		name = []byte(p.Name.Value)
	}
	return json.Marshal(struct {
		Name string `json:"name,omitempty"`
	}{
		Name: string(name),
	})

}

func main() {
	person := NewPerson()
	err := json.Unmarshal([]byte(simpleJsonData), &person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("Person: %+v\n", person)

	fmt.Println("----- unmarshalled-----")

	person1 := Person{Name: NullString{Value: "John Doe", IsNull: false, IsEmpty: false}, Age: NullInt{Value: 0, IsNull: false, IsEmpty: false}, Height: NullInt{Value: 0, IsNull: true, IsEmpty: false}, Married: NullBool{Value: false, IsNull: false, IsEmpty: false}, Pet: NullString{Value: "missing", IsNull: false, IsEmpty: false}, Job: NullString{Value: "", IsNull: false, IsEmpty: true}}

	fmt.Println("------ marshalling ------")
	data, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("err", err.Error())
		return
	}
	fmt.Println("data ->", string(data))

}
