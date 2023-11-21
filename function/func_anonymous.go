package main

import "fmt"

type Blacklist func(string) bool

func registerUser(name string, filter Blacklist) {
	if filter(name) {
		fmt.Println("You're blocked", name)
	} else {
		fmt.Println("Welcome", name)
	}
}

func main() {
	blacklist := func(name string) bool {
		return name == "stfu"
	}

	registerUser("Sandy", blacklist)
	registerUser("stfu", blacklist)

}
