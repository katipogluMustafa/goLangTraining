package main

import "fmt"

func main(){
	y := 42
	fmt.Println(y)
	f()
	r()
}

func f(){

//	fmt.Println(y) // Won't compile since y is not visible in this scope, y is undefined here

}

func r(){
	/*
	//This won't compile, order matters, k is visible from the point it declared to end of the block
	fmt.Println(k) //k is not visible here
	k := 5 //k is visible from now on
	*/
}
