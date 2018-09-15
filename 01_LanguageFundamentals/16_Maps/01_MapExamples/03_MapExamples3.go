package main

import "fmt"

func main() {

	m := make(map[string]int)

	m["route"] = 66
	m["sd"] = 76
	m["psd"] = 86

	// Retrieve the value stored under the key "route"
	i := m["route"]
	fmt.Println(i)

	//If the required key doesn't exist, we get the value type's zero value.
	// In this case the value type is int , so the zero value is 0
	j := m["root"]
	fmt.Println(j) // j return 0

	delete(m, "route")
	// The built in delete function removes an entry from the map
	// delete functions doesn't return anything, and won't do anything if the specified key doesn't exist

	// A two-value assignment tests for the existence of a key:
	i, ok := m["route"]
	if ok {
		fmt.Println("route exists, its value: ", i)
	}else{
		fmt.Println("route doesn't exit")
	}
	// In this statement, the first value (i) is assigned the value stored under the key "route". If that key doesn't exist, i is the value type's zero value (0). The second value (ok) is a bool that is true if the key exists in the map, and false if not.

	for key,value := range m{
		fmt.Println("Key: ", key, "Value: ", value)
	}
	//To iterate over the contents of a map, use the range keyword:



	// Items in Maps have unique keys
	myMap := make(map[string]int)
	myMap["Mustafa"] = 5
	fmt.Println(myMap)
	myMap["Mustafa"] = 7
	fmt.Println(myMap)
	/* If they key pair you inserted already exists, the map updates the value of the key */


}

/*
 * For optimization, map capacity should be prime number.
 */
