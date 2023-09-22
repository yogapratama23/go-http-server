package product

type ProductService struct {
	productRepo ProductRepository
}

func (s *ProductService) Create(p *CreateProductInput) error {
	if err := s.productRepo.Create(p); err != nil {
		return err
	}

	return nil
}

func (s *ProductService) FindAll() (*ListFindAllResponse, error) {
	products, err := s.productRepo.FindAllWithDetails()
	if err != nil {
		return nil, err
	}

	return products, nil
}
