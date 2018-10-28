package main

import "fmt"

type notifier interface {
	notify()
}

type user struct {
	name  string
	email string
}

func (u *user) notify() {
	fmt.Printf("Sending user email To %s<%s>\n", u.name, u.email)
}

/* admin represents an admin user with privileges */
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
	// Send the admin user a notification.
	// The embedded inner type's implementation of the interface
	// is "promoted" to the outer type.
	sendNotification(&ad)
}

func sendNotification(n notifier) {
	n.notify()
}

/*
 * This admin value(ad) now satisfies all the same interfaces as the inner type did
 * because we can promote that method set up.

 * And now this outer type, concrete piece of data , admin, now has the right behaviours because of inner type promotion to satisfy the notifier interface.
 */
