package main

import (
	"encoding/json"
	"fmt"
)

type Message struct {
	Author    string
	Content   string
	timestamp string
}

func simpleJsonMarshal() {
	// json package only accesses exported fields of struct types, so timestamp won't be encoded
	m := Message{"Alice", "Hello", "2020-01-01T00:00:00Z"}
	result, _ := json.Marshal(m)
	fmt.Println(string(result))
}

func simpleJsonUnmarshal() {
	// sample json
	// "extra" field will be ignored because it doesn't exist in Message struct
	// unexported fields (timestamp) will be ignored
	b := []byte(`{"author":"Alice","content":"Hello","timestamp":"2020-01-01T00:00:00Z","extra":"extra"}`)
	var m Message

	json.Unmarshal(b, &m)
	fmt.Println(m)
}

type User struct {
	// json package will use the name specified in the tag when encoding/decoding
	Name       string `json:"name"`
	Email      string `json:"email"`
	Password   string `json:"-"`                     // this tag means this field will be ignored in json
	MotherName string `json:"mother_name,omitempty"` // this tag means this field will be ignored in json if it's empty (zero value)
}

func jsonMarshalWithTags() {
	u := User{"Alice", "alice@gmail.com", "123456", ""}
	result, _ := json.Marshal(u)
	fmt.Println(string(result))

	u.MotherName = "Eve"
	result, _ = json.Marshal(u)
	fmt.Println(string(result))
}

func jsonUnmarshalWithTags() {
	b := []byte(`{"name":"Alice","email": "alice@gmail.com", "password": "123456", "mother_name":"Eve"}`)
	var u User

	json.Unmarshal(b, &u)
	// password won't be decoded because it's ignored in json
	fmt.Println(u)
}
