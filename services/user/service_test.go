package user

import (
	models "ThreeLayeredArchitecture/models/user"
	"errors"
	"go.uber.org/mock/gomock"
	"reflect"
	"testing"
)

func TestCreateUser(t *testing.T) {

	ctrl := gomock.NewController(t)
	mockStore := NewMockUserStore(ctrl)
	service := NewUserService(mockStore)

	validUser := models.User{ID: 1, Name: "John Doe"}
	invalidUser := models.User{}

	testCases := []struct {
		id       int
		desc     string
		input    models.User
		output   models.User
		expErr   error
		mockCall bool
	}{
		{1, "Valid User Creation", validUser, validUser, nil, true},
		{2, "Invalid User Creation", invalidUser, models.User{}, errors.New("invalid user"), true},
	}

	for _, tc := range testCases {
		if tc.mockCall {
			mockStore.EXPECT().CreateUser(tc.input).Return(tc.output, tc.expErr)
		}

		user, err := service.CreateUser(tc.input)
		if !errors.Is(err, tc.expErr) {
			t.Errorf("[%d] %s: expected (%v, %v), got (%v, %v)", tc.id, tc.desc, tc.output, tc.expErr, user, err)
		}
	}
}

//func TestAddUser(t *testing.T) {
//	test_cases := []struct {
//		id       int
//		desc     string
//		inp      string
//		exp      error
//		mockCall bool
//	}{
//		{1, "Testing for Valid Input", "Ram", nil, true},
//		{id: 2, desc: "Testing for Empty String", inp: "", exp: models.CustomError{Code: http.StatusBadRequest, Message: "Empty String given as input"}, mockCall: false},
//	}
//
//	ctrl := gomock.NewController(t)
//	mockStore := NewMockUserStore(ctrl)
//	svc := New(mockStore)
//
//	for _, test := range test_cases {
//		if test.mockCall {
//			mockStore.EXPECT().AddUser(test.inp).Return(test.exp)
//		}
//		err := svc.AddUser(test.inp)
//		if !errors.Is(err, test.exp) {
//			t.Errorf("Error in Testing: %v, Expected : %v, got : %v", test.exp, nil, err)
//		}
//	}

func TestGetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockUserStore(ctrl)
	service := NewUserService(mockStore)

	users := []models.User{
		{ID: 1, Name: "Alice"},
		{ID: 2, Name: "Bob"},
	}

	mockStore.EXPECT().GetAllUsers().Return(users, nil)

	result, err := service.GetAllUsers()
	if err != nil {
		t.Errorf("Expected no error, got %v", err)
	}
	if !reflect.DeepEqual(result, users) {
		t.Errorf("Expected %v, got %v", users, result)
	}
}

func TestGetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockStore := NewMockUserStore(ctrl)
	service := NewUserService(mockStore)

	user := models.User{ID: 1, Name: "Charlie"}

	testCases := []struct {
		id       int
		desc     string
		inputID  int
		expected models.User
		expErr   error
		mockCall bool
	}{
		{1, "Valid ID", 1, user, nil, true},
		{2, "User Not Found", 2, models.User{}, errors.New("user not found"), true},
	}

	for _, tc := range testCases {
		if tc.mockCall {
			mockStore.EXPECT().GetUserByID(tc.inputID).Return(tc.expected, tc.expErr)
		}

		result, err := service.GetUserByID(tc.inputID)
		if !errors.Is(err, tc.expErr) || !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("[%d] %s: expected (%v, %v), got (%v, %v)", tc.id, tc.desc, tc.expected, tc.expErr, result, err)
		}
	}
}

//
//
//mockgen -source interface.go -destination=mock_interface.go -package=task
//
//
//
//package user
//
//import (
//"awesomeProject/models"
//"errors"
//"go.uber.org/mock/gomock"
//"net/http"
//"reflect"
//"testing"
//)
//
//func TestAddUser(t *testing.T) {
//	test_cases := []struct {
//		id       int
//		desc     string
//		inp      string
//		exp      error
//		mockCall bool
//	}{
//		{1, "Testing for Valid Input", "Ram", nil, true},
//		{id: 2, desc: "Testing for Empty String", inp: "", exp: models.CustomError{Code: http.StatusBadRequest, Message: "Empty String given as input"}, mockCall: false},
//	}
//
//	ctrl := gomock.NewController(t)
//	mockStore := NewMockUserStore(ctrl)
//	svc := New(mockStore)
//
//	for _, test := range test_cases {
//		if test.mockCall {
//			mockStore.EXPECT().AddUser(test.inp).Return(test.exp)
//		}
//		err := svc.AddUser(test.inp)
//		if !errors.Is(err, test.exp) {
//			t.Errorf("Error in Testing: %v, Expected : %v, got : %v", test.exp, nil, err)
//		}
//	}
//
//}
//
//func TestViewTask(t *testing.T) {
//	test_cases := []struct {
//		id       int
//		desc     string
//		ifRow    bool
//		exp      models.UserSlice
//		expErr   error
//		mockCall bool
//	}{
//		{1, "Testing for Valid Input", true,
//			models.UserSlice{
//				{1, "Ram"},
//				{2, "Shyam"},
//			},
//			nil, true,
//		},
//		{
//			2, "Testing for no user", false,
//			models.UserSlice{},
//			models.CustomError{http.StatusNoContent, "No user Found"},
//			false,
//		},
//	}
//
//	ctrl := gomock.NewController(t)
//	mockStore := NewMockUserStore(ctrl)
//	svc := New(mockStore)
//
//	for _, test := range test_cases {
//		mockStore.EXPECT().CheckIfRowsExists().Return(test.ifRow)
//		if test.mockCall {
//			mockStore.EXPECT().ViewUser().Return(test.exp, test.expErr)
//		}
//		op, err := svc.ViewTask()
//		if !errors.Is(err, test.expErr) {
//			t.Errorf("Error in Testing: %v, Expected : %v, got : %v", test.desc, test.expErr, err)
//		}
//
//		if !reflect.DeepEqual(op, test.exp) {
//			t.Errorf("Error in Testing: %v, Expected : %v, got : %v", test.exp, test.exp, op)
//		}
//	}
//}
//
//func TestGetUserId(t *testing.T) {
//	test_cases := []struct {
//		id       int
//		desc     string
//		ifUser   bool
//		input    int
//		exp      models.User
//		expErr   error
//		mockCall bool
//	}{
//		{1, "Testing while user exists", true, 1, models.User{1, "Ram"}, nil, true},
//		{2, "Testing while user doesn't exists", false, 5, models.User{}, models.CustomError{Code: http.StatusNotFound, Message: "user does not exists"}, false},
//	}
//
//	ctrl := gomock.NewController(t)
//	mockStore := NewMockUserStore(ctrl)
//	svc := New(mockStore)
//
//	for _, test := range test_cases {
//		mockStore.EXPECT().CheckUserID(test.input).Return(test.ifUser)
//		if test.mockCall {
//			mockStore.EXPECT().GetUserByID(test.input).Return(test.exp, test.expErr)
//		}
//		op, err := svc.GetUserId(test.input)
//		if !errors.Is(err, test.expErr) {
//			t.Errorf("Error in Testing: %v, Expected : %v, got : %v", test.desc, test.expErr, err)
//		}
//		if !reflect.DeepEqual(op, test.exp) {
//			t.Errorf("Error in Testing: %v, Expected : %v, got : %v", test.desc, test.exp, op)
//		}
//	}
//}
//
//
