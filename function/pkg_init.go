package main

import (
	"fmt"
	database "function/database"
)

func main() {
	fmt.Println(database.GetDatabase())
}
