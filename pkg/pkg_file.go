package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"time"
)

func createNewFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(message)
	return nil
}

func readFile(name string) (string, error) {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return "", err
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var message string
	for {
		line, _, err := reader.ReadLine()
		message += string(line) + "\n"
		if err == io.EOF {
			break
		}
	}

	return message, nil
}

func appendFile(name string, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(message)
	return nil
}

func main() {
	var fileName string = "sample.log"
	createNewFile(fileName, "this is a sample log")
	appendFile(fileName, string("\n"+time.Now().Format(time.RFC3339)))
	message, _ := readFile(fileName)
	fmt.Println(message)
}
