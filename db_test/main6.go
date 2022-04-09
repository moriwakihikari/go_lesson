package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

var DbConnection *sql.DB

type Person struct {
	Name string
	Age  int
}

func main() {
	DbConnection, _ := sql.Open("sqlite3", "./example.sql")
	defer DbConnection.Close()

	cmd := "DELETE FROM persons WHERE name = ?"
	_, err := DbConnection.Exec(cmd, "hanako")
	if err != nil {
		log.Fatalln(err)
	}
}
