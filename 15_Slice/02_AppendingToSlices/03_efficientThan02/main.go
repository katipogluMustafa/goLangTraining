package main

import "fmt"

func main() {
	//But in this case you need to set the length
	// because we gotta make sure that all of that is available for use upfront
	data := make([]string, 1e5)
	lastCap := cap(data)

	// let's get rid of append call by indexing what we need
	// start 0
	for record := 0; record < 1e5; record++ {
		value := fmt.Sprintf("Rec: %d", record)
		data[record] = value

		if lastCap != cap(data) {
			changePercent := float64(cap(data)-lastCap) / float64(lastCap) * 100

			lastCap = cap(data)

			fmt.Printf("Addr[%p]\tIndex[%d]\t\tCap[%d - %2.f%%]\n",
				&data[0], record, cap(data), changePercent)
		}
	}

}
