package category

import (
	"strconv"
)

type CategoryService struct {
	categoryRepo CategoryRepository
}

func (s *CategoryService) FindAll() (*[]ListCategory, error) {
	categories, err := s.categoryRepo.FindAll()
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
