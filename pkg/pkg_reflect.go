package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name  string `required:"true" max:"100"`
	Email string `required:"true" max:"100"`
}

func main() {
	person := Person{"Sandy", "sandy.hasanudin@icloud.com"}

	fmt.Println(person)
	fmt.Println(reflect.TypeOf(person).Name())

	var refPerson reflect.Type = reflect.TypeOf(person)

	for i := 0; i < refPerson.NumField(); i++ {
		var structField reflect.StructField = refPerson.Field(i)
		fmt.Println(structField.Name, "with type", structField.Type)
		fmt.Println(structField.Tag.Get("required"))
		fmt.Println(structField.Tag.Get("max"))
	}
}
