package category

import (
	"errors"
	"log"
	"math"
	"time"

	db "github.com/yogapratama23/go-http-server/internal/database"
	"github.com/yogapratama23/go-http-server/internal/models"
	"github.com/yogapratama23/go-http-server/internal/response"
)

type CategoryRepository struct{}

func (r *CategoryRepository) FindAllPaginate(p *response.PaginationInput) (*PaginateListCategory, error) {
	var response PaginateListCategory
	query := `
		SELECT
			id, name
		FROM
			categories
		WHERE
			deleted_at IS NULL
		LIMIT ?, ?
	`

	// query for data
	rows, err := db.Connect.Query(query, p.Skip, p.Take)
	if err != nil {
		log.Println(err)
		return nil, errors.New("error finding list of categories")
	}
	defer rows.Close()

	for rows.Next() {
		var c ListCategory
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
	err = db.Connect.QueryRow("SELECT COUNT(id) FROM categories").Scan(&response.Total)
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	response.Page = int(math.Ceil(float64(p.Skip)/float64(p.Take))) + 1
	response.PerPage = p.Take
	response.PageCount = int(math.Ceil(float64(response.Total) / float64(p.Take)))

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

func (r *CategoryRepository) FindById(id *int) (*ListCategory, error) {
	var c ListCategory
	query := `
		SELECT
			id, name
		FROM
			categories
		WHERE
			deleted_at IS NULL
		AND
			id = ?
	`
	err := db.Connect.QueryRow(query, id).Scan(&c.ID, &c.Name)
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	return &c, nil
}

func (r *CategoryRepository) FindByName(n *string) (*ListCategory, error) {
	var c ListCategory
	query := `
		SELECT
			id, name
		FROM
			categories
		WHERE
			deleted_at IS NULL
		AND
			name LIKE CONCAT ('%', ?, '%')
	`
	err := db.Connect.QueryRow(query, n).Scan(&c.ID, &c.Name)
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	return &c, nil
}

func (r *CategoryRepository) SoftDelete(id *int) error {
	_, err := db.Connect.Exec(`
		UPDATE
			categories
		SET
			deleted_at = ?
		WHERE
			id = ?
	`, time.Now(), id)
	if err != nil {
		log.Println(err)
		return errors.New("delete category failed")
	}

	return nil
}
