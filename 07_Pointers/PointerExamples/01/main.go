package main

import "fmt"

type user struct {
	name   string
	email  string
	logins int
}

func main() {

	bill := user{
		name:  "Bill",
		email: "bill@ardanlabs.com",
	}
	display(&bill)
	increment(&bill.logins)
	display(&bill)
}

func display(u *user) {
	fmt.Printf("%p\t%+v\n", u, *u)
	fmt.Printf("Name: %q Email: %q Logins: %d\n\n", u.name, u.email, u.logins)
}

func increment(logins *int) {
	*logins++
	fmt.Printf("&logins[%p] logins[%p] *logins[%d]\n\n", &logins, logins, *logins)
}

/*
 * %q	a double-quoted string safely escaped with Go syntax
 */
