package main

import "fmt"

func main(){
	for i := 48; i <= 127; i++{
		fmt.Println(i, "-", string(i), "-", []byte(string(i)))
	}
	for i := 5000; i <= 5227; i++{
		fmt.Println(i, "-", string(i), "-", []byte(string(i)))
	}
}

/*
	Rune is alias for int32. It is used to represent UTF-8 Characters.
				https://golang.org/doc/go1#rune
*/
