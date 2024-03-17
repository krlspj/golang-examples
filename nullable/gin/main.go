package main

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
)

const simpleJsonData = `{
  "name": "John Doe",
  "age": 0,
  "height": null,
  "married": false,
  "pet": "",
  "status": {
    "available": false,
    "blocked": false
  },
  "children": [
    {"age": 5},
    {"name": "Bob", "age": null}
  ]
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

type Status struct {
	Available NullBool `json:"available"`
	Blocked   NullBool `json:"blocked"`
}

type Child struct {
	Name NullString `json:"name"`
	Age  NullInt    `json:"age"`
}

type Person struct {
	Name     NullString `json:"name"`
	Age      NullInt    `json:"age"`
	Height   NullInt    `json:"height"`
	Married  NullBool   `json:"married"`
	Pet      NullString `json:"pet"`
	Status   Status     `json:"status"`
	Children []Child    `json:"children"`
}

func NewEmptyChild() Child {
	return Child{
		Name: NewEmptyNullString(),
		Age:  NewEmptyNullInt(),
	}
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

func NewPerson() Person {
	return Person{
		Name:    NewEmptyNullString(),
		Age:     NewEmptyNullInt(),
		Height:  NewEmptyNullInt(),
		Married: NewEmptyNullBool(),
		Pet:     NewEmptyNullString(),
		Status: Status{
			Available: NewEmptyNullBool(),
			Blocked:   NewEmptyNullBool(),
		},
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

func (c *Child) UnmarshalJSON(data []byte) error {
	// Define a temporary struct to hold the JSON data
	var temp struct {
		Name NullString `json:"name"`
		Age  NullInt    `json:"age"`
	}
	temp.Name = NewEmptyNullString()
	temp.Age = NewEmptyNullInt()

	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	c.Name = temp.Name
	c.Age = temp.Age

	return nil
}

//func (c *Child) UnmarshalJSON(data []byte) error {
//  // Option 1: Unmarshal into a map[string]interface{} and check for null
//
//  var childMap map[string]interface{}
//  err := json.Unmarshal(data, &childMap)
//  if err != nil {
//    return err
//  }
//
//  // Initialize child fields with default values or handle missing fields
//  c.Name = "default_name" // Set a default name
//  if name, ok := childMap["name"]; ok {
//    if nameStr, ok := name.(string); ok {
//      c.Name = nameStr
//    }
//  }
//
//  c.Age = 0 // Set a default age
//  if age, ok := childMap["age"]; ok {
//    if ageInt, ok := age.(int); ok {
//      c.Age = ageInt
//    }
//  }
//
//  // Option 2: Unmarshal directly into Child struct with additional logic
//
//  // You can also unmarshal directly into the Child struct,
//  // but you'll need to handle cases where specific fields are missing.
//
//  // err = json.Unmarshal(data, c)
//  // if err != nil && (data == nil || len(data) == 0) {
//  //   // Handle empty JSON
//  // } else if err != nil {
//  //   return err
//  // }
//
//  return nil
//}

func main() {
	r := gin.Default()

	r.POST("/json", func(c *gin.Context) {
		var person Person = NewPerson()
		if err := c.ShouldBindJSON(&person); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"person": person})
	})

	r.Run(":8080")
}
