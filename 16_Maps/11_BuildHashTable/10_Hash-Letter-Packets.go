package main

import (
	"bufio"
	"log"
	"net/http"
)

func main(){
	// get the book moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()

	// Set the split function for the scanning operation
	scanner.Split(bufio.ScanWords)

	// Create slice to hold counts
	buckets := make([]int, 200)

}
