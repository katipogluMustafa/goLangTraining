package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Println(now)

	formattedTime := now.Format("3:04PM")
	fmt.Println(formattedTime)
	formattedTime = now.Format(time.ANSIC)
	fmt.Println(formattedTime)
}
