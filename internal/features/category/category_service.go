package category

import (
	"github.com/yogapratama23/go-http-server/internal/response"
)

type CategoryService struct {
	categoryRepo CategoryRepository
}

func (s *CategoryService) FindAll(p *response.PaginationInput, wc *FindAllWhereCond) (*ListCategoryResponse, error) {
	categories, err := s.categoryRepo.FindAll(p, wc)
	if err != nil {
		return nil, err
	}

	return categories, nil
}

func (s *CategoryService) Create(p *CreateCategoryInput) error {
	err := s.categoryRepo.Create(p)
	if err != nil {
		return err
	}

	return nil
}

func (s *CategoryService) SoftDelete(id int) error {
	err := s.categoryRepo.SoftDelete(&id)
	if err != nil {
		return err
	}

	return nil
}
