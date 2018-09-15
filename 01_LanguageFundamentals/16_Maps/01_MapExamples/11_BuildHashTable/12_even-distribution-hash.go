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
		n := HashBucketttt(scanner.Text(), 12)
		buckets[n]++
	}
	fmt.Println(buckets[65:123])
}
func HashBucketttt(word string, buckets int) int {
	var sum int

	for _, v := range word {
		sum += int(v)
	}

	return sum % buckets
}
