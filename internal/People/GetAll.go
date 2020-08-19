package People

import (
	"encoding/json"
	"github.com/dfrnks/go-rest-api/internal/Database"
	"github.com/dfrnks/go-rest-api/internal/Models"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request) {
	rows, err := Database.Connection().Query("SELECT p.id, p.firstname, p.lastname, a.city, a.state FROM person p LEFT JOIN address a on p.id = a.idperson")
	if err != nil {
		panic(err)
	}

	var people []Models.Person
	var person Models.Person

	for rows.Next() {
		var address Models.Address

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
