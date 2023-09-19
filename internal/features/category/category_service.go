package category

import (
	"strconv"

	"github.com/yogapratama23/go-http-server/internal/models"
)

type CategoryService struct {
	categoryRepo CategoryRepository
}

func (s *CategoryService) FindAll() (*[]models.Category, error) {
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

func (s *CategoryService) FindById(id string) (*models.Category, error) {
	i, _ := strconv.Atoi(id)
	c, err := s.categoryRepo.FindById(&i)
	if err != nil {
		return nil, err
	}

	return c, nil
}

func (s *CategoryService) FindByName(n *string) (*models.Category, error) {
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
