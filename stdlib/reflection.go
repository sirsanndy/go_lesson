package main

import (
	"fmt"
	"reflect"
)

type Sample struct {
	Name string
}

type Person struct {
	Name string
	Age  int
	Addr string
}

func readField(val any) {
	var valueType reflect.Type = reflect.TypeOf(val)
	fmt.Println("Type Name:", valueType.Name())
	for i := 0; i < valueType.NumField(); i++ {
		var valueField reflect.StructField = valueType.Field(i)
		fmt.Println(valueField.Name, "with type", valueField.Type)
	}
}

func main() {
	readField(Sample{"Sandy"})
	readField(Person{"Sandy", 30, ""})
}
