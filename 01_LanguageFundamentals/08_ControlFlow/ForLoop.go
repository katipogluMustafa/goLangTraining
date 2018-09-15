package main

import "fmt"

func main(){

	for i:=1;i<=10;i++{
		fmt.Printf("i: %d ", i)
	}

	fmt.Println()

	i := 0
	for i < 10{
		fmt.Printf("i: %d ", i)
		i++
	}

	fmt.Println()

	j :=0
	for {
		j++
		if j > 10{
			break
		}
		if j == 5 {
			continue
		}
		fmt.Printf("j: %d ", j)
	}

	for x := 1; x < 10; x++{
		for y := x; y < 10; y++{
			fmt.Printf("x : %d ~ y : %d\n",x,y)
		}
	}

}
