package userstore_test

import (
	models "ThreeLayeredArchitecture/models/user"
	userstore "ThreeLayeredArchitecture/store/user"
	"errors"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestCreateUser(t *testing.T) {
	testcases := []struct {
		name         string
		inputUser    models.User
		expectedUser models.User
		expectErr    error
		mockError    error
	}{
		{
			name:         "valid user Bhim",
			inputUser:    models.User{Name: "bhim"},
			expectedUser: models.User{ID: 1, Name: "bhim"},
			expectErr:    nil,
			mockError:    nil,
		},
		{
			name:         "valid user with empty name",
			inputUser:    models.User{Name: ""},
			expectedUser: models.User{ID: 2, Name: ""},
			expectErr:    nil,
			mockError:    nil,
		},
		{
			name:         "error while inserting user",
			inputUser:    models.User{Name: "error_case"},
			expectedUser: models.User{},
			expectErr:    errors.New("insertion error"),
			mockError:    errors.New("insertion error"),
		},
	}

	for _, tc := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		defer db.Close()

		store := userstore.NewUserStore(db)

		exec := mock.ExpectExec("INSERT INTO user (name) VALUES (?)").
			WithArgs(tc.inputUser.Name)

		if tc.mockError != nil {
			exec.WillReturnError(tc.mockError)
		} else {
			exec.WillReturnResult(sqlmock.NewResult(int64(tc.expectedUser.ID), 1))
		}

		u, err := store.CreateUser(tc.inputUser)

		if (err != nil) != (tc.expectErr != nil) {
			t.Errorf("expected error: %v, got: %v", tc.expectErr, err)
		}

		if u.ID != tc.expectedUser.ID || u.Name != tc.expectedUser.Name {
			t.Errorf("unexpected user: got %+v, expected %+v", u, tc.expectedUser)
		} else {
			t.Log("PASS")
		}
	}
}

func TestGetAllUsers(t *testing.T) {
	testcases := []struct {
		expectedUser models.User
		expectErr    error
	}{
		{
			expectedUser: models.User{ID: 1, Name: "alan"},
			expectErr:    nil,
		},
		{
			expectedUser: models.User{ID: 2, Name: "ram"},
			expectErr:    nil,
		},
	}

	for _, tc := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		defer db.Close()

		store := userstore.NewUserStore(db)

		rows := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(tc.expectedUser.ID, tc.expectedUser.Name)

		mock.ExpectQuery("SELECT id, name FROM user").
			WillReturnRows(rows)

		users, err := store.GetAllUsers()

		if err != tc.expectErr {
			t.Errorf("expected error: %v, got: %v", tc.expectErr, err)
			continue
		}

		if len(users) != 1 {
			t.Fatalf("expected 1 user, got %d", len(users))
		}

		user := users[0]
		if user.ID != tc.expectedUser.ID || user.Name != tc.expectedUser.Name {
			t.Errorf("unexpected user: got %+v, expected %+v", user, tc.expectedUser)
		} else {
			t.Log("PASS")
		}
	}
}

func TestGetUserByID(t *testing.T) {
	testcases := []struct {
		name         string
		inputID      int
		expectedUser models.User
		expectErr    error
	}{
		{
			name:         "valid user Bhim",
			inputID:      1,
			expectedUser: models.User{ID: 1, Name: "Bhim"},
			expectErr:    nil,
		},
		{
			name:         "valid user Ram",
			inputID:      2,
			expectedUser: models.User{ID: 2, Name: "Ram"},
			expectErr:    nil,
		},
	}

	for _, tc := range testcases {
		t.Run(tc.name, func(t *testing.T) {
			db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
			if err != nil {
				t.Fatalf("unexpected error: %s", err)
			}
			defer db.Close()

			store := userstore.NewUserStore(db)

			rows := sqlmock.NewRows([]string{"id", "name"}).
				AddRow(tc.expectedUser.ID, tc.expectedUser.Name)

			mock.ExpectQuery("SELECT id, name FROM user WHERE id = ?").
				WithArgs(tc.inputID).
				WillReturnRows(rows)

			u, err := store.GetUserByID(tc.inputID)

			if (err != nil) != (tc.expectErr != nil) {
				t.Errorf("expected error: %v, got: %v", tc.expectErr, err)
				return
			}

			if u.ID != tc.expectedUser.ID || u.Name != tc.expectedUser.Name {
				t.Errorf("unexpected user: got %+v, expected %+v", u, tc.expectedUser)
			} else {
				t.Log("PASS")
			}
		})
	}
}
