package user

import models "ThreeLayeredArchitecture/models/user"

type UserService interface {
	CreateUser(models.User) (models.User, error)
	GetAllUsers() ([]models.User, error)
	GetUserByID(int) (models.User, error)
}
