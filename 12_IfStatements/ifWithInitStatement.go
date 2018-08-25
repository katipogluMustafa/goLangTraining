package main

import "fmt"

func main() {

	b := true
	if food := "Chocolate"; b { // We could also initialize variable
		/* We've checked the b as the if statement	*/
		fmt.Println(food)
	}

}

/*

	Why did we initialized a variable inside if ?
		-We always want to keep the scope tight, so we created food there and the food will only exist inside of the is blocks.
	Let's see a meaningful example of it

	if err:= file.Chmod(0664); err != nil {
		log.Print(err)
		return err
	}

	As you'll see in this example, we have an err but we don't want to move the err all over the places, its scope only the if blocks

*/
