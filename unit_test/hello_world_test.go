package unit_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSayHello(t *testing.T) {
	name := "Sandy"
	result := SayHello(name)
	if result != "Hello "+name {
		t.Error("SayHello doesn't give expected result 'Hello " + name + "'")
	}
}

func TestSayHelloAssertions(t *testing.T) {
	name := "Sandy"
	result := SayHello(name)
	assert.Equal(t, result, "Hello "+name, "SayHello doesn't give expected result 'Hello "+name+"'")
}

func TestSayHelloNegative(t *testing.T) {
	name := "Sandy"
	result := SayHello(name)

	if result == "Hello "+"Assertions" {
		t.Fatal("SayHello doesn't give expected result 'Hello " + name + "'")
	}
}

func TestSayHelloNegativeAssertions(t *testing.T) {
	name := "Sandy"
	result := SayHello(name)
	assert.NotEqual(t, result, "Hello "+"Assertions", "SayHello doesn't give expected result 'Hello "+name+"'")
}
