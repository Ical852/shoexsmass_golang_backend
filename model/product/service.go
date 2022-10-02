package product

type Service interface {
	GetAll() ([]Product, error)
	GetByCategoryID(input SomethingWithID) ([]Product, error)
	GetByID(input SomethingWithID) (Product, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetAll() ([]Product, error) {
	products, err := s.repository.GetAll()
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *service) GetByCategoryID(input SomethingWithID) ([]Product, error) {
	products, err := s.repository.GetByCategoryID(input.ID)
	if err != nil {
		return products, err
	}

	return products, nil
}

func (s *service) GetByID(input SomethingWithID) (Product, error) {
	product, err := s.repository.GetByID(input.ID)
	if err != nil {
		return product, err
	}

	return product, nil
}
