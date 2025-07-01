package user

import (
	models "ThreeLayeredArchitecture/models/user"
	"bytes"
	"errors"
	"go.uber.org/mock/gomock"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUserHandler_CreateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := NewMockUserService(ctrl)
	handler := NewUserHandler(mockService)

	testCases := []struct {
		id           int
		desc         string
		inputUser    models.User
		inputJSON    string
		mockReturn   models.User
		mockError    error
		expectStatus int
	}{
		{
			id:           1,
			desc:         "Successful user creation",
			inputUser:    models.User{ID: 1, Name: "ram"},
			inputJSON:    `{"ID":1,"Name":"ram"}`,
			mockReturn:   models.User{ID: 1, Name: "ram"},
			mockError:    nil,
			expectStatus: http.StatusOK,
		},
		{
			id:           2,
			desc:         "Invalid JSON body",
			inputJSON:    `{"ID":1, "Name":}`,
			expectStatus: http.StatusBadRequest,
		},
		{
			id:           3,
			desc:         "Service returns error",
			inputUser:    models.User{ID: 2, Name: "ErrorUser"},
			inputJSON:    `{"ID":2,"Name":"ErrorUser"}`,
			mockReturn:   models.User{},
			mockError:    errors.New("create error"),
			expectStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		var body io.Reader
		if tc.inputJSON != "" {
			body = bytes.NewBufferString(tc.inputJSON)
		}

		req := httptest.NewRequest(http.MethodPost, "/users", body)
		w := httptest.NewRecorder()

		if tc.mockError != nil || tc.mockReturn != (models.User{}) {
			mockService.EXPECT().CreateUser(tc.inputUser).Return(tc.mockReturn, tc.mockError)
		}

		handler.CreateUser(w, req)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != tc.expectStatus {
			t.Errorf("[Test ID %d] %s: expected status %d, got %d", tc.id, tc.desc, tc.expectStatus, res.StatusCode)
		}
	}
}

func TestUserHandler_GetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := NewMockUserService(ctrl)
	handler := NewUserHandler(mockService)

	testCases := []struct {
		id           int
		desc         string
		mockReturn   []models.User
		mockError    error
		expectStatus int
	}{
		{
			id:           1,
			desc:         "Successful user fetch",
			mockReturn:   []models.User{{ID: 1, Name: "ram"}, {ID: 2, Name: "bhim"}},
			mockError:    nil,
			expectStatus: http.StatusOK,
		},
		{
			id:           2,
			desc:         "Failure in service layer",
			mockReturn:   nil,
			mockError:    errors.New("db error"),
			expectStatus: http.StatusInternalServerError,
		},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/users", nil)
		w := httptest.NewRecorder()

		mockService.EXPECT().GetAllUsers().Return(tc.mockReturn, tc.mockError)

		handler.GetAllUsers(w, req)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != tc.expectStatus {
			t.Errorf("[Test ID %d] %s: expected status %d, got %d", tc.id, tc.desc, tc.expectStatus, res.StatusCode)
		}
	}
}

func TestUserHandler_GetUserByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockService := NewMockUserService(ctrl)
	handler := NewUserHandler(mockService)

	testCases := []struct {
		id           int
		desc         string
		inputID      string
		mockID       int
		mockReturn   models.User
		mockError    error
		expectStatus int
	}{
		{
			id:           1,
			desc:         "Valid ID, user found",
			inputID:      "1",
			mockID:       1,
			mockReturn:   models.User{ID: 1, Name: "ram"},
			mockError:    nil,
			expectStatus: http.StatusOK,
		},
		{
			id:           2,
			desc:         "Invalid ID format",
			inputID:      "abc",
			expectStatus: http.StatusBadRequest,
		},
		{
			id:           3,
			desc:         "User not found",
			inputID:      "99",
			mockID:       99,
			mockReturn:   models.User{},
			mockError:    errors.New("not found"),
			expectStatus: http.StatusNotFound,
		},
	}

	for _, tc := range testCases {
		req := httptest.NewRequest(http.MethodGet, "/users/"+tc.inputID, nil)
		req.SetPathValue("id", tc.inputID)

		w := httptest.NewRecorder()

		if tc.mockError != nil || tc.mockReturn != (models.User{}) {
			mockService.EXPECT().GetUserByID(tc.mockID).Return(tc.mockReturn, tc.mockError)
		}

		handler.GetUserbyID(w, req)

		res := w.Result()
		defer res.Body.Close()

		if res.StatusCode != tc.expectStatus {
			t.Errorf("[Test ID %d] %s: expected status %d, got %d", tc.id, tc.desc, tc.expectStatus, res.StatusCode)
		}
	}
}
