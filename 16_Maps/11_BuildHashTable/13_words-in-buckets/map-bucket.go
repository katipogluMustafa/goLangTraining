package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
)

func main() {
	// get the book adventures of sherlock holmes
	res, err := http.Get("http://www.gutenberg.org/cache/epub/1661/pg1661.txt")
	if err != nil {
		log.Fatal(err)
	}

	// scan the page
	scanner := bufio.NewScanner(res.Body)
	defer res.Body.Close()
	/*
	 * Create a map with a key of int
	 * with a key of string, which will be the word
	 * and a value of int, which will be the number of times the words occurs
	 */
	buckets := make(map[int]map[string]int)
	//Create slices to hold words words
	for i := 0; i < 12; i++ {
		buckets[i] = make(map[string]int)
	}
	// Loop over the words
	for scanner.Scan() {
		word := scanner.Text()
		n := hashBuckett(word, 12)
		buckets[n][word]++
	}
	// Print the words in a bucket
	for k, v := range buckets[6] {
		fmt.Println(v, " \t- ", k)
	}
}
func hashBuckett(word string, buckets int) int {
	var sum int
	for _, v := range word {
		sum += int(v)
	}
	return sum % buckets
}
