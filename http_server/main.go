package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type server struct{}

// Student is a struct
type Student struct {
	FirstName string `json:"first_name" bson:"full_name" validate:"required"`
	LastName  string `json:"last_name"`
	Age       int    `json:"age"`
	Class     string `json:"class_name"`
}

var students = []Student{
	{"Phi", "Hoang", 30, "Golang"},
	{"Binh", "Hoang", 12, "Python"},
	{"Yen", "Hoang", 23, "Java"},
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "hello world"}`))
}

func (s *server) GetStudents() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		bs, err := json.Marshal(students)

		if err != nil {
			log.Fatal(bs)
			log.Fatal(err)
			return
		}

		w.Write(bs)
	}
}

func main() {
	fmt.Println(students)

	s := &server{}
	http.Handle("/", s)
	http.Handle("/getStudents", s.GetStudents())
	log.Fatal(http.ListenAndServe(":8080", nil))
}

/* Reference from https://dev.to/moficodes/build-your-first-rest-api-with-go-2gcj

Lets break down this code.
At the top we have our package main all go executable need a main package.

We have our imports. log for logging some error if it happens. net/http because we are writing a rest api.

Then we have a struct called server. It has no fields. We will add a method to this server ServeHTTP and that will satisfy the Handler interface. One thing you will notice in go we don't have to explicitly say the interface we are implementing. The compiler is smart enough to figure that out. In the ServeHTTP method we set httpStatus 200 to denote its the request was a success. We se the content type to application/json so the client understands when we send back json as payload. Finally we write
{"message": "hello world"}
To the response. */
