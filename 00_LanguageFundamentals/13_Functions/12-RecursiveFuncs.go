package main

import "fmt"

func factorial(x int) int {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}
func main() {
	fmt.Println(factorial(5))

	/* Analyze Following Functions, Predict The Outputs */
	fmt.Println(fun(4, 3))
	fun1(25)
	fmt.Println(fun4(11))

	/* What Does the Following Function Do ?*/
	fmt.Println(fun2(5, 6))
	fmt.Println(fun3(3, 2))

	/* Good One :) */
	x := 15
	fmt.Printf("%d", fun5(5, &x))

	/* Answers */
	/*
	 * fun  -- 13
	 * fun1 -- 10011
	 * fun2 -- 30 && Multiplication
	 * fun3 -- 9 && x to the power of y
	 * fun4 -- 5
	 * fun5 -- 8
	 */
}

// Recursion Quiz ~ 0
func fun(x int, y int) int {
	if x == 0 {
		return y
	}
	return fun(x-1, x+y)
}

// Recursion Quiz ~ 1
func fun1(n int) {
	if n == 0 {
		return
	}
	fmt.Printf("%d", n%2)
	fun1(n / 2)
}

// Recursion Quiz ~ 2
func fun2(x int, y int) int {
	if y == 0 {
		return 0
	}
	return x + fun2(x, y-1)
}

// Recursion Quiz ~ 3
func fun3(x int, y int) int {
	if y == 0 {
		return 1
	}
	return fun2(x, fun3(x, y-1)) // Be careful first one fun2
}

// Recursion Quiz ~ 4
func fun4(n int) int {
	if n <= 1 {
		return 1
	}
	if n%2 == 0 {
		return fun4(n / 2)
	}
	return fun4(n/2) + fun4(n/2+1)
}

// Recursion Quiz ~ 5
func fun5(n int, p *int) int {
	var t, f int
	if n <= 1 {
		*p = 1
		return 1
	}
	t = fun5(n-1, p)
	f = t + *p
	*p = t
	return f
}
