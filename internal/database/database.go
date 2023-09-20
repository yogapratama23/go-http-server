package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type WhereCondition struct{}

var Connect *sql.DB

func Init() {
	dsn := os.Getenv("DSN")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println("Database connected!")
	Connect = db
}
