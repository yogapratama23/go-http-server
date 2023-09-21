package product

import (
	"errors"
	"log"
	"time"

	db "github.com/yogapratama23/go-http-server/internal/database"
)

type ProductRepository struct{}

func (r *ProductRepository) Create(p *CreateProductInput) error {
	params := []interface{}{p.Name, p.CategoryId, time.Now(), time.Now()}
	query := `
		INSERT INTO
			products (
				name, category_id, created_at, updated_at
			)
		VALUES (
			?, ?, ?, ?
		)
	`

	if _, err := db.Connect.Exec(query, params...); err != nil {
		log.Println(err)
		return errors.New("create product failed")
	}

	return nil
}
