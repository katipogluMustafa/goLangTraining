package main

import "fmt"

func main(){

	var student1 []string
	//student1[0] = "Todd" // Doesn't work , since underlying structure hasn't been created properly
	student1 = append(student1, "Todd") // Works fine, append will handle the problem and add item
	fmt.Println("student1",student1)

	student2 := []string{}
	// student2[0] = "Ahmet" // Same, doesn't work, slice is not ready yet
	student2 = append(student2, "Todd")
	fmt.Println("student2",student2)

	student := make([]string, 5) // Create slice properly
	students := make([][]string, 5)
	student[0] = "Todd" // Works fine since its length set to 35
	fmt.Println(student)
	fmt.Println(students)
}