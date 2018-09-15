package main

import "fmt"

// user represents a user in the system.
type user struct {
	name  string
	email string
}

func main() {
	u1 := createUserV1()
	u2 := createUserV2()

	fmt.Printf("&u1: %p\n", &u1)
	fmt.Printf("&u2: %p\n", &u2)
	fmt.Println("u1", u1, "u2", u2)

}

// CreateUserV1 creates a user value and passed
// a copy back to the caller
func createUserV1() user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}
	fmt.Printf("V1 : %p\n", &u)
	return u
}

//createUserV2 creates a user value and shares
// the value with the caller
func createUserV2() *user {
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}
	fmt.Printf("V2 : %p \n", &u)
	return &u
}
