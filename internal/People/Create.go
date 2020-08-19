package People

import (
	"encoding/json"
	"github.com/dfrnks/go-rest-api/internal/Database"
	"github.com/dfrnks/go-rest-api/internal/Models"
	"net/http"
)

func Create(w http.ResponseWriter, r *http.Request) {
	var person Models.Person
	if err := json.NewDecoder(r.Body).Decode(&person); err != nil {
		panic(err)
	}

	stmt, err := Database.Connection().Prepare("INSERT INTO person(id, firstname, lastname) values(?,?,?)")
	if err != nil {
		panic(err)
	}

	res, _ := stmt.Exec(person.ID, person.Firstname, person.Lastname)

	if res != nil && person.Address != nil {
		stmt, err := Database.Connection().Prepare("INSERT INTO address(idperson, city, state) values(?,?,?)")
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
