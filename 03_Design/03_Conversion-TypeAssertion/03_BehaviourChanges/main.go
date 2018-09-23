// Sample program to show how method sets can affect behavior.

package main

import "fmt"

// user defines a user in the system.
type user struct {
	name  string
	email string
}

// String implements the fmt.Stringer interface.
func (u *user) String() string {
	return fmt.Sprintf("My name is %q and my email is %q", u.name, u.email)
}

func main() {

	// Create a value of type user.
	u := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}

	// Display the values.
	fmt.Println(u)
	fmt.Println(&u)
}

/*
 * You need to user pointer semantics otherwise you don't overwrite the method.
 */
