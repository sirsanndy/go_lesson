package main

import (
	"fmt"
	database "function/database"
	_ "function/internal"
)

func main() {
	fmt.Println(database.GetDatabase())
}
