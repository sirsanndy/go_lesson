package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string
}

type Person struct {
	Name string `required:"true" max:"110"`
	Age  int    `required:"true" max:"110"`
	Addr string `required:"true" max:"255"`
}

func readField(val any) {
	var valueType reflect.Type = reflect.TypeOf(val)
	fmt.Println("Type Name:", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Println(valueField.Name, "with type", valueField.Type)
		fmt.Println("required:", valueField.Tag.Get("required"))
		fmt.Println("max:", valueField.Tag.Get("max"))
	}
}

func main() {
	readField(Sample{"Sandy"})
	readField(Person{"Sandy", 30, ""})
}
