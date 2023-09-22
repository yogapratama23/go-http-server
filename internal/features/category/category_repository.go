package category

import (
	"errors"
	"log"
	"time"

	"github.com/yogapratama23/go-http-server/internal/constants/message"
	db "github.com/yogapratama23/go-http-server/internal/database"
	"github.com/yogapratama23/go-http-server/internal/response"
)

type CategoryRepository struct{}

func (r *CategoryRepository) FindAll(p *response.PaginationInput, wc *FindAllWhereCond) (*ListCategoryResponse, error) {
	params := []interface{}{}
	countParams := []interface{}{}
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
		countParams = append(countParams, wc.Id)
	}

	if wc.Search != "" {
		query += ` AND name like CONCAT('%', ?, '%')`
		queryCount += ` AND name like CONCAT('%', ?, '%')`
		params = append(params, wc.Search)
		countParams = append(countParams, wc.Search)
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
		return nil, errors.New(message.ErrorFindingCategories)
	}
	defer rows.Close()

	for rows.Next() {
		var c CategoryResponse
		err := rows.Scan(&c.ID, &c.Name)
		if err != nil {
			log.Println(err)
			return nil, errors.New(message.ErrorFindingCategories)
		}
		response.Categories = append(response.Categories, c)
	}

	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, errors.New(message.ErrorFindingCategories)
	}

	// query for pagination
	err = db.Connect.QueryRow(queryCount, countParams...).Scan(&response.Total)
	if err != nil {
		log.Println(err)
		return nil, nil
	}

	return &response, nil
}

func (r *CategoryRepository) FindAllWithProducts() (*[]FindWithProductsResponse, error) {
	var response []FindWithProductsResponse
	query := `
		SELECT
			c.id, c.name, p.id, p.category_id, p.name
		FROM
			categories AS c
		LEFT JOIN
			products AS p ON c.id = p.category_id
		WHERE
			c.deleted_at IS NULL
	`

	rows, err := db.Connect.Query(query)
	if err != nil {
		log.Println(err)
		return nil, errors.New(message.ErrorFindingCategories)
	}
	defer rows.Close()

	for rows.Next() {
		var c *CategoryResponse = new(CategoryResponse)
		var p *ProductResponse = new(ProductResponse)
		var cp FindWithProductsResponse
		err := rows.Scan(&c.ID, &c.Name, &p.ID, &p.CategoryId, &p.Name)
		if err != nil {
			log.Println(err)
		}

		if len(response) > 0 {
			found := false
			for i := 0; i < len(response); i++ {
				if response[i].ID == c.ID {
					found = true
					response[i].Products = append(response[i].Products, *p)
					break
				}
				continue
			}

			if !found {
				cp.ID = c.ID
				cp.Name = c.Name
				if p.ID != 0 {
					cp.Products = append(cp.Products, *p)
				}
				response = append(response, cp)
			}
		} else {
			cp.ID = c.ID
			cp.Name = c.Name
			if p.ID != 0 {
				cp.Products = append(cp.Products, *p)
			}
			response = append(response, cp)
		}

	}

	err = rows.Err()
	if err != nil {
		log.Println(err)
		return nil, errors.New(message.ErrorFindingCategories)
	}

	return &response, nil
}

func (r *CategoryRepository) Create(p *CreateCategoryInput) error {
	params := []interface{}{p.Name, time.Now(), time.Now()}

	_, err := db.Connect.Exec(`
		INSERT INTO 
			categories (
				name, created_at, updated_at
			)
		VALUES (
			?, ?, ?
		)
	`, params...)
	if err != nil {
		log.Println(err)
		return errors.New(message.CreateCategoryFailed)
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
		return errors.New(message.DeleteCategoryFailed)
	}

	return nil
}

func (r *CategoryRepository) Update(id *int, p *UpdateCategoryInput) error {
	params := []interface{}{time.Now()}
	query := `
		UPDATE
			categories
		SET
			updated_at = ?
	`

	if p.Name != "" {
		query += " , name = ?"
		params = append(params, p.Name)
	}

	query += ` WHERE id = ?`
	params = append(params, id)

	_, err := db.Connect.Exec(query, params...)
	if err != nil {
		log.Println(err)
		return errors.New(message.UpdateCategoryFailed)
	}

	return nil
}
