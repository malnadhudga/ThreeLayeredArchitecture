package user

import (
	userModel "ThreeLayeredArchitecture/models/user"
	"gofr.dev/pkg/gofr"
)

type Service struct {
	Store Store
}

func NewUserService(store Store) *Service {
	return &Service{Store: store}
}

func (s *Service) CreateUser(ctx *gofr.Context, user userModel.User) (userModel.User, error) {
	return s.Store.CreateUser(ctx, user)
}

func (s *Service) GetAllUsers(ctx *gofr.Context) ([]userModel.User, error) {
	return s.Store.GetAllUsers(ctx)
}

func (s *Service) GetUserByID(ctx *gofr.Context, id int) (user userModel.User, err error) {
	return s.Store.GetUserByID(ctx, id)
}
