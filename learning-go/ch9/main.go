package main

import (
	"errors"
	"fmt"
)

var ErrInvalidID = fmt.Errorf("ID is invalid")

func returnInvalidIDErr() error {
	return fmt.Errorf("haha: %w", ErrInvalidID)
}

type Employee struct {
	FirstName string
	LastName  string
	Company   string
}

type EmptyFieldErr struct {
	Message string
}

func (e EmptyFieldErr) Error() string {
	return e.Message
}

func validateEmployee(e Employee) error {
	errorSlice := []error{}
	if len(e.FirstName) == 0 {
		errorSlice = append(errorSlice, EmptyFieldErr{"FirstName cannot be empty"})
	}
	if len(e.LastName) == 0 {
		errorSlice = append(errorSlice, EmptyFieldErr{"LastName cannot be empty"})
	}
	if len(e.Company) == 0 {
		errorSlice = append(errorSlice, EmptyFieldErr{"Company cannot be empty"})
	}
	if len(errorSlice) != 0 {
		return errors.Join(errorSlice...)
	}
	return nil
}

func main() {
	err := returnInvalidIDErr()

	if errors.Is(err, ErrInvalidID) {
		fmt.Println(err)
	}

	err2 := validateEmployee(Employee{"eren", "", ""})
	var emptyFieldErr = EmptyFieldErr{}
	if errors.As(err2, &emptyFieldErr) {
		fmt.Println(err2)
	}
}
