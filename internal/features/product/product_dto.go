package product

type CreateProductInput struct {
	Name       string `json:"name"`
	CategoryId int    `json:"category_id"`
}
