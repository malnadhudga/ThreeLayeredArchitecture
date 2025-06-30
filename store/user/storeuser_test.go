package userstore_test

import (
	models "ThreeLayeredArchitecture/models/user"
	userstore "ThreeLayeredArchitecture/store/user"

	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestCreateUser(t *testing.T) {
	testcases := []struct {
		name       string
		expectedID int64
	}{
		{
			name:       "Alice",
			expectedID: 1,
		},
		{
			name:       "Bob",
			expectedID: 2,
		},
	}

	for _, tc := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		store := userstore.NewUserStore(db)

		mock.ExpectExec("INSERT INTO user (name) VALUES (?)").
			WithArgs(tc.name).
			WillReturnResult(sqlmock.NewResult(tc.expectedID, 1))

		u, err := store.CreateUser(models.User{Name: tc.name})

		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if int64(u.ID) != tc.expectedID || u.Name != tc.name {
			t.Errorf("unexpected user: %+v", u)
		} else {
			t.Log("PASS")
		}
	}
}

func TestGetAllUsers(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	store := userstore.NewUserStore(db)

	rows := sqlmock.NewRows([]string{"id", "name"}).
		AddRow(1, "Alice").
		AddRow(2, "Bob")

	mock.ExpectQuery("SELECT id, name FROM user").
		WillReturnRows(rows)

	users, err := store.GetAllUsers()
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}
	if len(users) != 2 {
		t.Errorf("expected 2 users, got %d", len(users))
	}
	if users[0].Name != "Alice" || users[1].Name != "Bob" {
		t.Errorf("unexpected users: %+v", users)
	}
}

func TestGetUserByID(t *testing.T) {
	testcases := []struct {
		id   int
		name string
	}{
		{
			id:   1,
			name: "Bhim",
		},
		{id: 2,
			name: "Ram",
		},
	}

	for _, tc := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
		store := userstore.NewUserStore(db)

		rows := sqlmock.NewRows([]string{"id", "name"}).
			AddRow(tc.id, tc.name)

		mock.ExpectQuery("SELECT id, name FROM user WHERE id = ?").
			WithArgs(tc.id).
			WillReturnRows(rows)

		u, err := store.GetUserByID(tc.id)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if u.ID != tc.id || u.Name != tc.name {
			t.Errorf("unexpected user: %+v", u)
		}
	}
}
