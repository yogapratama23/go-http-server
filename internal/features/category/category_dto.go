package category

type CreateCategoryInput struct {
	Name string `json:"name"`
}

type ListCategory struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type PaginateListCategory struct {
	Categories []ListCategory `json:"categories"`
	Total      int            `json:"total"`
	Page       int            `json:"page"`
	PerPage    int            `json:"per_page"`
	PageCount  int            `json:"page_count"`
}
