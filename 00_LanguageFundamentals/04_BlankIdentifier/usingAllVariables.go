package main

import (
	"net/http"
	"log"
	"io/ioutil"
	"fmt"
)

func main(){
	res,err := http.Get("http://www.teknoarktik.com")
	if err != nil {
		log.Fatal(err)
	}
	page, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", page)
}
/*
	In the code, you need to use every variable, every identifier, everything you declare.
The Blank Identifier is allow you to tell computer you aren't using something

*/
