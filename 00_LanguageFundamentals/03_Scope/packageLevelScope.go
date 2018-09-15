package main

import "fmt"

var x int = 42

func main(){
	foo()
}

func foo(){
	fmt.Println(x)
	fmt.Println(z)//z is visible even though it is declared after this line since its scope is package level
}
/* z is visible throughout the entire package*/
var z string = "hey"
