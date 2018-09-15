package main

import "fmt"

func main(){
 	c := make(chan int)
	go func(){
		for i := 0; i < 10; i++ {
			c <- i
		}
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

}

/*
 * When a channel is closed, you can no longer put values onto the channel.
 * You can still receive values from the channel but you can't put values onto the channel.
 */