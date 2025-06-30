package user

import userModel "ThreeLayeredArchitecture/models/user"

type UserStore interface {
	CreateUser(userModel.User) (userModel.User, error)
	GetAllUsers() ([]userModel.User, error)
	GetUserByID(int) (userModel.User, error)
}
