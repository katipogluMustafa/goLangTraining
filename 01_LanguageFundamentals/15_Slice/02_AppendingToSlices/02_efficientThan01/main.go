package main

import "fmt"

func main() {
	// instead of declaring nil slice
	// since we know the capacity
	//this is an opportunity to set the capacity ahead of time
	data := make([]string, 0, 1e5)
	// See lower of the page for other comment
	lastCap := cap(data)

	for record := 1; record <= 1e5; record++ {
		value := fmt.Sprintf("Rec: %d", record)
		data = append(data, value)

		if lastCap != cap(data) {
			changePercent := float64(cap(data)-lastCap) / float64(lastCap) * 100

			lastCap = cap(data)

			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n",
				&data[0], record, cap(data), changePercent)
		}
	}
}

/*
 * Why we set the length of the data slice to 0 ?
 * If we set the length to 1e5 we're gonna have 1e5 piece zero valued string and another 1e5 actual string
 */

/*
 * We don't see any output when we run this
 * because we don't ever change the capacity of the slice
 */
