package main

import (
	"net/http"
	"log"
	"bufio"
	"fmt"
)

func main(){
	// get the book moby dick
	res, err := http.Get("http://www.gutenberg.org/files/2701/old/moby10b.txt")
	if err != nil {
		log.Fatalln(err)
	}

	// scan  the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	// Set the split function for the scanning operation.
	scanner.Split(bufio.ScanWords)
	// Create slice to hold counts
	buckets := make([]int, 200)
	// Loop over the words
	for scanner.Scan() {
		n := HashBucket1(scanner.Text(), 12)
		buckets[n]++
	}
	fmt.Println(buckets)

}

func HashBucket1(word string, buckets int) int{

	letter := int(word[0])
	bucket := letter % buckets

	return bucket
}
