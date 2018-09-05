package main

import "fmt"

type Rectange struct {
	width 	float64
	height	float64
}

func newRectangle(w float64, h float64) *Rectange {
	return &Rectange{width: w, height:h,}
}

func (r *Rectange) area() float64 {
	return r.width * r.height
}

func main(){
	myRectangle := newRectangle(3,5)
	fmt.Println(myRectangle.area())
}