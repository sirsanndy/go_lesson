package unit_test

import (
	"fmt"
	"testing"
)

func TestSayHello(t *testing.T) {
	name := "Sandy"
	result := SayHello(name)
	fmt.Println(result)
	if result != "Hello "+name {
		panic("Result should be 'Hello " + name + "'")
	}
}
