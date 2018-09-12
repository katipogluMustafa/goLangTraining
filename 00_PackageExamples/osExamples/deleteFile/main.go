package main

import (
	"fmt"
	"log"
	"os"
	"sync"
)

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	/* Create a File */
	go func() {
		newFile, err := os.Create("test1.txt")
		check(err)
		defer newFile.Close()
		fmt.Println("test.txt successfully created.")
		wg.Done()
	}()
	wg.Wait()

	wg.Add(1)
	/* Remove the file*/
	go func() {
		err := os.Remove("test1.txt")
		check(err)
		fmt.Println("test.txt successfully deleted.")
		wg.Done()
	}()
	wg.Wait()
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
