package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func main() {
	var input = strings.NewReader("This is a short sentence\nwith new line")
	var reader = bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(string(line))
	}
}
