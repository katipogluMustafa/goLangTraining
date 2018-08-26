package main

import "fmt"

func hello() {
	fmt.Print("Hello ")
}
func world() {
	fmt.Print("world ")
}
func main() {
	defer world()
	defer hello()
	defer hello()
	defer world()
	hello()
}
/*
 * defer keyword is used to delay the function
 * when used execution is delayed until all the other codes inside this func runs
 * before the function quits, the deferred functions runs
 * Each deferred function is added to the stack(First in Last out)
 * The first deferred func will be executed last
 * The last deferred func will be executed first
 */