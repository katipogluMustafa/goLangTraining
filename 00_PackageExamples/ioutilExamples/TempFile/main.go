package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	content := []byte("temporary file's content is this sentence")
	tmpFile, err := ioutil.TempFile(".", "example")
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(tmpFile.Name())

	if _, err := tmpFile.Write(content); err != nil {
		log.Fatal(err)
	}
	/* See the current directory*/
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}

	/* See the contents of the tempFile*/
	contnt, err := ioutil.ReadFile(tmpFile.Name())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(contnt))

	if err := tmpFile.Close(); err != nil {
		log.Fatal(err)
	}
}
