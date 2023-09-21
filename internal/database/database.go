package db

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/yogapratama23/go-http-server/internal/constants/message"
)

var Connect *sql.DB

func Init() {
	dsn := os.Getenv("DSN")
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err.Error())
	}

	log.Println(message.DatabaseConnected)
	Connect = db
}
