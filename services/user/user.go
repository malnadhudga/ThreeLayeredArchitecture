package user

import (
	userModel "ThreeLayeredArchitecture/models/user"
)

type Service struct {
	Store Store
}

func NewUserService(store Store) *Service {
	return &Service{Store: store}
}

func (s *Service) CreateUser(user userModel.User) (userModel.User, error) {
	return s.Store.CreateUser(user)
}

func (s *Service) GetAllUsers() ([]userModel.User, error) {
	return s.Store.GetAllUsers()
}

func (s *Service) GetUserByID(id int) (userModel.User, error) {
	return s.Store.GetUserByID(id)
}
