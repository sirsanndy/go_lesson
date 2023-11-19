package main

import "fmt"

type Person struct {
	Name, Address string
	Age           int16
}

func main() {
	var person Person = Person{"Sandy", "Bogor", 29}
	var person1 *Person = &person

	person1.Address = "Batu"
	*person1 = Person{"Ilham", "Yogyakarta", 19}
	fmt.Println(person)
	fmt.Println(person1)
}
