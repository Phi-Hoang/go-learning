package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Student is a struct
type Student struct {
	FirstName string `json:"first_name" bson:"first_name" validate:"required"`
	LastName  string `json:"last_name" bson:"last_name"`
	Age       int    `json:"age" bson:"age"`
	Class     string `json:"class_name" bson:"class_name"`
}

func main() {
	resp, err := http.Get("http://localhost:8080/getStudents")
	if err != nil {
		log.Fatalln(err)
	}

	// defer is similar with 'finally' in Java
	defer resp.Body.Close()

	// fmt.Println(resp)
	// fmt.Println(resp.Body)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(string(body))

	inputJSON := string(body)
	var outStudents []Student
	err = json.Unmarshal([]byte(inputJSON), &outStudents)
	if err != nil {
		fmt.Println(err)
	}
	for _, st := range outStudents {
		fmt.Printf("%+v\n", st)
	}

	// Initialising and connecting
	// ========================================================================================

	// create a new timeout context
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// create a mongo client
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://phihdn:4Wi3CuK4FuknQu5f@cluster0-5kahe.gcp.mongodb.net/test?retryWrites=true&w=majority"))
	if err != nil {
		log.Fatal(err)
	}

	// disconnect from mongo
	defer client.Disconnect(ctx)

	// select collection from database
	col := client.Database("go-learning").Collection("students")

	// InsertOne
	// ========================================================================================
	{
		for _, st := range outStudents {
			res, err := col.InsertOne(ctx, st)
			if err != nil {
				log.Fatal(err)
			}
			fmt.Printf("%s inserted id: %s\n", st.FirstName, res.InsertedID.(primitive.ObjectID).Hex())
		}
	}
}

// reference for http request https://medium.com/@masnun/making-http-requests-in-golang-dd123379efe7
// reference for mongodb https://vkt.sh/go-mongodb-driver-cookbook/
// https://blog.ruanbekker.com/blog/2019/04/17/mongodb-examples-with-golang/
