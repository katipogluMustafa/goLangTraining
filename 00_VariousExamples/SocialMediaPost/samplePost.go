package main

import (
	"fmt"
	"goLangTraining/23_SocialMediaPost/socialMedia"
)

func main(){
	myPost := socialMedia.NewPost("MustafaKatipoglu", socialMedia.Moods["thrilled"], "Go is awesome", "Check out the Go web site!", "https://golang.org", "", "", []string{"go","golang","programming language"})
	fmt.Printf("myPost: %+v\n\n", myPost)
	fmt.Printf("myPost: %v\n", myPost)
}
/*
 * when using printf, %+v prints the field names of the struct when %v just prints the values
 */