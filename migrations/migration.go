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
	createUsers()
	createTokens()
	indexCategoryIdProducts()
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

func createUsers() {
	query := `
		CREATE TABLE IF NOT EXISTS users (
			id INT AUTO_INCREMENT,
			username VARCHAR(255),
			password VARCHAR(255),
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

func createTokens() {
	query := `
		CREATE TABLE IF NOT EXISTS tokens (
			id INT AUTO_INCREMENT,
			token VARCHAR(255),
			user_id INT,
			created_at DATETIME,
			updated_at DATETIME,
			deleted_at DATETIME,
			PRIMARY KEY (id),
			FOREIGN KEY (user_id) REFERENCES users(id)
		);
	`

	if _, err := db.Connect.Exec(query); err != nil {
		log.Println(err)
	}
}

func indexCategoryIdProducts() {
	query := `
		CREATE INDEX
			idx_category_id
		ON
			products (category_id)
	`

	if _, err := db.Connect.Exec(query); err != nil {
		log.Println(err)
	}
}
