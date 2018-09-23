* Start from concrete data then move up this way.
 
 ![Figure 1](img/theWayToSolveProblems.png?)

* Functions should be your first choice until they are not reasonable or practical to do so.

* We must be consistent with semantics.Data drives the semantic.So once the semantic is chosen either everything is value or everything is pointer.And the method and the code you write has to respect the semantic for the data.

* When naming receivers , you want that receiver name to be short , and sweet :P 
    * Receiver name should never be more than three letters long.

* When it comes to calling a method, go only cares that the data is of type.It doesn't care that value or pointer of that data.It all cares we are working with that type.
---
* What does go do underneath the method calls ?
   
```go
package main

import "fmt"

// data is a struct to bind methods to.
type data struct {
	name string
	age  int
}

// displayName provides a pretty print view of the name.
func (d data) displayName() {
	fmt.Println("My Name Is", d.name)
}

// setAge sets the age and displays the value.
func (d *data) setAge(age int) {
	d.age = age
	fmt.Println(d.name, "Is Age", d.age)
}

func main() {

	// Declare a variable of type data.
	d := data{
		name: "Bill",
	}

	fmt.Println("Proper Calls to Methods:")

	// How we actually call methods in Go.
	d.displayName()
	d.setAge(45)

	fmt.Println("\nWhat the Compiler is Doing:")

	// This is what Go is doing underneath.
	data.displayName(d)
	(*data).setAge(&d, 45)

```

* Functions  in go are values. They are literally typed values that means we can pass a function by its name anywhere we want in our program as long as type information is the same{functions signature}.

* Escape analysis algorithms can't track whether or not this value can stay on a stack or not.In other words, even though there is no reason for d to end up on the heap, it has to now allocate.
```go
	fmt.Println("\nCall Value Receiver Methods with Variable:")

	// Declare a function variable for the method bound to the d variable.
	// The function variable will get its own copy of d because the method
	// is using a value receiver.
	f1 := d.displayName

	// Call the method via the variable.
	f1()

	// Change the value of d.
	d.name = "Joan"

	// Call the method via the variable. We don't see the change.
	f1()
}	
```
* There is a cost to decoupling in go. When you decouple a piece of concrete data in go you're going to have the cost of indirection and allocation.There is no way around it right now :/

![Figure 1](img/decouplingAllocation.png?)

* We have to make sure that if we're decoupling, cost of indirection and allocation worth it.Do not wanna blindly decouple code in go. 