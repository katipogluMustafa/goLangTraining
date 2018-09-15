package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("Contains:", strings.Contains("test", "es"))
	fmt.Println("Count:", strings.Count("testt", "t"))
	fmt.Println("HasPrefix:", strings.HasPrefix("test", "te"))
	fmt.Println("HasSuffix:", strings.HasSuffix("test", "st"))
	fmt.Println("Index of e:", strings.Index("test", "e"))
	fmt.Println("Index of t:", strings.Index("test", "t"))
	fmt.Println("Join:", strings.Join([]string{"aa", "bb"}, ","))
	fmt.Println("Repeat:", strings.Repeat("-repeat-", 5))
	fmt.Println("Replace:", strings.Replace("foo", "o", "0", -1))
	fmt.Println("Replace:", strings.Replace("foo", "o", "0", 1))
	fmt.Println("Split:", strings.Split("a-b-c-d-e", "-"))
	fmt.Println("ToLower:", strings.ToLower("TESt"))
	fmt.Println("ToUpper:", strings.ToUpper("tesT"))

	fmt.Println("Len:", len("hello"))
	fmt.Println("Char:", "hello"[0])
}
