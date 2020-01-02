package main

import (
	"database/sql"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	var dbfile string = "./test.db"

	os.Remove(dbfile)

	//	db, err := sql.Open("sqlite3", ":memory:")
	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`CREATE TABLE "world" ("id" INTEGER PRIMARY KEY AUTOINCREMENT, "country" VARCHAR(255), "capital" VARCHAR(255))`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(
		`INSERT INTO "world" ("country", "capital") VALUES (?, ?) `,
		"日本",
		"東京",
	)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare(`INSERT INTO "world" ("country", "capital") VALUES (?, ?) `)
	if err != nil {
		panic(err)
	}

	if _, err = stmt.Exec("アメリカ", "ワシントンD.C."); err != nil {
		panic(err)
	}
	if _, err = stmt.Exec("ロシア", "モスクワ"); err != nil {
		panic(err)
	}
	if _, err = stmt.Exec("イギリス", "ロンドン"); err != nil {
		panic(err)
	}
	if _, err = stmt.Exec("オーストラリア", "シドニー"); err != nil {
		panic(err)
	}
	stmt.Close()

	db.Close()
}
