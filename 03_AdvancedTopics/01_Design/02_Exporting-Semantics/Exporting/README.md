# Exporting

* Private and public is about the data, exporting is about identifier.

* Field level encapsulation
```go
package users

type User struct {
	Name     string  // Exported Field
	ID       int     // Exported Field
	
	password string  // Unexported field
}
```

*
```go
package users

// User represents information about a user.
type user struct { // Unexported
	Name    string // Exported Field
	ID      int    // Exported Field
}

// Manager represents information about a manager.
type Manager struct { // Exported
	Title string  // Exported Field
	
	user         // Embedded unexported user
}

/*
 * Because of inner type promotion
 * when we try to access to Manager from outside the users package
 * we are gonna have access to Title, Name , ID
 */

```
```go
package main

import (
	"users" /* import the users package */
    "fmt"
	)
func main(){
	
	// Create a value of type Manager from the users package.
	u := users.Manager{
		Title: "Dev Manager",
	}
	/*
	 * The problem is during the construction
	 * we can't initialize the inner value
	 * because that name's been identified using the lower case letter(user unexported field)
	 */
	
	// Set the exported fields from the unexported user inner type.
	u.Name = "Chole"
	u.ID = 10
	
	fmt.Printf("User: %#v\n",u)
}

```