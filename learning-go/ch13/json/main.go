package main

import (
	"encoding/json"
	"fmt"
)

type Employee struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
}

func main() {
	jsonStr := `
	{
		"employees":[
			{"firstName":"John", "lastName":"Doe"},
			{"firstName":"Anna", "lastName":"Smith"},
			{"firstName":"Peter", "lastName":"Jones"}
		]
	}`

	var m struct {
		Employees []Employee `json:"employees"`
	}
	err := json.Unmarshal([]byte(jsonStr), &m)
	if err != nil {
		panic(err)
	}

	fmt.Println(m)

	b, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
