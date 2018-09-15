package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	First       string
	Last        string
	Age         int
	notExported int
}
type Monkey struct {
	alias    string
	age      int
	Exported int
}

func main() {
	p1 := Person{"James", "Bond", 20, 107}
	bs, _ := json.Marshal(p1)
	fmt.Println(bs)
	fmt.Printf("%T \n", bs)
	fmt.Println(string(bs)) // see, not exported field is not in the output

	m1 := Monkey{"LittleMonkey", 1, 10}
	ms, _ := json.Marshal(m1)
	fmt.Println(ms)
	fmt.Printf("%T \n", ms)
	fmt.Println(string(ms)) // see, alias and age fields is not shown
}

/*
 * JSON -- JavaScript Object Notation
 */

/* func Marshal(v interface{})([]byte, error)
 * Marshal returns the JSON encoding of v.
 */
