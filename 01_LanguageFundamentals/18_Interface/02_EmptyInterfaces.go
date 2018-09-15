package main

import (
	"fmt"
	"goLangTraining/18_Interface/simpleShapes"
	"math/rand"
	"time"
)

func randomShape() interface{}{
	var shape interface{}
	randomShapesSlice := make([]interface{},2)

	//Seed the random generator
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	// Pick a number, either 1 or 0
	randomNumber := r.Intn(2)
	fmt.Printf("Random Number:%d\n",randomNumber)

	// Let's make a new rectangle
	rectangle := simpleShapes.NewRectangle(3,6)

	// Let's make a new triangle
	triangle := simpleShapes.NewTriangle(9,18)

	// Let's store the shapes into a slice data structure
	randomShapesSlice[0] = rectangle
	randomShapesSlice[1] = triangle
	shape = randomShapesSlice[randomNumber]
	return shape
}

func main(){
	myRectangle := simpleShapes.NewRectangle(4,5)
	myTriangle  := simpleShapes.NewTriangle(2,7)
	shapesSlice := []interface{}{myRectangle,myTriangle}

	for i,s := range shapesSlice {
		fmt.Printf("Shape in index, %d, of the shapesSlice is a %T\n",i,s)
	}

	aRandomShape := randomShape()
	fmt.Printf("The type of aRondomShape is %T\n",aRandomShape)

	/* Type switch to discover dynamic type of this interface variable*/
	switch t := aRandomShape.(type) {
	case *simpleShapes.Rectangle:
		fmt.Println("I got back a rectangle with an area equal to ", t.Area())
	case *simpleShapes.Triangle:
		fmt.Println("I got back a triangle with an area equal to ", t.Area())
	default:
		fmt.Println("I don't recognize what I got back!")
	}

}

/*
 * Every type in go implements the empty interface(interface{})
 * This includes all of  the primitive types as well as types that we create.
 * Some Use Cases:
 * 		A function that returns a value of interface{} can return any type.
 * 		We can store heterogeneous values in an array, slice, or map using  the interface{} type.
 */
