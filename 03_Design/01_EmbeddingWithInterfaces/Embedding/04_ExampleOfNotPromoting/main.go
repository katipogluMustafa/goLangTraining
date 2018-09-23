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

type admin struct {
	user
	level string
}

// this time admin type also implements notifier interface
func (a *admin) notify() {
	fmt.Printf("Sending admin email To %s<%s>\n", a.name, a.email)
}

func main() {

	ad := admin{
		user: user{
			name:  "Mustafa Katipoğlu",
			email: "not-real-email@hotmail.com",
		},
		level: "super",
	}
	/*
	 * Send the admin user a notification.
	 * The embedded inner type's implementation of the interface
	 * is NOT "promoted" to the outer type.
	 */
	sendNotification(&ad)

	// We can access the inner type's method directly.
	ad.user.notify()

	// The inner type's method is NOT promoted
	ad.notify()
}

func sendNotification(n notifier) {
	n.notify()
}

/*
 * When the outer type implements the same thing as the inner type then there is no promotion.
 * The outer type will overwrite the inner type's implementation.
 */
