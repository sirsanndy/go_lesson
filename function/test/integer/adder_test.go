package integers

import (
	"testing"
)

func TestAdder(t *testing.T) {
	var sum = Add(2, 2)
	var expected = 4
	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}
