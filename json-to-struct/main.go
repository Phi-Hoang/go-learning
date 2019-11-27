package main

import (
	"encoding/json"
	"fmt"
)

// Student type
type Student struct {
	FirstName string `json:"first_name" bson:"full_name" validate:"required"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Class     string `json:"class_name"`
}

// GetFullName returns the fullname of a student
func (s Student) GetFullName() string {
	return s.FirstName + " " + s.LastName
}

func main() {
	inputJSON := `[
		{"first_name": "Victor", "last_name": "Nguyen", "age": 100, "class_name":"golang"},
		{"first_name": "Anh", "last_name": "Dinh", "age":200, "class_name":"golang"}
		]`
	var outStudents []Student
	err := json.Unmarshal([]byte(inputJSON), &outStudents)
	if err != nil {
		fmt.Println(err)
	}
	for _, st := range outStudents {
		fmt.Printf("%+v\n", st)
	}

	// struct -> json string
	bs, err := json.Marshal(outStudents)

	if err != nil {
		fmt.Println(bs)
		return
	}

	jsonString := string(bs)

	fmt.Println(jsonString)
}
