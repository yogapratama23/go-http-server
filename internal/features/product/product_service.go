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
