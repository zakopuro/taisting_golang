package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	var dbfile string = "./test.db"

	db, err := sql.Open("sqlite3", dbfile)
	if err != nil {
		panic(err)
	}

	stmt, err := db.Prepare(`UPDATE "world" SET capital=? WHERE country=?`)
	if err != nil {
		panic(err)
	}

	res, err := stmt.Exec("キャンベラ", "オーストラリア")
	if err != nil {
		panic(err)
	}

	affect, err := res.RowsAffected()
	if err != nil {
		panic(err)
	}

	fmt.Println(affect, "個のデータを更新しました。")
	stmt.Close()

	db.Close()
}
