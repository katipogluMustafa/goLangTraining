package main

import (
	"fmt"
	"sync"
)

func main(){
	c := make(chan int)

	var wg sync.WaitGroup

	go func(){
		wg.Add(1)
		defer wg.Done()
		for i := 0; i < 10; i++ {
			c <- i
		}
	}()

	go func() {
		wg.Add(1)
		defer wg.Done()
		for i := 0; i < 10; i++{
			c <- i
		}
	}()

	go func() {
		wg.Wait()
		close(c)
	}()

	for n := range c {
		fmt.Println(n)
	}

}