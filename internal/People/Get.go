package People

import (
	"database/sql"
	"encoding/json"
	"github.com/dfrnks/go-rest-api/internal/Database"
	"github.com/dfrnks/go-rest-api/internal/Models"
	"github.com/gorilla/mux"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var person Models.Person
	var address Models.Address

	row := Database.Connection().QueryRow("SELECT * FROM person WHERE id = $1", params["id"])

	err := row.Scan(&person.ID, &person.Firstname, &person.Lastname)
	switch err {
	case sql.ErrNoRows:
	case nil:
		row := Database.Connection().QueryRow("SELECT city, state FROM address WHERE idperson = $1", person.ID)
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
