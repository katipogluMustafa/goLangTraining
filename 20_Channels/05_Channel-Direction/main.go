package main

import "fmt"

func main(){
	c := incrementor()
	cSum := puller(c)
	for n := range cSum {
		fmt.Println(n)
	}
}

func incrementor() <-chan int {
	out := make(chan int)
	go func(){
		for i := 0; i < 10; i++ {
			out <- i
		}
		close(out)
	}()
	return out
}

func puller(c <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()

	return out
}

/*
 * puller and incrementor returns channels of it which you can only receive from,so the return would be receive only channel
 * <-chan int	can only be used to receive ints
 * chan<- int	can only be used to send ints
 * chan int		can be used to send and receive ints //bidirectional channel
 */