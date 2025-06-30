package taskstore_test

import (
	"ThreeLayeredArchitecture/store/task"
	"github.com/DATA-DOG/go-sqlmock"
	"testing"
)

func TestAdd(t *testing.T) {
	testcases := []struct {
		desc       string
		completed  bool
		expectedID int64
	}{
		{
			desc:       "Clean room",
			completed:  false,
			expectedID: 1,
		},
		{
			desc:       "Play",
			completed:  false,
			expectedID: 2,
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
			WillReturnResult(sqlmock.NewResult(testcases.expectedID, 1))

		task, err := store.Add(testcases.desc)
		if err != nil {
			t.Errorf("expected no error, got %v", err)
		}
		if int64(task.ID) != testcases.expectedID || task.Description != testcases.desc || task.Completed {
			t.Errorf("unexpected task: %+v", task)
		} else {
			t.Log("PASS")
		}
	}

}

func TestGetByID(t *testing.T) {

	testcases := []struct {
		desc       string
		completed  bool
		expectedID int64
	}{
		{
			desc:       "Clean room",
			completed:  false,
			expectedID: 1,
		},
		{
			desc:       "Play",
			completed:  false,
			expectedID: 2,
		},
		{
			desc:       "Play",
			completed:  false,
			expectedID: 3,
		},
	}

	for i, testcases := range testcases {
		i = i + 1
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}

		store := taskstore.NewTaskStore(db)

		rows := sqlmock.NewRows([]string{"id", "description", "completed"}).
			AddRow(i, testcases.desc, testcases.completed)

		mock.ExpectQuery("SELECT id, description, completed FROM task WHERE id = ?").
			WithArgs(1).
			WillReturnRows(rows)

		task, err := store.GetByID(1)
		if err != nil || int64(task.ID) != testcases.expectedID {
			t.Errorf("unexpected result: %v, err: %v", task, err)
		}
	}
}

func TestGetPending(t *testing.T) {
	testcases := []struct {
		desc       string
		completed  bool
		expectedID int64
	}{
		{
			desc:       "Clean room",
			completed:  false,
			expectedID: 1,
		},
		{
			desc:       "Play",
			completed:  false,
			expectedID: 2,
		},
		{
			desc:       "Dance",
			completed:  true,
			expectedID: 3,
		},
	}
	for i, testcases := range testcases {
		i = i + 1
		db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
		if err != nil {
			t.Errorf("unexpected error: %s", err)
		}

		store := taskstore.NewTaskStore(db)

		rows := sqlmock.NewRows([]string{"id", "description", "completed"}).
			AddRow(i, testcases.desc, testcases.completed)

		mock.ExpectQuery("SELECT id, description, completed FROM task WHERE completed = FALSE ORDER BY id").
			WillReturnRows(rows)

		tasks, err := store.GetPending()
		if err != nil || len(tasks) != 1 {
			t.Errorf("unexpected result: %v, err: %v", tasks, err)
		}
		if int64(tasks[0].ID) != testcases.expectedID || tasks[0].Description != testcases.desc || tasks[0].Completed {
			t.Errorf("unexpected task: %+v", tasks[0])
		}
	}
}

func TestMarkComplete(t *testing.T) {

	testcases := []struct {
		id int64
	}{
		{
			id: 1,
		},
		{
			id: 2,
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
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	}
}

func TestDelete(t *testing.T) {

	testcases := []struct {
		id int64
	}{
		{
			id: 1,
		},
		{
			id: 2,
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
		if err != nil {
			t.Errorf("unexpected error: %v", err)

		}
	}
}
