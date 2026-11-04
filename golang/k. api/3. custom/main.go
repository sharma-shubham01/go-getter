package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

// Custom Marshal and Unmarshal
// 1. The encoding/json provides Marshaler and Unmarshaler interfaces

// Marshaler Interface
// 1. By implementing Marshaler interface, we can customize marshaling of our type
// 2. Signature -
type Marshaler interface {
	MarshalJSON() ([]byte, error)
}

// Unmarshaler Interface
// 1. By implementing Unmarshaler interface, we can customize unmarshalling of our type
// 2. Signature -
type Unmarshaler interface {
	UnmarshalJSON([]byte) error
}

type CarType int

const (
	_ CarType = iota
	Hatchback
	Sedan
	SUV
)

var (
	stringToCarType = map[string]CarType{
		"hatchback": Hatchback,
		"sedan":     Sedan,
		"suv":       SUV,
	}

	carTypeToString = map[CarType]string{
		Hatchback: "hatchback",
		Sedan:     "sedan",
		SUV:       "suv",
	}
)

func (c CarType) string() string {
	return carTypeToString[c]
}

func (c CarType) MarshalJSON() ([]byte, error) {
	t, ok := carTypeToString[c]
	if !ok {
		return nil, errors.New("invalid car type")
	}

	return []byte(`"` + t + `"`), nil
}

func (c *CarType) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	t, ok := stringToCarType[s]
	if !ok {
		return errors.New("invalid car type")
	}

	*c = t
	return nil
}

func main() {
	c := Hatchback
	b, err := json.Marshal(c)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b)) // "hatchback"

	var unmarshalledCarType CarType
	err = json.Unmarshal([]byte(`"suv"`), &unmarshalledCarType)
	if err != nil {
		panic(err)
	}

	fmt.Println(unmarshalledCarType) // suv
}
