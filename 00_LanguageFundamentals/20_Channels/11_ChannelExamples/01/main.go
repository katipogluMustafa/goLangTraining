package main

import (
	"fmt"
	"strconv"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	strbry_cs := make(chan string)
	choco_cs := make(chan string)

	// Factory Methods
	wg.Add(2)
	go makeCakeAndSend(choco_cs, "Chocolate", 3, &wg)
	go makeCakeAndSend(strbry_cs, "Strawberry", 3, &wg)

	// Receiver and packer
	wg.Add(1)
	go receiveCakeAndPack(strbry_cs, choco_cs, &wg)
	wg.Wait()
}

func makeCakeAndSend(cs chan string, flavor string, count int, group *sync.WaitGroup) {
	defer group.Done()
	for i := 1; i <= count; i++ {
		cakeName := flavor + " Cake " + strconv.Itoa(i)
		cs <- cakeName // send a strawberry cake
	}
	close(cs)
}

func receiveCakeAndPack(strbry_cs chan string, choco_cs chan string, group *sync.WaitGroup) {
	defer group.Done()
	strbry_closed, choco_closed := false, false

	for {
		if strbry_closed && choco_closed {
			return
		}
		fmt.Println("Waiting for a new cake...")

		select {
		case cakeName, strbry_ok := <-strbry_cs:
			if !strbry_ok {
				strbry_closed = true
				fmt.Println("...Strawberry channel closed!")
			} else {
				fmt.Println("Received from strawberry channel. Now packing", cakeName)
			}
		case cakeName, choco_ok := <-choco_cs:
			if !choco_ok {
				choco_closed = true
				fmt.Println(" ... Chocolate channel closed!")
			} else {
				fmt.Println("Received from Chocolate channel.  Now packing", cakeName)
			}
		}
	}
}
