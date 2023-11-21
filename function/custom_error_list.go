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

func saveData(id string) []error {
	var err []error

	if id == "" {
		err = append(err, &validationError{Message: "Validation error empty string"})
	}

	if id != "Sandy" {
		err = append(err, &notFoundError{Message: "Data not found"})
	}

	if len(err) > 0 {
		return err
	} else {
		return nil
	}
}

func main() {
	var errs = saveData("Sandy")

	if errs != nil {
		var errVal *validationError
		var errNotFound *notFoundError
		for _, err := range errs {
			if e, ok := err.(*validationError); ok {
				fmt.Println("validation error:", e.Message)
			}

			if e, ok := err.(*notFoundError); ok {
				fmt.Println("not found error:", e.Message)
			}

			if !errors.As(err, &errVal) && !errors.As(err, &errNotFound) {
				fmt.Println("unknown error:", err.Error())
			}
		}
	} else {
		fmt.Println("Success")
	}

}
