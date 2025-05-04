package main

import (
	_ "embed"
	"fmt"
	"os"
)

//go:embed english_rights.txt
var englishRights string

//go:embed french_rights.txt
var frenchRights string

//go:embed chinese_rights.txt
var chineseRights string

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments")
		os.Exit(1)
	}
	language := os.Args[1]

	switch language {
	case "english":
		fmt.Println("???")
		fmt.Println(englishRights)
	case "french":
		fmt.Println(frenchRights)
	case "chinese":
		fmt.Println(chineseRights)
	default:
		fmt.Println("Not a choice")
	}
}
