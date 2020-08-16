package main

func syncDataBase() {
	createTables()
	insertRows()
}

func createTables() {
	person, err := db.Prepare("CREATE TABLE IF NOT EXISTS person (`id` INTEGER PRIMARY KEY AUTOINCREMENT,`firstname` VARCHAR(64) NULL,`lastname` VARCHAR(64) NULL);")
	if err != nil {
		panic(err)
	}

	if _, err := person.Exec(); err != nil {
		panic(err)
	}

	address, err := db.Prepare("CREATE TABLE IF NOT EXISTS address(id INTEGER PRIMARY KEY AUTOINCREMENT,idperson INTEGER NOT NULL REFERENCES person ON UPDATE CASCADE ON DELETE CASCADE,city VARCHAR(255) NULL,state VARCHAR(255) NULL);")
	if err != nil {
		panic(err)
	}

	if _, err := address.Exec(); err != nil {
		panic(err)
	}
}

func insertRows() {
	var people []Person

	people = append(people, Person{
		ID:        1,
		Firstname: "John",
		Lastname:  "Doe",
		Address: &Address{
			City:  "City X",
			State: "State X",
		},
	})
	people = append(people, Person{
		ID:        2,
		Firstname: "Koko",
		Lastname:  "Doe",
		Address: &Address{
			City:  "City Z",
			State: "State Y",
		},
	})
	people = append(people, Person{
		ID:        3,
		Firstname: "Francis",
		Lastname:  "Sunday",
		Address: &Address{
			City:  "City W",
			State: "State H",
		},
	})

	for _, item := range people {
		stmt, err := db.Prepare("INSERT INTO person(id, firstname, lastname) values(?,?,?)")
		if err != nil {
			panic(err)
		}

		res, _ := stmt.Exec(item.ID, item.Firstname, item.Lastname)

		if res != nil && item.Address != nil {
			stmt, err := db.Prepare("INSERT INTO address(idperson, city, state) values(?,?,?)")
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
