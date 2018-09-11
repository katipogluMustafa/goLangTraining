package main

import "fmt"

func fibonacci(out, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case out <- x:
			x, y = y, x+y
		case <-quit: // if any data come to quit chan., then break infinite loop
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	out := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-out) // pop data from out chan. 10 times
		}
		quit <- 0 // then send 0 to quit
	}()

	fibonacci(out, quit)
}
