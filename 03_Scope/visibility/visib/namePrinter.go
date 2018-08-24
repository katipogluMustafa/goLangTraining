package visib

import "fmt"

func PrintName(){ //Visible From outside of the package(uppercase start)
	fmt.Println(name())
}