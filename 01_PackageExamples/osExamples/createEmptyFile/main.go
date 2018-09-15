package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	var newFile *os.File
	var err error

	newFile, err = os.Create("test.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("test.txt successfully created.")
	defer newFile.Close()
}
