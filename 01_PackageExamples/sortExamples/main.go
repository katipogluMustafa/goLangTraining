/* how to walk through a map by alphabetical key order */
package main

import (
	"fmt"
	"sort"
)

type user struct {
	name    string
	surname string
}

func main() {
	users := map[string]user{
		"Roy":     {"Rob", "Roy"},
		"Ford":    {"Henry", "Ford"},
		"Mouse":   {"Mickey", "Mouse"},
		"Jackson": {"Michael", "Jackson"},
	}

	// Pull the keys from the map
	var keys []string
	for key := range users {
		keys = append(keys, key)
	}

	// sort the keys alphabetically
	sort.Strings(keys)

	// Walk through the keys and pull each value from the map
	for _, key := range keys {
		fmt.Println(key, users[key])
	}
}
