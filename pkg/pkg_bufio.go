package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var input = strings.NewReader("Hello!\nMy Name is Sandy")
	var reader = bufio.NewReader(input)
	for {
		line, prefix, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		fmt.Println(prefix, string(line))
	}

	var writer = bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Hello World")
	writer.Flush()
}
