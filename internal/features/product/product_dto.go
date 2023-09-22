package product

import "github.com/yogapratama23/go-http-server/internal/features/category"

type CreateProductInput struct {
	Name       string `json:"name"`
	CategoryId int    `json:"category_id"`
}

type FindAllWithDetailsResponse struct {
	ID       int                       `json:"id"`
	Name     string                    `json:"name"`
	Category category.CategoryResponse `json:"category"`
}

type ProductResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ListFindAllResponse struct {
	Products []FindAllWithDetailsResponse `json:"products"`
}
