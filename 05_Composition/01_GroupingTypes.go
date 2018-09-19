/*
 * This is an example of using composition and interfaces.
 * This is something we want to do in Go.
 * We will group common types by their behavior and not by their state.
 * This pattern does provide a good design principle in a Go program.
 */
package main

import "fmt"

/*
 * Speaker provide a common behavior for all concrete types
 * to follow if they want to be a part of this group.
 * This is a contract for these concrete types to follow.
 */
type Speaker interface {
	Speak()
}

type Dog struct {
	Name       string
	IsMammal   bool
	PackFactor int
}

func (d *Dog) Speak() {
	fmt.Printf("Woof! My name is %s, it is %t I am a mammal with a pack factor of %d.\n",
		d.Name,
		d.IsMammal,
		d.PackFactor,
	)
}

type Cat struct {
	Name        string
	IsMammal    bool
	ClimbFactor int
}

func (c *Cat) Speak() {
	fmt.Printf(
		"Meow! My name is %s, it is %t I am a mammal with a climb factor of %d.\n",
		c.Name,
		c.IsMammal,
		c.ClimbFactor,
	)
}

func main() {
	speakers := []Speaker{
		&Dog{
			Name:       "fido",
			IsMammal:   true,
			PackFactor: 5,
		},
		&Cat{
			Name:        "Milo",
			IsMammal:    true,
			ClimbFactor: 4,
		},
	}

	for _, spkr := range speakers {
		spkr.Speak()
	}

}

// Notes:
/*
 * Declare types that represent something new or unique.
 * Validate that a value of any type is created or used on its own.
 * Embed types to reuse existing behaviors you need to satisfy.
 * Question types that are an alias or abstraction for an existing type.
 * Question types whose sole purpose is to share common state.
 */
