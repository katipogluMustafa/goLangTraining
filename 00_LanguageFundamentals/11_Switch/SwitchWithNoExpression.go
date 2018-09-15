package main

import "fmt"

func main() {
	var myFriendsName string
	fmt.Print("What is My Friend's name: ")
	fmt.Scan(&myFriendsName)
	switch {
	case len(myFriendsName) == 2:
		fmt.Println("Wassup my friend with name of length 2")
	case myFriendsName == "Tim":
		fmt.Println("Wassup Tim")
	case myFriendsName == "Jenny":
		fmt.Println("Wassup Jenny")
	case myFriendsName == "Marcus", myFriendsName == "Medhi":
		fmt.Println("Your name is either Marcus or Medhi")
	case myFriendsName == "Julian":
		fmt.Println("Wassup Julian")
	case myFriendsName == "Sushant":
		fmt.Println("Wassup Sushant")
	default:
		fmt.Println("Nothing matched; This is default.")
	}
}

/*
	Expression not needed
		If no expression provided, go checks for the first case that evaluates to true
		This makes the switch operate like if/ if else/else
		Cases can be expressions
*/
