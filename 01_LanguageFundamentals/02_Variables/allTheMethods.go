package main

import "fmt"

func main(){

		/* Declare Many Variables At Once */
	{
		var a,b,c,d int
		fmt.Println(a,b,c,d)
	}
		/* Initialize Many Variables At Once */
	{
		var a,b,c,d int = 1,2,3,4
		fmt.Println(a,b,c,d)
	}
		/* Infer Type */
	{
		var a,b,c,d = 5,6,7,8
		fmt.Println(a,b,c,d);
	}
		/*  Infer mixed up types */
	{
		var a,b,c,d = 9, true, "hey", 5.2
		fmt.Println(a,b,c,d)
	}
		/* Initialize shorthand */
	{
		a,b,c := 1,false,"hy"
		d := true
		o := 'e'
		fmt.Println(a,b,c,d,o)
	}


}
