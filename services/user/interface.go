package user

import (
	userModel "ThreeLayeredArchitecture/models/user"
	"gofr.dev/pkg/gofr"
)

type Store interface {
	CreateUser(ctx *gofr.Context, user userModel.User) (userModel.User, error)
	GetAllUsers(ctx *gofr.Context) ([]userModel.User, error)
	GetUserByID(ctx *gofr.Context, id int) (userModel.User, error)
}
