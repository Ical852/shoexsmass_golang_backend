package category

type Service interface {
	GetALl() ([]Category, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetALl() ([]Category, error) {
	categories, err := s.repository.GetAll()
	if err != nil {
		return categories, err
	}

	return categories, nil
}
