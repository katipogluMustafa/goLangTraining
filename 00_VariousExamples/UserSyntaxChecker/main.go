package main

import (
	"flag"
	"fmt"
	"goLangTraining/22_UserSyntaxChecker/usernameSyntaxChecker"
)

const UsernameRegex string = `^@?(\w){1,15a}$`

func main() {
	var usernameInput string
	flag.StringVar(&usernameInput, "username", "Gopher", "The GopherFace Username To Check")
	flag.Parse()

	fmt.Println("GopherFace Username Validation Checker v3")
	fmt.Printf("Checking Syntax for username, %s resulted in: %v", usernameInput, validationKit.CheckUsernameSyntax(usernameInput))
}
