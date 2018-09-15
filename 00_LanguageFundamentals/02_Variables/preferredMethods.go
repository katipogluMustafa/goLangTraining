package main

import (
	"fmt"
)

func main(){



	/* Initialize, Declare and Assign to a Value */
	{
		fmt.Println("~~~~ Style 1 ~~~")
		var a float64 = 4.0
		var b bool = true
		var c string = "success"
		var d int = 4
		fmt.Printf("%v : %T\n",a,a)
		fmt.Printf("%v : %T\n",b,b)
		fmt.Printf("%v : %T\n",c,c)
		fmt.Printf("%v : %T\n",d,d)

	}

	/* Shorthand Initialize, Declare and Assign to a Value, Types Omitted */
	{
		fmt.Println("~~~~ Style 2 ~~~")
		a := 3.0
		b := false
		c := "heyoo"
		d := 5
		fmt.Printf("%v : %T \n",a,a)
		fmt.Printf("%v : %T \n",b,b)
		fmt.Printf("%v : %T \n",c,c)
		fmt.Printf("%v : %T \n",d,d)

	}

	/* Initialize, Declare which will automatically be assigned to 0 of the same type*/
	{
		fmt.Println("~~~~ Style 3 ~~~")
		var a float64
		var b bool
		var c string
		var d int
		fmt.Printf("%v : %T \n",a,a)
		fmt.Printf("%v : %T \n",b,b)
		fmt.Printf("%v : %T \n",c,c)
		fmt.Printf("%v : %T \n",d,d)

	}


}
