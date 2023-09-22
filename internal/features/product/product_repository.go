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

func (r *ProductRepository) FindAllWithDetails() (*ListFindAllResponse, error) {
	var response ListFindAllResponse
	query := `
		SELECT
			p.id, p.name, c.id, c.name
		FROM
			products AS p
		LEFT JOIN
			categories AS c ON p.category_id = c.id
		WHERE
			p.deleted_at IS NULL
	`

	rows, err := db.Connect.Query(query)
	if err != nil {
		log.Println(err)
		return nil, errors.New("error find all product with details")
	}
	defer rows.Close()

	for rows.Next() {
		var p FindAllWithDetailsResponse
		err := rows.Scan(&p.ID, &p.Name, &p.Category.ID, &p.Category.Name)
		if err != nil {
			log.Println(err)
			return nil, errors.New("error find all product with details")
		}
		response.Products = append(response.Products, p)
	}

	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, errors.New("error find all product with details")
	}

	return &response, nil
}
