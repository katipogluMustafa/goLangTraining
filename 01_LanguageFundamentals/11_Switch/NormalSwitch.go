package main

import "fmt"

func main() {

	switch "Medhi" {
	case "Daniel":
		fmt.Println("Wassup Daniel")
	case "Medhi":
		fmt.Println("Wassup Medhi")
	case "Jenny":
		fmt.Println("Wassup Jenny")
	default:
		fmt.Println("Who Are u ?")
	}

}

/*
* No default fallthrough
* Fallthrough is optional, you can specify fallthrough by explicitly stating it
* Break isn't needed like other programming languages
 */
