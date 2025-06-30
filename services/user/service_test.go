package user

import (
	models "ThreeLayeredArchitecture/models/user"
	"errors"
	"testing"
)

type MockUser struct{}

func (m *MockUser) GetUserByID(id int) (models.User, error) {
	if id == 1 {
		user1 := models.User{ID: 1, Name: "Ram"}
		return user1, nil
	} else {

		return models.User{}, errors.New("User not found")
	}
}

func (m *MockUser) GetAllUsers() ([]models.User, error) {
	alluser := []models.User{
		{
			ID:   1,
			Name: "ram",
		},
		{
			ID:   2,
			Name: "raj",
		},
	}

	return alluser, nil
}

func (m *MockUser) CreateUser(user models.User) (models.User, error) {
	return models.User{}, nil
}

func TestCreateUser(t *testing.T) {
	mock := &MockUser{}
	service := NewUserService(mock)

	newUser := models.User{Name: "Alice"}
	_, err := service.CreateUser(newUser)

	if err != nil {
		t.Error("Expected no error, got:", err)
	} else {
		t.Log("Successfully created user")
	}
}

func TestGetUserByID(t *testing.T) {
	mock := &MockUser{}
	service := NewUserService(mock)

	msg1, err := service.GetUserByID(1)
	if err != nil {
		t.Error(err)
	} else {
		t.Logf("Successfully get user by id: %d and name %s", msg1.ID, msg1.Name)
	}

	_, err = mock.GetUserByID(2)
	if err == nil {
		t.Error(err)
	} else {
		t.Logf("PASS")
	}
}

func TestGetUsers(t *testing.T) {
	mock := &MockUser{}
	service := NewUserService(mock)

	_, err := service.GetAllUsers()
	if err != nil {
		t.Error(err)
	} else {
		t.Log("Successfully created user")
	}

}
