package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

func makePerson(firstName string, lastName string, age int) person {
	return person{firstName, lastName, age}
}

func makePersonPointer(firstName string, lastName string, age int) *person {
	return &person{firstName, lastName, age}
}

func updateSlice(strSlice []string, str string) {
	sliceLen := len(strSlice)
	strSlice[sliceLen-1] = str
}

// shouldnt update strSlice
func growSlice(strSlice []string, str string) {
	strSlice = append(strSlice, str)
}

func main() {
	p := makePerson("eren", "yeager", 18)
	p2 := makePersonPointer("eren", "yeager", 18)

	fmt.Println(p)
	fmt.Println(p2)

	mySlice := []string{"aaa", "bbb", "ccc"}
	updateSlice(mySlice, "ddd")
	fmt.Println(mySlice)

	growSlice(mySlice, "eee")
	fmt.Println(mySlice)

	personSlice := make([]person, 0, 10_000_000)
	for i := 0; i < 10_000_000; i++ {
		personSlice = append(personSlice, person{"a", "b", 1})
	}
}
