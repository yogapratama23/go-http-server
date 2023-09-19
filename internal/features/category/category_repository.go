package category

import (
	"time"

	db "github.com/yogapratama23/go-http-server/internal/database"
	"github.com/yogapratama23/go-http-server/internal/models"
)

type CategoryRepository struct{}

func (r *CategoryRepository) FindAll() (*[]ListCategory, error) {
	var categories []ListCategory
	query := `
		SELECT
			id, name
		FROM
			categories
		WHERE
			deleted_at IS NULL;
	`

	rows, err := db.Connect.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var c ListCategory
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			return nil, err
		}
		categories = append(categories, c)
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &categories, nil
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
		);
	`, payload.Name, payload.CreatedAt, payload.UpdatedAt)
	if err != nil {
		return err
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
			id = ?;
	`
	err := db.Connect.QueryRow(query, id).Scan(&c.ID, &c.Name)
	if err != nil {
		return nil, err
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
			name LIKE CONCAT ('%', ?, '%');
	`
	err := db.Connect.QueryRow(query, n).Scan(&c.ID, &c.Name)
	if err != nil {
		return nil, err
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
			id = ?;
	`, time.Now(), id)
	if err != nil {
		return err
	}

	return nil
}
