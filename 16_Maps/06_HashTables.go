package main

import "fmt"

func main(){
		// Items in Maps have unique keys
		myMap := make(map[string]int)
		myMap["Mustafa"] = 5
		fmt.Println(myMap)
		myMap["Mustafa"] = 7
		fmt.Println(myMap)
		/* If they key pair you inserted already exists, the map updates the value of the key */

}

/*
 * For optimization, Hash table capacity should be prime number.It shouldn't be close to the powers of 2 or 10.
 */
 /*
  * Look at other details about Hash Tables and Maps
  */