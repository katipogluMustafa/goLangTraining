// Sample program to show how to implement custom error type
// based on the json package in the standard library.

package main

import (
	"fmt"
	"reflect"
)

// An UnmarshalTypeError describes a JSON value that was
// not appropriate for a value of a specific Go type.
type UnmarshalTypeError struct {
	Value string       // Description of JSON value
	Type  reflect.Type // type of Go value it could not be assigned to
}

// Error implements the error interface.
func (e *UnmarshalTypeError) Error() string {
	return "json: cannot unmarshal " + e.Value + " into Go value of type " + e.Type.String()
}

// An InvalidUnmarshalError describes an invalid argument passed to Unmarshal.
// (The argument to Unmarshal must be a non-nil pointer.)
type InvalidUnmarshalError struct {
	Type reflect.Type
}

// Error implements the error interface.
func (e *InvalidUnmarshalError) Error() string {
	if e.Type == nil {
		return "json: Unmarshal(nil)"
	}

	if e.Type.Kind() != reflect.Ptr {
		return "json: Unmarshal(non-pointer " + e.Type.String() + ")"
	}

	return "json: Unmarshal(nil " + e.Type.String() + ")"
}

// User is a type for using in the Unmarshal call.
type user struct {
	Name int
}

func main() {
	var u user

	err := Unmarshal([]byte(`{"name":"bill"}`), u) // Run with a value and pointer

	if err != nil {
		switch e := err.(type) { // type assertion on switch, we get a copy of what we stord
		case *UnmarshalTypeError:
			fmt.Printf("UnmarshalTypeError : Value[%s]  Type[%v]\n", e.Value, e.Type)
		case *InvalidUnmarshalError:
			fmt.Printf("InvalidUnmarshalError: Type[%v]\n", e.Type)
		default:
			fmt.Println(err)
			return
		}
	}

	fmt.Println("Name: ", u.Name)
}

// Unmarshal simulates an unmarshal call that always fails.
func Unmarshal(data []byte, v interface{}) error {
	rv := reflect.ValueOf(v) // this is how we always start reflection.

	/* Is the concrete data inside of v , not a pointer or nil ? */
	if rv.Kind() != reflect.Ptr || rv.IsNil() {
		return &InvalidUnmarshalError{reflect.TypeOf(v)}
	}

	return &UnmarshalTypeError{"string", reflect.TypeOf(v)}
}
