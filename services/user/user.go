package user

import (
	userModel "ThreeLayeredArchitecture/models/user"
)

type UserService struct {
	Store UserStore
}

func NewUserService(store UserStore) *UserService {
	return &UserService{Store: store}
}

func (s *UserService) CreateUser(user userModel.User) (userModel.User, error) {
	return s.Store.CreateUser(user)
}

func (s *UserService) GetAllUsers() ([]userModel.User, error) {
	return s.Store.GetAllUsers()
}

func (s *UserService) GetUserByID(id int) (userModel.User, error) {
	return s.Store.GetUserByID(id)
}
