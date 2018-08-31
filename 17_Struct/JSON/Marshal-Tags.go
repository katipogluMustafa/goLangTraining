package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First string
	Last  string `json:"-"`
	Age   int    `json:"wisdom score"`
}

func main() {
	p1 := Person{"James", "Bond", 20}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Println(string(bs)) // çıktıdan görüldüğü üzere wisdom score 20
}
