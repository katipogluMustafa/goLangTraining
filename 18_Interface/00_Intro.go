package main

import "fmt"

type Square struct {
	side float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

type Shape interface {
	area() float64
}

func info(z Shape){
	fmt.Println(z)
	fmt.Println(z.area())
}

func main(){
	s := Square{10}
	info(s)
}

/*
 * Interface enables functions to have a parameter that can act as a placeholder for multiple types that implement that interface.
 */
