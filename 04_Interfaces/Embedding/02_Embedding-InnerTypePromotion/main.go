package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n", u.name, u.email)
}

/* User is the inner type, admin is the outer type.*/
type admin struct {
	user  // Embedding Type
	level string
}

func main() {
	// Create an admin user.
	ad := admin{
		user: user{
			name:  "Mustafa Katipoğlu",
			email: "not-real-email@hotmail.com",
		},
		level: "super",
	}

	// We can access the inner type's method directly.
	ad.user.notify()

	// Because of inner type promotion
	// we can also access notify directly through the outer type value.
	ad.notify() // The inner type's method is promoted.

}

/*
 * Everything related to the inner type can be promoted up to the outer type.
 * In other words promotion allows access to those things from the inner type through the outer type.
 */
