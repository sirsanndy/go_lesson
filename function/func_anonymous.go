package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, filter Blacklist) {
	if filter(name) {
		fmt.Println("You're blocked")
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		if name == "stfu" {
			return true
		} else {
			return false
		}
	}

	registerUser("Sandy", blacklist)
	registerUser("stfu", blacklist)

}
