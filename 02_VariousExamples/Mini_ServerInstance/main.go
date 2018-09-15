package main

import (
	"fmt"
	"goLangTraining/00_VariousExamples/Mini_ServerInstance/validationKit"
	"net/http"
)

func helloGopherHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello Gopher!")
}

func checkUsernameSyntaxHandler(w http.ResponseWriter, r *http.Request) {

	var usernameSyntaxResult bool
	username := r.URL.Query().Get("username")

	if username == "" {
		http.Error(w, "Username not provided!", http.StatusInternalServerError)
	} else {
		usernameSyntaxResult = validationKit.CheckUsernameSyntax(username)
		fmt.Fprintf(w, "Syntax Check Result fo %v is %v", username, usernameSyntaxResult)
	}
}

func main() {
	http.HandleFunc("/hello-gopher", helloGopherHandler)
	http.HandleFunc("/username-syntax-checker", checkUsernameSyntaxHandler)
	http.ListenAndServe(":8080", nil)
}

/*
 * after running main.go, you can see the outputs
 * http://localhost:8080/username-syntax-checker?username=mu@stafa
 * http://localhost:8080/hello-gopher
 */
