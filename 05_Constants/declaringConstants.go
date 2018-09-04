package main

import "fmt"

func main(){
	const x string = "mustafa"
	fmt.Println(x)

	const (
		speedOfLight	= 299792458
		pi				= 3.14
		myFavoriteNumber= 98
	)
	fmt.Printf("speedOfLight: %v\tpi: %v\tmyFavoriteNumber: %v\n",speedOfLight,pi,myFavoriteNumber)
	const (
		a int = 0
		b	  = 1.8 + 3i
		c	  = 2.7
	)
	fmt.Printf("a: %v\tb: %v\tc: %v\n",a,b,c)
}