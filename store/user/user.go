package user

import (
	"ThreeLayeredArchitecture/models/user"
	"database/sql"
	"log"
)

type UserStore struct {
	DB *sql.DB
}

func NewUserStore(db *sql.DB) *UserStore {
	return &UserStore{DB: db}
}

func (s *UserStore) CreateUser(user models.User) (models.User, error) {
	query := "INSERT INTO user (name) VALUES (?)"
	result, err := s.DB.Exec(query, user.Name)
	if err != nil {
		return models.User{}, err
	}
	id, _ := result.LastInsertId()
	user.ID = int(id)
	return user, nil
}

func (s *UserStore) GetAllUsers() ([]models.User, error) {
	rows, err := s.DB.Query("SELECT id, name FROM user")
	if err != nil {
		return nil, err
	}
	defer func() {
		if err := rows.Close(); err != nil {
			log.Printf("Error closing Task DB: %v", err)
		}
	}()

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

func (s *UserStore) GetUserByID(id int) (models.User, error) {
	row := s.DB.QueryRow("SELECT id, name FROM user WHERE id = ?", id)
	var user models.User
	err := row.Scan(&user.ID, &user.Name)
	if err != nil {
		return models.User{}, err
	}
	return user, nil
}
