package main

import "fmt"

func main(){
	greet := map[int]string{
		0: "Selamu Aleyküm",
		1: "Good Morning",
	}

	for key, val := range greet{
		fmt.Println(key," - ", val)
	}
}
