package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	var id int
	var country string
	var capital string

	db, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}

	rows, err := db.Query(`SELECT * FROM world`)
	if err != nil {
		panic(err)
	}

	for rows.Next() {
		err = rows.Scan(&id, &country, &capital)
		if err != nil {
			panic(err)
		}
		fmt.Println(id, " ", country, " ", capital)
	}

	db.Close()
}
