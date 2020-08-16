package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {
	if err := json.NewEncoder(w).Encode(people); err != nil {
		panic(err)
	}
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for _, item := range people {
		if item.ID == params["id"] {
			if err := json.NewEncoder(w).Encode(item); err != nil {
				panic(err)
			}
			return
		}
	}

	if err := json.NewEncoder(w).Encode(&Person{}); err != nil {
		panic(err)
	}
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		panic(err)
	}

	for _, item := range people {
		if item.ID == person.ID {
			if err := json.NewEncoder(w).Encode(item); err != nil {
				panic(err)
			}
			return
		}
	}

	people = append(people, person)

	if err := json.NewEncoder(w).Encode(person); err != nil {
		panic(err)
	}
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range people {
		if item.ID == params["id"] {
			people = append(people[:index], people[index+1:]...)
			break
		}
	}

	if err := json.NewEncoder(w).Encode(people); err != nil {
		panic(err)
	}
}
