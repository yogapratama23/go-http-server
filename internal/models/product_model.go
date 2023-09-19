package models

import "time"

type Product struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	CategoryId int       `json:"category_id"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	DeletedAt  time.Time `json:"deleted_at"`
}
