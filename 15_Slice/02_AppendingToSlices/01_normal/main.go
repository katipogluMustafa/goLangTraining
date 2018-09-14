// how append grows the capacity of the underlying array.

package main

import "fmt"

func main() {
	// nil slice
	var data []string

	lastCap := cap(data)

	for record := 1; record <= 1e5; record++ {
		value := fmt.Sprintf("Rec: %d", record)
		data = append(data, value)

		// When capacity changes, display the changes.
		if lastCap != cap(data) {
			changePercent := float64(cap(data)-lastCap) / float64(lastCap) * 100

			lastCap = cap(data)

			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n",
				&data[0], record, cap(data), changePercent)
		}
	}
}
