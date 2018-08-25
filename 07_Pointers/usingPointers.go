package main

import "fmt"

func main(){
	a := 45
	fmt.Println("Before the change a : ", a)

	var b *int = &a

	*b = 60

	fmt.Println("After the change a : ",a)
}