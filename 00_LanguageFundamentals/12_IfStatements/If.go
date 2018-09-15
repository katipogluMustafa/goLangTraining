package main

import (
	"fmt"
	"math"
)

func main() {

	/* Classic If */

	if true {
		fmt.Println("This runs")
	}

	/* Use not operator*/

	if !true {
		fmt.Println("This doesn't run")
	}

	/* Normal Condition Checking */

	if 5 > math.Pi {
		fmt.Println("If true , This runs")
	} else {
		fmt.Println("Else, This runs")
	}

	/* If - Else*/

	if len("mustafa") > len("katipoglu") {
		fmt.Println("not true, this won't run")
	} else {
		fmt.Println("The condition return false so this else block runs")
	}

	/* If - Else if - Else Structure */

	if false {
		fmt.Println("First Statement")
	} else if false {
		fmt.Println("Second Statement")
	} else if true {
		fmt.Println("Third Statement, This runs since it is the one which evaluates true for the first time")
	} else {
		fmt.Println("Fourth Statement")
	}

	/* If in action */
	for i := 1; i <= 15; i++ {
		if i%3 == 0 {
			fmt.Print(i, " ")
		}
	}
}
