package main

import (
	"fmt"
	"math"
)

type Square struct {
	side float64
}
// another shape
type Circle struct {
	radius float64
}

type Shape interface {
	area() float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(z Shape){
	fmt.Println(z)
	fmt.Println(z.area())
}

func main(){
	s := Square{10}
	c := Circle{5}
	info(s)
	info(c)
}
