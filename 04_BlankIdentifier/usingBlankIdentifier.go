package main

import (
	"net/http"
	"io/ioutil"
	"fmt"
)

func main(){
	res, _ := http.Get("http://www.teknoarktik.com")
	page, _ := ioutil.ReadAll(res.Body)
	res.Body.Close()
	fmt.Printf("%s", page)
}

/*

	Despite the methods return error, we didn't get it this time and we used blank identifier which shows we won't be dealing with error in this code so we just throw it away

*/
