package main

import "fmt"

func main(){
	mustafa := []int{8,7,6,5,4}

	fmt.Printf("%T,%v",mustafa,mustafa)
	changeTheElements(mustafa)
	fmt.Println()
	fmt.Println(mustafa)
}
func changeTheElements(x []int){
	x[2] = 4
	x[1] = 1
}
