package db

import (
	"database/sql"

	_ "github.com/lib/pq" // _ means excute just during it is running
)

func ConnectDB() *sql.DB {
	conn := "user=postgres dbname=alura_store password=123qweasd host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conn)
	if err != nil {
		panic(err.Error())
	}
	return db
}
