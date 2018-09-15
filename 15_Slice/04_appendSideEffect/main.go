package main

import "fmt"

type user struct {
	likes int
}

func main() {
	// declare a slice of 3 users.
	users := make([]user, 3)

	// Share the user at index 1
	shareUser := &users[1]

	shareUser.likes++

	for i := range users {
		fmt.Printf("User: %d Likes: %d\n", i, users[i].likes)
	}

	// add a new user
	users = append(users, user{})

	//Add another like for the user that was shared
	shareUser.likes++
	fmt.Println("*************")
	for i := range users {
		fmt.Printf("User: %d Likes: %d\n", i, users[i].likes)
	}

}

/*
 * Notice the last like has not been recorded.
 * Since we appended new element to the users slice
 * append() created new underlying array and copied the old data into new array
 * but the shareUser pointer still points to the old underlying array
 * This is a memory leak, there is a reference to an unused array
 */
