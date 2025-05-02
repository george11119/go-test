package main

import (
	"errors"
	"fmt"
	"strconv"
)

func add(a, b int) (int, error) {
	return a + b, nil
}

func sub(a, b int) (int, error) {
	return a - b, nil
}

func mul(a, b int) (int, error) {
	return a * b, nil
}

func div(a, b int) (int, error) {
	if b == 0 {
		err := errors.New("cannot divide by 0")
		return 0, err
	}
	return a / b, nil
}

func mod(a, b int) (int, error) {
	if b == 0 {
		err := errors.New("cannot divide by 0")
		return 0, err
	}
	return a % b, nil
}

var operatorsMap = map[string]func(int, int) (int, error){
	"+": add,
	"-": sub,
	"*": mul,
	"/": div,
	"%": mod,
}

func parseExpression(expression []string) (int, error) {
	if len(expression) != 3 {
		err := errors.New("expression must be of length 3")
		return 0, err
	}

	val1, val1Err := strconv.Atoi(expression[0])
	if val1Err != nil {
		err := errors.New("first operand is invalid")
		return 0, err
	}

	val2, val2Err := strconv.Atoi(expression[2])
	if val2Err != nil {
		err := errors.New("second operand is invalid")
		return 0, err
	}

	operator := expression[1]

	opFunc, ok := operatorsMap[operator]
	if !ok {
		err := errors.New("invalid operator")
		return 0, err
	}

	result, err := opFunc(val1, val2)
	if err != nil {
		return 0, err
	}

	return result, nil
}

func main() {
	expressions := [][]string{
		{"2", "+", "3"},
		{"2", "-", "3"},
		{"2", "*", "3"},
		{"2", "/", "3"},
		{"2", "%", "3"},
		{"two", "+", "three"},
		{"5"},
		{"10", "/", "0"},
		{"0", "/", "10"},
	}

	for _, v := range expressions {
		result, err := parseExpression(v)

		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(result)
		}
	}
}
