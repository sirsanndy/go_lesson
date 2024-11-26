package unit_test

import (
	"testing"
)

func TestSayHello(t *testing.T) {
	name := "Sandy"
	result := SayHello(name)
	if result != "Hello "+name {
		t.Error("SayHello doesn't give expected result 'Hello " + name + "'")
	}
}

func TestSayHelloNegative(t *testing.T) {
	name := "Sandy"
	result := SayHello(name)

	if result == "Hello "+"Eko" {
		t.Fatal("SayHello doesn't give expected result 'Hello " + name + "'")
	}
}
