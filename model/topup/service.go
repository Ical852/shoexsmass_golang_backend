package topup

import "shoexsmass/model/user"

type Service interface {
	CreateTopUp(input CreateTopUp) (TopUp, error)
	GetByUserID(input SomethingWithID) ([]TopUp, error)
	GetByID(input SomethingWithID) (TopUp, error)
}

type service struct {
	repository Repository
	userRepository user.Repository
}

func NewService(repository Repository, userRepository user.Repository) *service {
	return &service{repository, userRepository}
}

func (s *service) CreateTopUp(input CreateTopUp) (TopUp, error) {
	topUpCreate := TopUp{
		UserID:      input.UserID,
		OrderID: string(input.OrderID),
		GrossAmount: input.GrossAmount,
	}

	user, err := s.userRepository.FindById(input.UserID)
	if err != nil {
		return topUpCreate, err
	}

	topUpCreate.User = user

	topUp, err := s.repository.TopUP(topUpCreate)
	if err != nil {
		return topUp, err
	}

	user.Balance = user.Balance + input.GrossAmount

	_, err = s.userRepository.Update(user)
	if err != nil {
		return topUp, err
	}

	return topUp, nil
}

func (s *service) GetByUserID(input SomethingWithID) ([]TopUp, error) {
	topUps, err := s.repository.GetByUserID(input.ID)
	if err != nil {
		return topUps, err
	}

	return topUps, nil
}

func (s *service) GetByID(input SomethingWithID) (TopUp, error) {
	topUp, err := s.repository.GetByID(input.ID)
	if err != nil {
		return topUp, err
	}

	return topUp, nil
}
