package main

import "fmt"

type person struct {
	first string
	last string
	age int
}
// func receiver    identifier parameters return type
// func (p person)  fullName   ()         string
func (p person) fullName() string {
	return p.first + " " + p.last
}
//Receiver attaches this function to the person type
func main(){
	p1 := person{"James", "Bond", 20}
	p2 := person{"Todd", "Mc", 18}
	fmt.Println(p1.fullName() )
	fmt.Println(p2.fullName() )
}