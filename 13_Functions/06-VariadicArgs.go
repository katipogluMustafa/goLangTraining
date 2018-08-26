package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	n := average1(data...) // ... says take the slice and pull each item out one at a time
	fmt.Println(n)
}
func average1(sf ...float64) float64 {
	var total float64

	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}

/*
 * In the parameters ... are in the front which means unlimited number of the type in this case float64
 * In the arguments ... are at the end which means take this data structure and pull each item one at a time
 */
