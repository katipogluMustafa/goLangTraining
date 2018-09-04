package main

import "fmt"

func main(){
	c := make(chan int)

	go func(){
		for  i := 0; i < 10; i++{
			c <- i
		}
		close(c) // if you are gonna use range, you need to close the channel!!!
	}()

	for i := range c{
		fmt.Println(i)
	}

}

//Why does this only print zero ?