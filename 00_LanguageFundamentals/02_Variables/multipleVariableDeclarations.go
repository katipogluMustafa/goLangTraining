package main

import "fmt"

func main(){
	var variable1,variable2,variable3 string
	var variable4,variable5,variable6 int = 4,5,6
	var variable7,variable8,variable9 = 7,8,9
	variable10,variable11,variable12 := 10,11,12

	for i := 1; i < 12; i++ {
		s := fmt.Sprintf("%s%d","variable",i)
		fmt.Printf(s)
	}
	fmt.Println(variable1,variable2,variable3)
	fmt.Println(variable4,variable5,variable6)
	fmt.Println(variable7,variable8,variable9)
	fmt.Println(variable10,variable11,variable12)

}
