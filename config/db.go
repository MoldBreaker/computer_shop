package config

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

func GetConnection() *sql.DB {
	db, err := sql.Open("mysql", os.Getenv("CONNECTION_STRING"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected Successfully")
	return db
}
