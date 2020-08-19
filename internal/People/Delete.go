package People

import (
	"encoding/json"
	"github.com/dfrnks/go-rest-api/internal/Database"
	"github.com/gorilla/mux"
	"net/http"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	stmt, err := Database.Connection().Prepare("delete from person where id=?")
	if err != nil {
		panic(err)
	}

	_, err = stmt.Exec(params["id"])
	if err != nil {
		panic(err)
	}

	stmt, err = Database.Connection().Prepare("delete from address where idperson=?")
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
