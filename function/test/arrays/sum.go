package main

func Sum(number []int) int {
	var sum = 0
	for _, n := range number {
		sum += n
	}

	return sum
}
