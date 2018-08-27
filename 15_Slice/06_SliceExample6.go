package main

import "fmt"

func main() {
	days1 := []string{"Monday", "Tuesday"}
	days2 := []string{"Wednesday", "Thursday", "Friday"}

	days1 = append(days1, days2...)
	fmt.Println(days1)

	/* Let's Delete Wednesday */
	days1 = append(days1[:2], days1[3:]...) // Both of the arguments days1
	fmt.Println(days1)

}
