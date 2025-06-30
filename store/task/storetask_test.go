package taskstore_test

import (
	model "ThreeLayeredArchitecture/models/task"
	"ThreeLayeredArchitecture/store/task"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestAdd(t *testing.T) {
	testcases := []struct {
		desc          string
		completed     bool
		expectedTask  model.Task
		expectedError error
	}{
		{
			desc:      "Clean",
			completed: false,
			expectedTask: model.Task{
				ID:          1,
				Description: "Clean",
				Completed:   false,
			},
			expectedError: nil,
		},
		{
			desc:      "Play",
			completed: false,
			expectedTask: model.Task{
				ID:          2,
				Description: "Play",
				Completed:   false,
			},
			expectedError: nil,
		},
	}

	for _, testcases := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}

		store := taskstore.NewTaskStore(db)

		mock.ExpectExec("INSERT INTO task (description, completed) VALUES (?, ?)").
			WithArgs(testcases.desc, testcases.completed).
			WillReturnResult(sqlmock.NewResult(int64(testcases.expectedTask.ID), 1))

		task, err := store.Add(testcases.desc) // assume Add takes both desc and completed
		if err != testcases.expectedError {
			t.Errorf("expected error %v, got %v", testcases.expectedError, err)
		}

		if task.ID != testcases.expectedTask.ID ||
			task.Description != testcases.expectedTask.Description ||
			task.Completed != testcases.expectedTask.Completed {
			t.Errorf("unexpected task: got %+v, expected %+v", task, testcases.expectedTask)
		} else {
			t.Log("PASS")
		}
	}
}

func TestGetByID(t *testing.T) {
	testcases := []struct {
		desc          string
		completed     bool
		id            int
		expectedTask  model.Task
		expectedError error
	}{
		{
			desc:      "Clean",
			completed: false,
			id:        1,
			expectedTask: model.Task{
				ID:          1,
				Description: "Clean",
				Completed:   false,
			},
			expectedError: nil,
		},
		{
			desc:      "Play",
			completed: false,
			id:        2,
			expectedTask: model.Task{
				ID:          2,
				Description: "Play",
				Completed:   false,
			},
			expectedError: nil,
		},
		{
			desc:      "Play",
			completed: false,
			id:        3,
			expectedTask: model.Task{
				ID:          3,
				Description: "Play",
				Completed:   false,
			},
			expectedError: nil,
		},
	}

	for _, testcases := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}

		store := taskstore.NewTaskStore(db)

		rows := sqlmock.NewRows([]string{"id", "description", "completed"}).
			AddRow(testcases.expectedTask.ID, testcases.desc, testcases.completed)

		mock.ExpectQuery("SELECT id, description, completed FROM task WHERE id = ?").
			WithArgs(testcases.id).
			WillReturnRows(rows)

		task, err := store.GetByID(testcases.id)
		if err != testcases.expectedError {
			t.Errorf("expected error %v, got %v", testcases.expectedError, err)
		}

		if task.ID != testcases.expectedTask.ID ||
			task.Description != testcases.expectedTask.Description ||
			task.Completed != testcases.expectedTask.Completed {
			t.Errorf("unexpected task: got %+v, expected %+v", task, testcases.expectedTask)
		} else {
			t.Log("PASS")
		}
	}
}

func TestGetPending(t *testing.T) {
	testcases := []struct {
		desc          string
		completed     bool
		id            int
		expectedTask  model.Task
		expectedError error
	}{
		{
			desc:      "Clean",
			completed: false,
			id:        1,
			expectedTask: model.Task{
				ID:          1,
				Description: "Clean",
				Completed:   false,
			},
			expectedError: nil,
		},
		{
			desc:      "Play",
			completed: false,
			id:        2,
			expectedTask: model.Task{
				ID:          2,
				Description: "Play",
				Completed:   false,
			},
			expectedError: nil,
		},
		{
			desc:      "Play",
			completed: false,
			id:        3,
			expectedTask: model.Task{
				ID:          3,
				Description: "Play",
				Completed:   false,
			},
			expectedError: nil,
		},
	}

	for _, testcases := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}

		store := taskstore.NewTaskStore(db)

		rows := sqlmock.NewRows([]string{"id", "description", "completed"}).
			AddRow(testcases.expectedTask.ID, testcases.desc, testcases.completed)

		mock.ExpectQuery("SELECT id, description, completed FROM task WHERE completed = FALSE ORDER BY id").
			WillReturnRows(rows)

		tasks, err := store.GetPending()
		if err != testcases.expectedError {
			t.Errorf("expected error %v, got %v", testcases.expectedError, err)
		}
		if len(tasks) != 1 {
			t.Errorf("expected 1 task, got %d", len(tasks))
			continue
		}
		task := tasks[0]
		if task.ID != testcases.expectedTask.ID ||
			task.Description != testcases.expectedTask.Description ||
			task.Completed != testcases.expectedTask.Completed {
			t.Errorf("unexpected task: got %+v, expected %+v", task, testcases.expectedTask)
		} else {
			t.Log("PASS")
		}
	}
}

func TestMarkComplete(t *testing.T) {

	testcases := []struct {
		id            int64
		expectedError any
	}{
		{
			id:            1,
			expectedError: nil,
		},
		{
			id:            2,
			expectedError: nil,
		},
	}
	for _, testcases := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}

		store := taskstore.NewTaskStore(db)

		mock.ExpectExec("UPDATE task SET completed = TRUE WHERE id = ?").
			WithArgs(testcases.id).
			WillReturnResult(sqlmock.NewResult(testcases.id, 1))

		err = store.MarkComplete(int(testcases.id))
		if err != testcases.expectedError {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

func TestDelete(t *testing.T) {

	testcases := []struct {
		id            int64
		expectederror any
	}{
		{
			id:            1,
			expectederror: nil,
		},
		{
			id:            2,
			expectederror: nil,
		},
	}
	for _, testcases := range testcases {
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}

		store := taskstore.NewTaskStore(db)

		mock.ExpectExec("DELETE FROM task WHERE id = ?").
			WithArgs(testcases.id).
			WillReturnResult(sqlmock.NewResult(testcases.id, 1))

		err = store.Delete(int(testcases.id))
		if err != testcases.expectederror {
			t.Errorf("unexpected error: %v", err)
		}
	}
}
