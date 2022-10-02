package notification

type Service interface {
	Create(input NotificationCreateInput) (Notification, error)
	GetByID(input SomethingWithID) (Notification, error)
	Delete(input SomethingWithID) (Notification, error)
	GetByUserID(input SomethingWithID) ([]Notification, error)
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) Create(input NotificationCreateInput) (Notification, error) {
	notif := Notification{
		UserID:  input.UserID,
		Message: input.Message,
		Date:    input.Date,
	}

	notification, err := s.repository.Create(notif)
	if err != nil {
		return notification, err
	}

	return notification, nil
}

func (s *service) GetByID(input SomethingWithID) (Notification, error) {
	notification, err := s.repository.GetById(input.ID)
	if err != nil {
		return notification, err
	}

	return notification, nil
}

func (s *service) Delete(input SomethingWithID) (Notification, error) {
	notification, err := s.repository.Delete(input.ID)
	if err != nil {
		return notification, err
	}

	return notification, nil
}

func (s *service) GetByUserID(input SomethingWithID) ([]Notification, error) {
	notifications, err := s.repository.GetByUserID(input.ID)
	if err != nil {
		return notifications, err
	}

	return notifications, nil
}
