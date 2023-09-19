package main

import (
	"log"

	"github.com/joho/godotenv"
	db "github.com/yogapratama23/go-http-server/internal/database"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Failed to load env .file in migrations")
	}

	db.Init()
	defer db.Connect.Close()

	createCategories()
	createProducts()
}

func createCategories() {
	query := `
		CREATE TABLE IF NOT EXISTS categories (
			id INT AUTO_INCREMENT,
			name VARCHAR(30),
			created_at DATETIME,
			updated_at DATETIME,
			deleted_at DATETIME,
			PRIMARY KEY (id)
		);
	`
	if _, err := db.Connect.Exec(query); err != nil {
		log.Println(err)
	}
}

func createProducts() {
	query := `
		CREATE TABLE IF NOT EXISTS products (
			id INT AUTO_INCREMENT,
			name VARCHAR(255),
			category_id INT,
			created_at DATETIME,
			updated_at DATETIME,
			deleted_at DATETIME,
			PRIMARY KEY (id),
			FOREIGN KEY (category_id) REFERENCES categories(id)
		);
	`
	if _, err := db.Connect.Exec(query); err != nil {
		log.Println(err)
	}
}
