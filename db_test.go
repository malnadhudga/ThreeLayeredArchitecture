package main

//
//import (
//	models "ThreeLayeredArchitecture/models/user"
//	"testing"
//)
//
////
////import (
////	models "ThreeLayeredArchitecture/models/user"
////	"testing"
////)
////
////////
////////import (
////////	models "ThreeLayeredArchitecture/models/user"
////////	"errors"
////////)
////////
////////type MockStote struct{}
////////
////////// user, err := h.Service.GetUserByID(id)
////////func (m MockStote) GetUserByID(id int) (models.User, error) {
////////	if id == 10 {
////////		return models.User, nil
////////	} else {
////////		return nil, errors.New("User not found")
////////	}
////////}
//
//package main
//
//import (
//"ThreeLayeredArchitecture/models/user"
//"testing"
//)
//
//type TaskStore interface {
//	GetUserByID(id int) (models.User, error)
//}
//
//type MockStore struct{}
//
////GetUserByID func(id int) (model.User, error)
//
//func (m *MockStore) GetUserByID(id int) (models.User, error) {
//	if id == 1 {
//		return
//	}
//}
//
//func TestUserGetByID(t *testing.T) {
//	mockStore := &MockStore{}
//
//	//service := store.NewUserService(mockStore)
//	//user, err := service.GetUserByID(1)
//	//if err != nil {
//	//t.Fatalf("expected no error, got %v", err)
//	//}
//	//if user.ID != 1 || user.Name != "Shreya" {
//	//t.Errorf("unexpected user: %+v", user)
//	//}
