package main

import (
	"database/sql"
	"encoding/json"
	"github.com/gorilla/mux"
	"net/http"
)

func GetPeople(w http.ResponseWriter, r *http.Request) {
	rows, err := db.Query("SELECT p.id, p.firstname, p.lastname, a.city, a.state FROM person p LEFT JOIN address a on p.id = a.idperson")
	if err != nil {
		panic(err)
	}

	var people []Person
	var person Person

	for rows.Next() {
		var address Address

		err = rows.Scan(&person.ID, &person.Firstname, &person.Lastname, &address.City, &address.State)
		if err != nil {
			panic(err)
		}

		person.Address = &address

		people = append(people, person)
	}

	err = rows.Close()
	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(people); err != nil {
		panic(err)
	}
}

func GetPerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var person Person
	var address Address

	row := db.QueryRow("SELECT * FROM person WHERE id = $1", params["id"])

	err = row.Scan(&person.ID, &person.Firstname, &person.Lastname)
	switch err {
	case sql.ErrNoRows:
	case nil:
		row := db.QueryRow("SELECT city, state FROM address WHERE idperson = $1", person.ID)
		err = row.Scan(&address.City, &address.State)
		switch err {
		case sql.ErrNoRows:
		case nil:
			person.Address = &address
		default:
			panic(err)
		}
	default:
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(&person); err != nil {
		panic(err)
	}
}

func CreatePerson(w http.ResponseWriter, r *http.Request) {
	var person Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		panic(err)
	}

	stmt, err := db.Prepare("INSERT INTO person(id, firstname, lastname) values(?,?,?)")
	if err != nil {
		panic(err)
	}

	res, _ := stmt.Exec(person.ID, person.Firstname, person.Lastname)

	if res != nil && person.Address != nil {
		stmt, err := db.Prepare("INSERT INTO address(idperson, city, state) values(?,?,?)")
		if err != nil {
			panic(err)
		}

		_, err = stmt.Exec(person.ID, person.Address.City, person.Address.State)
		if err != nil {
			panic(err)
		}
	}

	if err := json.NewEncoder(w).Encode(person); err != nil {
		panic(err)
	}
}

func DeletePerson(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	stmt, err := db.Prepare("delete from person where id=?")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err)
	}

	stmt, err = db.Prepare("delete from address where idperson=?")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err)
	}

	if err := json.NewEncoder(w).Encode(true); err != nil {
		panic(err)
	}
}
