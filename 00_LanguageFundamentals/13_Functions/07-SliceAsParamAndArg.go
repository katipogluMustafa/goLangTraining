package main

import "fmt"

func main() {
	data := []float64{43, 56, 87, 12, 45, 57}
	n := average2(data)
	fmt.Println(n)
}

func average2(sf []float64) float64 {
	var total float64

	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
