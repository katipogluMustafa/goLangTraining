package main

import "fmt"

type contact struct {
	greeting string
	name     string
}

func SwitchOnType(x interface{}) {
	switch x.(type) { // This is an assertion which says Hey Give me the the type of x
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("string")
	case contact:
		fmt.Println("contact")
	default:
		fmt.Println("unknown")
	}
}

func main() {
	SwitchOnType(7)
	SwitchOnType("Mustafa")
	var t = contact{"Good to see you", "tim"}
	SwitchOnType(t)
	SwitchOnType(t.greeting)
	SwitchOnType(t.name)
}
/*

	* Normally we switch on value of a variable but go allows us to switch on type of variable

 */