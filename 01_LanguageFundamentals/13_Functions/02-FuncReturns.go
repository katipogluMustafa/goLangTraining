package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(greet1("mustafa", "katipoğlu"))
}
func greet1(name string, surname string) string {
	return fmt.Sprintf("Selamu Aleykum %s %s", strings.Title(name), strings.Title(surname))
}

/*
 * fmt.Sprintf Concatenates the strings in the format we want and returns strings
 * In the other words, Sprintf formats according to a format specifier and returns the resulting string.
 * likewise we could be using fmt.Sprint or fmt.Sprintln funcs
 */

/*
 * greet1 is a func which takes two parameters and returns a string
 * greet1 doesn't print anything to screen, it just return string
 * then in main func, fmt.Println func prints the returning value of greet1 to screen
 */
