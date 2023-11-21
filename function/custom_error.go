package main

import (
	"errors"
	"fmt"
)

type validationError struct {
	Message string
}

func (v *validationError) Error() string {
	return v.Message
}

type notFoundError struct {
	Message string
}

func (n *notFoundError) Error() string {
	return n.Message
}

func SaveData(id string, data any) error {
	err := errors.Join(checkValidation(id), checkNotFound(id))
	if err != nil {
		return err
	}

	return nil
}

func checkValidation(id string) error {
	if id == "" {
		return &validationError{Message: "Validation error empty string"}
	}

	return nil
}

func checkNotFound(id string) error {
	if id != "Sandy" {
		return &notFoundError{Message: "Data not found"}
	}

	return nil
}

func main() {
	var err = SaveData("Say", nil)
	var errVal *validationError
	var errNotFound *notFoundError
	if err != nil {
		if errors.As(err, &errVal) {
			fmt.Println("validation error:", errVal.Message)
		}

		if errors.As(err, &errNotFound) {
			fmt.Println("not found error:", errNotFound.Message)
		}

		if !errors.Is(err, errVal) && !errors.Is(err, errNotFound) {
			fmt.Println("unknown error:", err.Error())
		}
	} else {
		fmt.Println("Success")
	}
}
