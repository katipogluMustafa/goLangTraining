package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(greet2("mustafa", "katipoğlu"))
}
func greet2(name string, surname string) (s string) {
	s = fmt.Sprintf("Selamu Aleykum %s %s", strings.Title(name), strings.Title(surname))
	return
}
