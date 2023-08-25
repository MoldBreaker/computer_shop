package config

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func GetConnection() *sql.DB {
	LoadENV()
	db, err := sql.Open("mysql", os.Getenv("CONNECTION_STRING"))
	if err != nil {
		log.Fatal(err)
	}
	return db
}
