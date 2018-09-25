package main

import (
	"fmt"
	"log"
	"os/exec"
)

func main() {
	out, err := exec.Command("echo", "hello", "world").Output()

	if err != nil {
		log.Fatalln(err)
	} else {
		fmt.Println("%s\n", out)

		for _, v := range out {
			fmt.Printf("%s ", string(v))
		}
	}

}
