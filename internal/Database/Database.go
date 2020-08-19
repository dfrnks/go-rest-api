package Database

import (
	"database/sql"
	"github.com/dfrnks/go-rest-api/internal/Models"
	"github.com/dfrnks/go-rest-api/internal/Utils"
	_ "github.com/mattn/go-sqlite3"
)

var db, err = sql.Open("sqlite3", "./database.sqlite")

func Connection() *sql.DB {
	if err != nil {
		panic(err)
	}

	return db
}

func SyncDataBase() {
	createTables()
	insertRows()
}

func createTables() {
	person, err := Connection().Prepare("CREATE TABLE IF NOT EXISTS person (`id` INTEGER PRIMARY KEY AUTOINCREMENT,`firstname` VARCHAR(64) NULL,`lastname` VARCHAR(64) NULL);")
	if err != nil {
		panic(err)
	}

	if _, err := person.Exec(); err != nil {
		panic(err)
	}

	address, err := Connection().Prepare("CREATE TABLE IF NOT EXISTS address(id INTEGER PRIMARY KEY AUTOINCREMENT,idperson INTEGER NOT NULL REFERENCES person ON UPDATE CASCADE ON DELETE CASCADE,city VARCHAR(255) NULL,state VARCHAR(255) NULL);")
	if err != nil {
		panic(err)
	}

	if _, err := address.Exec(); err != nil {
		panic(err)
	}
}

func insertRows() {
	var people []Models.Person

	people = append(people, Models.Person{
		ID:        1,
		Firstname: "John",
		Lastname:  "Doe",
		Address: &Models.Address{
			City:  Utils.NullString{sql.NullString{String: "City X", Valid: true}},
			State: Utils.NullString{sql.NullString{String: "State X", Valid: true}},
		},
	})
	people = append(people, Models.Person{
		ID:        2,
		Firstname: "Koko",
		Lastname:  "Doe",
		Address: &Models.Address{
			City:  Utils.NullString{sql.NullString{String: "City Z", Valid: true}},
			State: Utils.NullString{sql.NullString{String: "State Y", Valid: true}},
		},
	})
	people = append(people, Models.Person{
		ID:        3,
		Firstname: "Francis",
		Lastname:  "Sunday",
	})

	for _, item := range people {
		stmt, err := Connection().Prepare("INSERT INTO person(id, firstname, lastname) values(?,?,?)")
		if err != nil {
			panic(err)
		}

		res, _ := stmt.Exec(item.ID, item.Firstname, item.Lastname)

		if res != nil && item.Address != nil {
			stmt, err := Connection().Prepare("INSERT INTO address(idperson, city, state) values(?,?,?)")
			if err != nil {
				panic(err)
			}

			_, err = stmt.Exec(item.ID, item.Address.City, item.Address.State)
			if err != nil {
				panic(err)
			}
		}
	}
}
