package main

import "fmt"

func max(x int ) int {
	return 42 + x
}

func main(){
	fmt.Println(max(5))

	max := max(7)
	fmt.Println(max)
	//fmt.Println(max(5)) //You can't call max() function anymore in this scope, it is shadowed
	//max is variable anymore not a function in this scope
}
