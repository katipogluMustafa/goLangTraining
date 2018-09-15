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
		buckets = HashBucket(scanner.Text(), buckets)
	}
	fmt.Println(buckets[65:123])

	for i := 28; i <= 126; i++ {
	fmt.Printf("%v - %c - %v\n", i, i, buckets[i])
	}
}

func HashBucket(word string, bucket[] int) []int{

	for _ , e := range word {
		bucket[int(e)]++
	}

	return bucket
}
