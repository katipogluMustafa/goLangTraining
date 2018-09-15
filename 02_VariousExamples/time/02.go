package main

import (
	"fmt"
	"time"
)

func main() {
	then := time.Date(1998, 7, 15, 7, 5, 0, 0, time.Local)
	fmt.Println(then)
	fmt.Println(then.Year())
	fmt.Println(then.Month())
	fmt.Println(then.Day())
	fmt.Println(then.Hour())
	fmt.Println(then.Second())
	fmt.Println(then.Nanosecond())
	fmt.Println(then.Location())
	fmt.Println(then.Weekday())

	fmt.Println()

	now := time.Now()
	fmt.Println(then.Before(now))
	fmt.Println(then.After(now))
	fmt.Println(then.Equal(now))

	fmt.Println()

	diff := now.Sub(then)
	fmt.Println(diff)
	fmt.Println(diff.Hours())
	fmt.Println(diff.Minutes())
	fmt.Println(diff.Nanoseconds())
	fmt.Println(diff.Seconds())

	fmt.Println()

	fmt.Println(then.Add(diff))
	fmt.Println(then.Add(-diff))
}
