package category

import (
	"strconv"

	"github.com/yogapratama23/go-http-server/internal/response"
)

type CategoryService struct {
	categoryRepo CategoryRepository
}

func (s *CategoryService) FindAll(p *response.PaginationInput) (*PaginateListCategory, error) {
	if (p.Skip == 0) && (p.Take == 0) {
		p.Skip = 0
		p.Take = 10
	}
	categories, err := s.categoryRepo.FindAllPaginate(p)
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

func (s *CategoryService) FindById(id string) (*ListCategory, error) {
	i, _ := strconv.Atoi(id)
	c, err := s.categoryRepo.FindById(&i)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *CategoryService) FindByName(n *string) (*ListCategory, error) {
	c, err := s.categoryRepo.FindByName(n)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *CategoryService) SoftDelete(id string) error {
	i, _ := strconv.Atoi(id)
	err := s.categoryRepo.SoftDelete(&i)
	if err != nil {
		return err
	}

	return nil
}
