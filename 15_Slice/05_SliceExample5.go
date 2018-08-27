package main

import "fmt"

func main() {
	customerIds := []int{1, 2, 3, 4, 5}
	otherCustomerIds := []int{6, 7, 8, 9}

	fmt.Println(customerIds)
	fmt.Println(otherCustomerIds)

	customerIds = append(customerIds, otherCustomerIds...)

	fmt.Println()
	fmt.Println(customerIds)
	fmt.Println(otherCustomerIds)
}
