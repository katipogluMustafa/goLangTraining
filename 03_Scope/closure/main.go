package main

import "fmt"

func main(){

	x := 42
	fmt.Println(x)

	{

		fmt.Println(x)//x is visible since x is declared at outer scope
		y := "Some text"
		fmt.Println(y)//y is only visible inside of the braces

	}

	//fmt.Println(y) //won't work since y is not visible

}
