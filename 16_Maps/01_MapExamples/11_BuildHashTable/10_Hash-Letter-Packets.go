package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
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
	//loop over the words
	for scanner.Scan() {
		n := HashBuckett(scanner.Text())
		buckets[n]++
	}
	fmt.Println(buckets[65:123])
}
func HashBuckett(word string) int {
	return int(word[0])
}
