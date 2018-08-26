package main

import (
	"fmt"
	"strings"
)

func main() {

	/* First Exercise */
	/*
	 * Create a program that prints to the terminal asking for a user to enter their name.
	 * Use fmt.Scan to read a user’s name entered at the terminal.
	 * Print “Hello <NAME>” with <NAME> replaced with what the user entered at the terminal.
	 */
	var name string

	fmt.Print("Name: ")
	fmt.Scan(&name)

	fmt.Printf("Hello %s !", strings.Title(name))
	fmt.Println()

	/* Second Exercise */
	/*
	 * Create a program that prints to the terminal asking for a user to enter a small number and a larger number.
	 * Print the remainder of the larger number divided by the smaller number.
	 */
	var num1, num2 int

	fmt.Print("First Number: ")
	fmt.Scan(&num1)
	fmt.Print("Second Number: ")
	fmt.Scan(&num2)

	if num1 > num2 {
		fmt.Println("Remainder of Large divided by small : ", num1%num2)
	} else {
		fmt.Println("Remainder of Large divided by small : ", num2%num1)
	}
	fmt.Println()

	/* Third Exercise */
	/*
	 * Print all of the even numbers between 0 and 100.
	 */

	var numRange int

	fmt.Print("Range: ")
	fmt.Scan(&numRange)
	if numRange >= 1 {
		for i := 0; i <= numRange; i++ {
			if i%2 == 0 {
				fmt.Print(i, " ")
			}
		}
	}

	fmt.Println()

	/* Fourth Exercise */
	/*
	 * Write a program that prints the numbers from 1 to 100.
	 * But for multiples of three print "Fizz" instead of the number and for the multiples of five print "Buzz".
	 * For numbers which are multiples of both three and five print "FizzBuzz".
	 */

	for i := 1; i <= 105; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Print("FizzBuzz \n")
		} else if i%3 == 0 {
			fmt.Print("Fizz ")
		} else if i%5 == 0 {
			fmt.Print("Buzz ")
		} else {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()

	/* Fifth Exercise */
	/*
	 * If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9.
	 * The sum of these multiples is 23. Find the sum of all the multiples of 3 or 5 below 1000.
	 */

	var sum int
	for i := 0; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println("Sum: ", sum)
}
