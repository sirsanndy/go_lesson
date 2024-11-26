package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Println("Test Started")
	m.Run()
	fmt.Println("Test Finished")
}

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

func TestSubTest(t *testing.T) {
	t.Run("Positive Test", func(t *testing.T) {
		name := "Sandy"
		result := SayHello(name)
		assert.Equal(t, result, "Hello "+name, "SayHello doesn't give expected result 'Hello "+name+"'")
	})

	t.Run("Negative Test", func(t *testing.T) {
		name := "Sandy"
		result := SayHello(name)

		if result == "Hello "+"Assertions" {
			t.Fatal("SayHello doesn't give expected result 'Hello " + name + "'")
		}
	})
}
