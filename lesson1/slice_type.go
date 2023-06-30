package main

import "fmt"

func main() {
	var months = [...]string{
		"Jan",
		"Feb",
		"Mar",
		"Apr",
		"May",
		"Jun",
		"Jul",
		"Aug",
		"Sept",
		"Oct",
		"Nov",
		"Dec",
	}

	var days = [...]string{
		"Sunday",
		"Monday",
		"Tuesday",
		"Wednesday",
		"Thursday",
		"Friday",
		"Saturday",
	}

	var slice_one = months[4:7]
	fmt.Println(slice_one)
	fmt.Println(len(slice_one))
	fmt.Println(cap(slice_one))

	months[5] = "Change"
	slice_one[0] = "New Moon"

	fmt.Println(slice_one)
	fmt.Println(months)

	fmt.Println(days)

	var slice_two = days[3:4]
	fmt.Println(slice_two)

	var slice_three = append(slice_two, "Weekend")
	slice_three[0] = "Midweek"
	fmt.Println(slice_three)

	fmt.Println(slice_two)
	fmt.Println(days)

	new_slice := make([]string, 3, 5)

	new_slice[0] = "Muhamad"
	new_slice[1] = "Sandy"
	new_slice[2] = "Hasanudin"

	fmt.Println(new_slice)

	copy_slice := make([]string, len(new_slice), cap(new_slice))
	copy(copy_slice, new_slice)
	fmt.Println(copy_slice)
}
