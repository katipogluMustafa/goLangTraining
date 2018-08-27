package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println(ex(5))

	/* 2 */
	ex2 := func(x int) (int, bool) {
		return x / 2, x % 2 == 0
	}
	fmt.Println(ex2(8))
	fmt.Println( greatest(5,6,7,8,1,-5,2,3,-100) )

	fmt.Println(foo(1,2))
	fmt.Println(foo(1,2,3))
	fmt.Println(foo([]int{1,2,3,4}...))
	fmt.Println(foo())
}

/* 1 */
func ex(x int) (int, bool) {
	return x / 2, x % 2 == 0
}

/* 3 */
func greatest(x ...int) int {
	num := math.MinInt32
	for _, v := range x {
		if v > num {
			num = v
		}
	}
	return num
}
/* 4 */
/*
	Q: What's the value of this expression: (true && false) || (false && true) || !(false && false)?
	A: true
 */

 /* 5 */
 func foo(num ...int) int {
	var sum int
	for _,v := range num{
		sum += v
	}
	return sum
 }
 /* Question 1
  * Write a function which takes an integer.
  * The function will have two returns.
  * The first return should be the argument divided by 2.
  * The second return should be a bool that let’s us know whether or not the argument was even.
  * For example:
  *		half(1) returns (0, false)
  *		half(2) returns (1, true)
  */
/* Question 2
 * Modify the previous program to use a func expression.
 */
/* Question 3
 * Write a function with one variadic parameter that finds the greatest number in a list of numbers.
 */
/* Question 5
 * Write a function, foo, which can be called in all of these ways:
 * 		func main() {
 *			foo(1, 2)
 *			foo(1, 2, 3)
 *			aSlice := []int{1, 2, 3, 4}
 *			foo(aSlice...)
 * 			foo()
 *		}
 */
 /*
  * Find a problem at projecteuler.net then create the solution.
  * Answers are in another repository.
  * https://github.com/katipogluMustafa/eulerProject
  */