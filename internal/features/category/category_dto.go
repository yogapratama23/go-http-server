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
	Id     int
	Search string
}
