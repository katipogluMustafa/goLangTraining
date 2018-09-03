package main

import (
	"fmt"
	"time"
)

func main(){
	c := make(chan int)

	go func(){
		for i := 0; i < 10; i++{
			c <- i
		}
	}()
	/*
	 * When you put something on an unbuffered channel, this code blocked until that value taken off the unbuffered channel.
	 * One is waiting for the other until the other is ready.
	 */
	go func() {
		for {
			fmt.Println(<-c)
		}
	}()
	time.Sleep(time.Second)
}

