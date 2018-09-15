package main

import "fmt"

func main() {
	// If the entry exist, delete the entry from the map
	myGreeting := make(map[string]string)

	myGreeting["Mustafa"] = "Selamu Aleyküm"
	myGreeting["Ahmet"] = "Aleyküm Selam"
	myGreeting["Mahmut"] = "Ve Aleyküm Selam"

	fmt.Println("Before Deletion: ", myGreeting)

	// Keep the scope tight
	if val, exists := myGreeting["Mahmut"]; exists {
		delete(myGreeting, "Mahmut")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	} else {
		fmt.Println("That value doesn't exist.")
		fmt.Println("val: ", val)
		fmt.Println("exists: ", exists)
	}

	fmt.Println("After Deletion: ", myGreeting)
}
