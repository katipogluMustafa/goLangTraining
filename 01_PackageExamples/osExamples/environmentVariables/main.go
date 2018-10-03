package main

import (
	"fmt"
	"os"
)

// Getenv retrieves the value of the environment variable named by the key.
// It returns the value, which will be empty if the variable is not present.
// To distinguish between an empty value and an unset value, use LookupEnv.
func main(){
	/* os.Getenv */
	fmt.Printf("Home: %s \nGo-Root: %s", os.Getenv("HOME"), os.Getenv("GOROOT"))

	/* os.LookupEnv */
	Show := func(key string) string{
		val, ok := os.LookupEnv(key)

		if !ok {
			return "\nThe Environment Variable does not exist."
		}

		return val
	}
	fmt.Println(Show("USER"))
	fmt.Println(Show("HOME"))


}
