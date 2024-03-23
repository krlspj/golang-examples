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

func (p Person) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]interface{})

	if !p.Name.IsEmpty {
		_bytes, err := json.Marshal(p.Name)
		if err != nil {
			return nil, err
		}
		jsonData["name"] = json.RawMessage(_bytes)
	}
	if !p.Age.IsEmpty {
		_bytes, err := json.Marshal(p.Age)
		if err != nil {
			return nil, err
		}
		jsonData["age"] = json.RawMessage(_bytes)
	}
	if !p.Height.IsEmpty {
		_bytes, err := json.Marshal(p.Height)
		if err != nil {
			return nil, err
		}
		jsonData["height"] = json.RawMessage(_bytes)
	}
	if !p.Married.IsEmpty {
		_bytes, err := json.Marshal(p.Married)
		if err != nil {
			return nil, err
		}
		jsonData["married"] = json.RawMessage(_bytes)
	}
	if !p.Pet.IsEmpty {
		_bytes, err := json.Marshal(p.Pet)
		if err != nil {
			return nil, err
		}
		jsonData["pet"] = json.RawMessage(_bytes)
	}

	jsonData["status"] = map[string]interface{}{}
	statusMap := make(map[string]any)
	if !p.Status.Available.IsEmpty {
		_bytes, err := json.Marshal(p.Status.Available)
		if err != nil {
			return nil, err
		}
		statusMap["available"] = json.RawMessage(_bytes)
	}
	if !p.Status.Blocked.IsEmpty {
		_bytes, err := json.Marshal(p.Status.Blocked)
		if err != nil {
			return nil, err
		}
		statusMap["blocked"] = json.RawMessage(_bytes)
	}
	if len(statusMap) > 0 {
		jsonData["status"] = statusMap
	}

	if len(p.Children) > 0 { // Check if Children slice is not empty
		childrenBytes := make([]json.RawMessage, len(p.Children))
		for i, child := range p.Children {
			childBytes, err := json.Marshal(child)
			if err != nil {
				return nil, err
			}
			childrenBytes[i] = json.RawMessage(childBytes)
		}
		jsonData["children"] = childrenBytes
	}
	return json.Marshal(jsonData)
}

func (this Child) MarshalJSON() ([]byte, error) {
	jsonData := make(map[string]any)
	if !this.Name.IsEmpty {
		_bytes, err := json.Marshal(this.Name)
		if err != nil {
			return nil, err
		}
		jsonData["name"] = json.RawMessage(_bytes)
	}
	if !this.Age.IsEmpty {
		_bytes, err := json.Marshal(this.Age)
		if err != nil {
			return nil, err
		}
		jsonData["age"] = json.RawMessage(_bytes)
	}
	return json.Marshal(jsonData)

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

func (nb NullBool) MarshalJSON() ([]byte, error) {
	if nb.IsNull {
		return []byte("null"), nil
	}
	if nb.IsEmpty {
		return []byte("{}"), nil
	}
	return json.Marshal(nb.Value)
}

func (ni NullInt) MarshalJSON() ([]byte, error) {
	if ni.IsNull {
		return []byte("null"), nil
	}
	if ni.IsEmpty {
		return []byte("{}"), nil
	}
	return json.Marshal(ni.Value)
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

		//mperson, err := json.Marshal(&person)
		//_ = err
		//if err != nil {
		//	fmt.Println("error marshalling", err.Error())
		//	return
		//}
		//c.JSON(200, gin.H{"marshalledPerson": string(mperson)})
		//c.JSON(200, gin.H{"person": person, "marshalledPerson": mperson})
		c.JSON(200, gin.H{"person": person})
		//fmt.Println("string mperson", mperson)
	})

	r.Run(":8080")
}
