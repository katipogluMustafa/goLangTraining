package main

import "fmt"

func main(){
	var x int
	fmt.Printf("%v\n",x)

	/* Below line is a func expression ,you assign a func to a variable*/
	increment := func() int { //This is a anonymous func, func without a name
		x++
		return x
	}

	fmt.Println(increment())
	fmt.Println(increment())

	/*

		Closure helps us limit the scope of variables used by multiple functions without closure,
	for two or more functions to have access to the same variable, that variable would need to be package scope

	*/

}