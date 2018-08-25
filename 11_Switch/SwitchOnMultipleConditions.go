package main

import "fmt"

func main() {
	switch "Jenny" {
	case "Tim", "Jenny":
		fmt.Println("Wassup Tim , or  Jenny")
	case "Marcus", "Medhi":
		fmt.Println("Bot of your names start with M")
	case "Julian", "Sushant":
		fmt.Println("Wassup Julian / Sushant")
	}
}
