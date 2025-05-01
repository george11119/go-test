package main

import (
	"fmt"
)

func main() {
	// ex1
	greetings := []string{"Hello", "Hola", "नमस्कार", "こんにちは", "Привіт"}
	subslice1 := greetings[:2]
	subslice2 := greetings[1:4]
	subslice3 := greetings[3:]

	fmt.Println(greetings)
	fmt.Println(subslice1)
	fmt.Println(subslice2)
	fmt.Println(subslice3)

	// ex2
	message := "Hi 👩 and 👨"
	messageRuneSlice := []rune(message)
	fmt.Println(string(messageRuneSlice[3]))

	// ex3
	type employee struct {
		firstName string
		lastName  string
		id        int
	}

	eren := employee{
		"eren",
		"yeager",
		1,
	}

	armin := employee{
		firstName: "armin",
		lastName:  "arlert",
		id:        2,
	}

	var mikasa employee

	mikasa.firstName = "mikasa"
	mikasa.lastName = "ackerman"
	mikasa.id = 3

	fmt.Println(eren)
	fmt.Println(armin)
	fmt.Println(mikasa)
}
