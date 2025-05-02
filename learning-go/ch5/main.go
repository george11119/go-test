package main

import (
	"fmt"
	"os"
)

func fileLen(fileName string) (int, error) {
	bytes, err := os.ReadFile(fileName)
	if err != nil {
		return 0, err
	}

	return len(bytes), nil
}

func prefixer(prefix string) func(string) string {
	prefix_func := func(str string) string {
		return prefix + " " + str
	}
	return prefix_func
}

func main() {
	fileLen, err := fileLen("test1")

	if err != nil {
		panic(err)
	}

	fmt.Println(fileLen)

	helloPrefix := prefixer("Hello")
	fmt.Println(helloPrefix("Bob"))   // should print Hello Bob
	fmt.Println(helloPrefix("Maria")) // should print Hello Maria
}
