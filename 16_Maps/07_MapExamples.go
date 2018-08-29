package main

import (
	"net/http"
	"log"
	"bufio"
	"fmt"
	"os"
)

func main(){
	res, err := http.Get("https://raw.githubusercontent.com/dwyl/english-words/master/words.txt")
	if err != nil {
		log.Fatalln(err)
	}

	words := make(map[string]string)

	sc := bufio.NewScanner(res.Body)
	sc.Split(bufio.ScanWords)

	for sc.Scan() {
		words[sc.Text()] = ""
	}

	if err := sc.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input: ", err)
	}

	i := 0

	for k, _ := range words {
		fmt.Println(k)
		if i == 1000 {
			break
		}
		i++
	}

}
