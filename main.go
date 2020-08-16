package main

import (
	"log"
	"net/http"
)

var people []Person

func main() {
	router := NewRouter()

	people = append(people, Person{
		ID:        "1",
		Firstname: "John",
		Lastname:  "Doe",
		Address: &Address{
			City:  "City X",
			State: "State X",
		},
	})

	people = append(people, Person{
		ID:        "2",
		Firstname: "Koko",
		Lastname:  "Doe",
		Address: &Address{
			City:  "City Z",
			State: "State Y",
		},
	})

	people = append(people, Person{
		ID:        "3",
		Firstname: "Francis",
		Lastname:  "Sunday",
	})

	log.Fatal(http.ListenAndServe(":8000", router))
}
