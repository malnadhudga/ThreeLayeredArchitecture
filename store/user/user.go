package userstore

import (
	"ThreeLayeredArchitecture/models/user"
	"gofr.dev/pkg/gofr"
)

type UserStore struct {
}

func NewUserStore() *UserStore {
	return &UserStore{}
}

func (*UserStore) CreateUser(ctx *gofr.Context, user models.User) (models.User, error) {
	query := "INSERT INTO user (name) VALUES (?)"
	result, err := ctx.SQL.Exec(query, user.Name)

	if err != nil {
		return models.User{}, err
	}

	id, _ := result.LastInsertId()
	user.ID = int(id)

	return user, nil
}

func (*UserStore) GetAllUsers(ctx *gofr.Context) ([]models.User, error) {
	rows, err := ctx.SQL.Query("SELECT id, name FROM user")

	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return []models.User{}, rows.Err()
	}

	var users []models.User

	for rows.Next() {
		var user models.User
		err = rows.Scan(&user.ID, &user.Name)

		if err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

func (*UserStore) GetUserByID(ctx *gofr.Context, id int) (models.User, error) {
	row := ctx.SQL.QueryRow("SELECT id, name FROM user WHERE id = ?", id)

	var user models.User
	err := row.Scan(&user.ID, &user.Name)

	if err != nil {
		return models.User{}, err
	}

	return user, nil
}
