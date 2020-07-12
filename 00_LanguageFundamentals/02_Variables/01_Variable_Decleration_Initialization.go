package main

import "fmt"

func main() {

	/* Declare Many Variables At Once */
	{
		var a, b, c, d int
		var customerName string
		var length, width float32

		customerName = "Mustafa"

		// Assign one by one
		length = 1.78
		width = 67

		// Assign at once
		length, width = 1.79, 68

		// Use The Variables
		fmt.Println(a, b, c, d)
		fmt.Printf("Customer Name is %s ", customerName)
		fmt.Printf("\nLength: %.2f, Width: %.2f\n", length, width)
		fmt.Println("--------------------")
	}
	/* Initialize Many Variables At Once */
	{
		var a, b, c, d int = 1, 2, 3, 4
		fmt.Println(a, b, c, d)
		fmt.Println("--------------------")
	}
	/* Infer Type */
	{
		// If you assign value the moment you declare, no need to give type info
		var a, b, c, d = 5, 6, 7, 8
		fmt.Println(a, b, c, d)
		fmt.Println("--------------------")
	}
	/*  Infer mixed up types */
	{
		var a, b, c, d = 9, true, "hey", 5.2
		fmt.Println(a, b, c, d)
		fmt.Println("--------------------")
	}
	/* Initialize shorthand */
	{
		a, b, c := 1, false, "hy"
		d := true
		o := 'e'
		fmt.Println(a, b, c, d, o)
		fmt.Println("--------------------")
	}

}
