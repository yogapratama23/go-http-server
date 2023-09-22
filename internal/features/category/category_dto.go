package category

type CreateCategoryInput struct {
	Name string `json:"name"`
}

type CategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type ListCategoryResponse struct {
	Categories []CategoryResponse `json:"categories"`
	Total      int                `json:"total"`
}

type FindAllWhereCond struct {
	Id     int    `json:"id"`
	Search string `json:"search"`
}

type UpdateCategoryInput struct {
	Name string `json:"name"`
}

type ProductResponse struct {
	ID         int    `json:"id"`
	CategoryId int    `json:"category_id"`
	Name       string `json:"name"`
}

type FindWithProductsResponse struct {
	ID       int               `json:"id"`
	Name     string            `json:"name"`
	Products []ProductResponse `json:"products"`
}
