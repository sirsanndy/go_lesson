package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("Validation error")
	NotFoundError   = errors.New("Not found error")
)

func GetById(id string) error {
	var err error
	if id == "" {
		err = errors.Join(ValidationError)
	}

	if id != "Sandy" {
		err = errors.Join(err, NotFoundError)
	}

	if err != nil {
		return err
	}

	return nil
}

func main() {
	err := GetById("")

	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Error Validation :", ValidationError.Error())
		}

		if errors.Is(err, NotFoundError) {
			fmt.Println("Error Not Found :", NotFoundError.Error())
		}

		if !errors.Is(err, ValidationError) && !errors.Is(err, NotFoundError) {
			fmt.Println("unknown error:", err.Error())
		}
	} else {
		fmt.Println("Success")
	}
}
