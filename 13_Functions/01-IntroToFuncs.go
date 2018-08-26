package main

import (
	"fmt"
	"strings"
)

/*
 * You can only have one func main() in your entire program
 * It needs to be inside package main
 * func main() is the entry point to your program
 * That's where your program starts
 */
func main() {
	greet("mustafa")                  // We pass the "mustafa" strings as an argument to the func greet
	greet("Ahmet")                    // "    "    "    "Ahmet"    "     "   "   "       "  "   "    "
	greetings("Muhammed", "Abdullah") // We pass two string arguments to the func greetings
	takeGreetings()
}

/*
 * name is a parameter which has type of string
 * scope of the name variable is only the greet function's blocks
 */
func greet(name string) {
	fmt.Println("Selamu aleykum", strings.Title(name), "!")
}

func greetings(name1 string, name2 string) {
	fmt.Println("Selamu aleykum", strings.Title(name1), ",", strings.Title(name2), "!")
}

func takeGreetings() {
	fmt.Println("Aleykum selam")
}

/*
 * Functions allow you to abstract the internal functionality of a partition of the program.
 * Instead of having long code section, you break your program into little modules.
 * We call this modules functions or func as in the go programming language.
 * These functions also allow us to reuse code.
 */
/*
 * greet ---> This one just a variable
 * greet() -> This one a func without parameters
 */
