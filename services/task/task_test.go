package taskservice

import (
	models "ThreeLayeredArchitecture/models/task"
	"errors"
	"testing"
)

type MockTaskStore struct{}

func (m *MockTaskStore) GetPending() ([]models.Task, error) {
	return []models.Task{
		{ID: 1, Description: "Do homework", Completed: false},
	}, nil
}

func (m *MockTaskStore) Add(desc string) (models.Task, error) {
	return models.Task{ID: 2, Description: desc, Completed: false}, nil
}

func (m *MockTaskStore) Delete(id int) error {
	if id == 1 {
		return nil
	}
	return errors.New("Task not found")
}

func (m *MockTaskStore) MarkComplete(id int) error {
	if id == 1 {
		return nil
	} else {
		return errors.New("Task not found")
	}
}

func (m *MockTaskStore) GetByID(id int) (models.Task, error) {
	if id == 1 {
		return models.Task{ID: 1, Description: "Do homework", Completed: false}, nil
	}
	return models.Task{}, errors.New("Task not found")
}

func TestAddTask(t *testing.T) {
	testcases := []struct {
		desc string
	}{
		{desc: "play"},
		{desc: "jogging"},
	}

	store := &MockTaskStore{}
	service := NewTaskService(store)

	for _, tc := range testcases {
		task, err := service.Add(tc.desc)
		if err != nil {
			t.Errorf("unexpected error adding task: %v", err)
		}
		if task.Description != tc.desc {
			t.Errorf("expected task description to be '%s', got '%s'", tc.desc, task.Description)
		} else {
			t.Log("PASS")
		}
	}
}

func TestDeleteTask(t *testing.T) {
	testcases := []struct {
		id          int
		expectError bool
	}{
		{id: 1, expectError: false},
		{id: 42, expectError: true},
	}

	store := &MockTaskStore{}
	service := NewTaskService(store)

	for _, tc := range testcases {
		err := service.Delete(tc.id)
		if tc.expectError && err == nil {
			t.Errorf("expected error for id=%d, got nil", tc.id)
		}
		if !tc.expectError && err != nil {
			t.Errorf("unexpected error for id=%d: %v", tc.id, err)
		}
	}
}

func TestGetPending(t *testing.T) {
	store := &MockTaskStore{}
	service := NewTaskService(store)

	tasks, err := service.GetPending()
	if err != nil {
		t.Errorf("unexpected error getting pending tasks: %v", err)
	}
	if len(tasks) == 0 {
		t.Error("expected at least one pending task, got none")
	} else {
		t.Log("PASS")
	}
}

func TestGetByID(t *testing.T) {
	testcases := []struct {
		id          int
		expectError bool
	}{
		{id: 1, expectError: false},
		{id: 99, expectError: true},
	}

	store := &MockTaskStore{}
	service := NewTaskService(store)

	for _, tc := range testcases {
		task, err := service.GetByID(tc.id)
		if tc.expectError && err == nil {
			t.Errorf("expected error for id=%d, got nil", tc.id)
		}
		if !tc.expectError && err != nil {
			t.Errorf("unexpected error for id=%d: %v", tc.id, err)
		}
		if !tc.expectError && task.ID != tc.id {
			t.Errorf("expected ID=%d, got %d", tc.id, task.ID)
		}
	}
}

func TestMarkComplete(t *testing.T) {
	testcases := []struct {
		id          int
		expectError bool
	}{
		{id: 1, expectError: false},
		{id: 4, expectError: true},
	}

	store := &MockTaskStore{}
	service := NewTaskService(store)

	for _, tc := range testcases {
		_, err := service.MarkComplete(tc.id)
		if tc.expectError && err == nil {
			t.Errorf("expected error for id=%d, got nil", tc.id)
		}
		if !tc.expectError && err != nil {
			t.Errorf("unexpected error for id=%d: %v", tc.id, err)
		} else if !tc.expectError {
			t.Log("PASS")
		}
	}
}
