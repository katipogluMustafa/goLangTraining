package main

import (
	"fmt"
	"reflect"
)

func main(){
	// Found out the Types
	fmt.Println(reflect.TypeOf(42))
	fmt.Println(reflect.TypeOf("Probably A String, Right?"))
	fmt.Println(reflect.TypeOf(42.2222))
	fmt.Println(reflect.TypeOf('r')) // remember runes are unicode
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf(42.00000000000000000000000000000000000000000000000000000000000000000000000000000))
}