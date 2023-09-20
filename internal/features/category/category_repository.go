package category

import (
	"errors"
	"log"
	"time"

	db "github.com/yogapratama23/go-http-server/internal/database"
	"github.com/yogapratama23/go-http-server/internal/models"
	"github.com/yogapratama23/go-http-server/internal/response"
)

type CategoryRepository struct{}

func (r *CategoryRepository) FindAllPaginate(p *response.PaginationInput, wc *FindAllWhereCond) (*ListCategoryResponse, error) {
	params := []interface{}{}
	paramsCount := []interface{}{}
	var response ListCategoryResponse
	query := `
		SELECT
			id, name
		FROM
			categories
		WHERE
			deleted_at IS NULL
	`
	queryCount := `
		SELECT
			COUNT(id)
		FROM
			categories
		WHERE
			deleted_at IS NULL
	`

	if wc.Id != 0 {
		query += " AND id = ?"
		queryCount += " AND id = ?"
		params = append(params, wc.Id)
		paramsCount = append(paramsCount, wc.Id)
	}

	if wc.Search != "" {
		query += ` AND name like CONCAT('%', ?, '%')`
		queryCount += ` AND name like CONCAT('%', ?, '%')`
		params = append(params, wc.Search)
		paramsCount = append(paramsCount, wc.Search)
	}

	if (p.Skip != 0) && (p.Take != 0) {
		query += " LIMIT ?, ?"
		params = append(params, p.Skip)
		params = append(params, p.Take)
	}
	// query for data
	rows, err := db.Connect.Query(query, params...)
	if err != nil {
		log.Println(err)
		return nil, errors.New("error finding list of categories")
	}
	defer rows.Close()

	for rows.Next() {
		var c CategoryResponse
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			log.Println(err)
			return nil, errors.New("error finding list of categories")
		}
		response.Categories = append(response.Categories, c)
	}

	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, errors.New("error finding list of categories")
	}

	// query for pagination
	err = db.Connect.QueryRow(queryCount, paramsCount...).Scan(&response.Total)
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	return &response, nil
}

func (r *CategoryRepository) Create(p *CreateCategoryInput) error {
	payload := models.Category{
		Name:      p.Name,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	_, err := db.Connect.Exec(`
		INSERT INTO 
			categories (
				name, created_at, updated_at
			)
		VALUES (
			?, ?, ?
		)
	`, payload.Name, payload.CreatedAt, payload.UpdatedAt)
	if err != nil {
		log.Println(err)
		return errors.New("create category failed")
	}

	return nil
}

func (r *CategoryRepository) SoftDelete(id *int) error {
	params := []interface{}{time.Now(), id}
	_, err := db.Connect.Exec(`
		UPDATE
			categories
		SET
			deleted_at = ?
		WHERE
			id = ?
	`, params...)
	if err != nil {
		log.Println(err)
		return errors.New("delete category failed")
	}

	return nil
}
