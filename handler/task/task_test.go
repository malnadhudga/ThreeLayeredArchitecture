package taskhandler

import (
	models "ThreeLayeredArchitecture/models/task"
	taskservice "ThreeLayeredArchitecture/services/task"
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
	}
	return errors.New("Task not found")
}

func (m *MockTaskStore) GetByID(id int) (models.Task, error) {
	if id == 1 {
		return models.Task{ID: 1, Description: "Do homework", Completed: false}, nil
	}
	return models.Task{}, errors.New("Task not found")
}

func TestDeleteTask(t *testing.T) {
	store := &MockTaskStore{}
	service := taskservice.NewTaskService(store)

	if err := service.Delete(1); err != nil {
		t.Errorf("expected no error deleting existing task, got: %v", err)
	}

	if err := service.Delete(42); err == nil {
		t.Errorf("expected error deleting non-existent task, got nil")
	}
}

func TestAddTask(t *testing.T) {
	store := &MockTaskStore{}
	service := taskservice.NewTaskService(store)

	task, err := service.Add("Buy milk")
	if err != nil {
		t.Errorf("unexpected error adding task: %v", err)
	}
	if task.Description != "Buy milk" {
		t.Errorf("expected task description to be 'Buy milk', got '%s'", task.Description)
	} else {
		t.Error("PASS ")
	}
}

func TestGetPending(t *testing.T) {
	store := &MockTaskStore{}
	service := taskservice.NewTaskService(store)

	tasks, err := service.GetPending()
	if err != nil {
		t.Errorf("unexpected error getting pending tasks: %v", err)
	}
	if len(tasks) == 0 {
		t.Error("expected at least one pending task, got none")
	} else {
		t.Error("PASS ")
	}
}

func TestGetByID(t *testing.T) {
	store := &MockTaskStore{}
	service := taskservice.NewTaskService(store)

	task, err := service.GetByID(1)
	if err != nil {
		t.Errorf("expected no error, got: %v", err)
	}
	if task.ID != 1 {
		t.Errorf("expected task ID 1, got %d", task.ID)
	}

	_, err = service.GetByID(99)
	if err != nil {
		t.Errorf("Error %v", err)
	} else {
		t.Error("PASS ")
	}
}

//func TestMarkComplete(t *testing.T) {
//	store := &MockTaskStore{}
//	service := taskservice.NewTaskService(store)
//
//	if err := service.MarkComplete(1); err != nil {
//		t.Errorf("unexpected error completing task: %v", err)
//	}
//
//	if err := service.MarkComplete(99); err == nil {
//		t.Error("expected error marking non-existent task complete, got nil")
//	}
//}
