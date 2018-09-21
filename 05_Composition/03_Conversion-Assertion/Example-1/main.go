package main

import "fmt"

// Mover supports for moving things.
type Mover interface {
	Move()
}

// Locker provides support for locking and unlocking things.
type Locker interface {
	Lock()
	Unlock()
}

// MoveLocker provides support for moving and locking things.
type MoveLocker interface {
	Mover
	Locker
}

/*
 * Use empty struct when you might wanna method based API where there is no state.
	- Maybe we're using bike as namespace.
 * Remember functions always gonna be more precise than methods
    - but in this particular case since there is no state
    - than the API would have to be as precise as any function
*/

// bike represents a concrete type for the example.
type bike struct{}

// Move can change the position of a bike.
func (bike) Move() {
	fmt.Println("Moving the bike")
}

// Lock prevents a bike from moving.
func (bike) Lock() {
	fmt.Println("Locking the bike")
}

func (bike) Unlock() {
	fmt.Println("Unlocking the bike")
}

func main() {

	// Declare variables to their zero value
	var ml MoveLocker
	var m Mover

	// Create a value of type bike, assign the value to the MoveLocker
	ml = bike{}

	// An interface value of type MoveLocker can be implicitly
	// converted into a value of type Mover. They both declare a method named move.

	m = ml

	/* Perform a type assertion against the Mover interface value 	 * to access a COPY of the concrete type value of type bike that was stored inside of it.
	 * Then assign the COPY of the concrete type to the MoveLocker interface.
	 */

	b := m.(bike)
	ml = b

	/*
	 * It's important to note that the type assertion syntax provides a way
	 * to state what type of value is stored inside the interface.
	 * This is more powerful from a language and readability standpoint,
	 * than using a casting syntax, like in other languages.
	 */

}

/*
 * Type assertions happens at runtime.
 * Type assertions and usage of pointer semantics in the method give the ability to overwrite methods.
 */
