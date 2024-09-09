package main

import (
	"bufio"
	"os"
)

func main() {
	var writer = bufio.NewWriter(os.Stdout)
	_, _ = writer.WriteString("Hello World")
	writer.Flush()
}
