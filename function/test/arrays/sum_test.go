package main

import "testing"

func TestSum(t *testing.T) {
	var number = []int{1, 2, 3, 4, 5}

	var got = Sum(number)
	var want = 15
	if got != want {
		t.Errorf("got '%d' want '%d' give, '%v'", got, want, number)
	}
}
