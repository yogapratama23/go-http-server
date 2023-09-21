package category

import (
	"strconv"

	"github.com/yogapratama23/go-http-server/internal/response"
)

type CategoryService struct {
	categoryRepo CategoryRepository
}

func (s *CategoryService) FindAll(p *response.PaginationInput, wc *FindAllWhereCond) (*ListCategoryResponse, error) {
	categories, err := s.categoryRepo.FindAllPaginate(p, wc)
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

func (s *CategoryService) SoftDelete(id string) error {
	i, _ := strconv.Atoi(id)
	err := s.categoryRepo.SoftDelete(&i)
	if err != nil {
		return err
	}

	return nil
}
