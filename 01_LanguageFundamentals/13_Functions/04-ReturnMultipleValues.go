package main

import "fmt"

func main() {
	fmt.Println(greet3("Mustafa ", "Katipoğlu "))
}

func greet3(fname string, lname string) (string, string) {
	return fmt.Sprint(fname, lname), fmt.Sprint(lname, fname)
}
