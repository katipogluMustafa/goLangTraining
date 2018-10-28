package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n", u.name, u.email)
}

type admin struct {
	person user // Not embedding
	level  string
}

func main() {
	// Create an admin user.
	ad := admin{
		person: user{
			name:  "Mustafa Katipoğlu",
			email: "not-real-email@hotmail.com",
		},
		level: "super",
	}

	// We can access fields methods.
	ad.person.notify()
}
