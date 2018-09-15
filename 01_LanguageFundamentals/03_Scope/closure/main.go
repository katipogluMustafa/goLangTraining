package main

import "fmt"

func main(){

	/* Anything declared in the outer scope is accessible in scope that are enclosed by that outer scope */

	x := 42
	fmt.Println(x)
	fmt.Println(k)//k is declared outer scope, it is visible in this enclosed scope
	{

		fmt.Println(x)//x is visible since x is declared at outer scope and outer scope enclosing inner scope
		y := "Some text"
		fmt.Println(y)//y is only visible inside of the braces

	}

	//fmt.Println(y) //won't work since y is not visible

}
