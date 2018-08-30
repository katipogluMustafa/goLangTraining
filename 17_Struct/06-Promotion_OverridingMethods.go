package main

import "fmt"

type Person struct {
	Name string
	Age int
}

type DoubleZero struct {
	Person
	LicenseToKill bool
}

func (p Person) greeting(){
	fmt.Println("I'm just a regular person")
}
func (dz DoubleZero) greeting(){
	fmt.Println("Miss Moneypenny, so good to see u")
}

func main(){
	p1 := Person{
		Name: "Ian Flemming",
		Age : 44,
	}
	p2 := DoubleZero{
		Person : Person{
			Name : "James Bond",
			Age : 23,
		},
		LicenseToKill : true,
	}
	p1.greeting()
	p2.greeting()
	p2.Person.greeting()
}