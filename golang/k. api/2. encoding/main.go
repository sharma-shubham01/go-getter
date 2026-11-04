package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Marshaling
// Converting Go data structures to JSON data is called marshalling

// Unmarshalling
// Converting JSON data into Go data structures is called unmarshalling

// Package encoding/json provides function to marshal and unmarshal
// Tags should be used for mapping

// JSON Tags
// 1. JSON tags are used in Goto map a field of a struct to it's JSON counterpart
// 2. Added inside the backticks at the end of struct field declaration
// 3. Inside the tags you provide the mapping and the other options

// Common JSON Tags Option
// 1. 'Omitempty' - It is used to omit the struct field from marshaling when zero valued
// 2. '-' - It is used to indicate that the field should completely be ignored from marshaling

// Prefer using Tags
// 1. If no tags are provided to the map field names, by default the field name should match it's JSON counterpart
// 2. JSON Marshal and Unmarshal requires the struct fields to be exported

//{
//	"user_id": 12345,
//	"name": "User A",
//	"age": 35,
//	"password": "my-pwd",
//	"roles": ["admin", "read"]
//}

type User struct {
	ID          int      `json:"user_id"`
	Name        string   `json:"name,omitempty"`
	Age         int      `json:"age"`
	Password    string   `json:"-"`
	Permissions []string `json:"roles"`
}

func jsonMarshal() {
	u := User{
		ID:          1,
		Name:        "User One",
		Age:         10,
		Password:    "my-pwd",
		Permissions: []string{"admin", "read"},
	}

	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(b))
}

type Person struct {
	Name string `json:"full_name"`
	Age  int    `json:"years_old,omitempty"`
}

func jsonUnmarshal() {
	jsonData := `{"full_name":"user a", "years_old":25}`

	var person Person

	err := json.Unmarshal([]byte(jsonData), &person)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(person)
}

// JSON Encoding and Decoding
// 1. The io.Reader and io.Writer interface support reading and writing byte data
// 2. It is efficient to use io.Reader and io.Writer to read JSON byte data directly into Go structs and vice-versa
// 3. This is where encoding/json.Encoder and encoding/json.Decoder comes into play

type Human struct {
	Name   string `json:"username"`
	Age    int    `json:"age"`
	Active bool   `json:"is_active"`
}

func jsonEncoder() {
	h := Human{
		Name:   "Alice",
		Age:    20,
		Active: true,
	}

	f, err := os.Create("output.json")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	encoder := json.NewEncoder(f)
	err = encoder.Encode(h)
	if err != nil {
		panic(err)
	}
}

func jsonDecoder() {
	var h Human

	f, err := os.Open("output.json")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	decoder := json.NewDecoder(f)
	err = decoder.Decode(&h)
	if err != nil {
		panic(err)
	}

	fmt.Println(h)
}

func main() {
	jsonMarshal()
	jsonUnmarshal()
	jsonEncoder()
	jsonDecoder()
}
