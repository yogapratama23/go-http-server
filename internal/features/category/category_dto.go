package category

type CreateCategoryInput struct {
	Name string `json:"name"`
}

type ListCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
