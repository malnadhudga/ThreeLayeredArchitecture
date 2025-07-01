package task

import (
	"ThreeLayeredArchitecture/models/task"
)

type TaskService struct {
	Store TaskStore
}

func NewTaskService(store TaskStore) *TaskService {
	return &TaskService{Store: store}
}

func (s *TaskService) Add(desc string) (models.Task, error) {
	return s.Store.Add(desc)
}

func (s *TaskService) GetPending() ([]models.Task, error) {
	return s.Store.GetPending()
}

func (s *TaskService) GetByID(id int) (models.Task, error) {
	return s.Store.GetByID(id)
}

func (s *TaskService) MarkComplete(id int) (string, error) {
	task, err := s.Store.GetByID(id)
	if err != nil {
		return "", err
	}

	if task.Completed {
		return "Task already completed", nil
	}

	err = s.Store.MarkComplete(id)
	if err != nil {
		return "", err
	}

	return "Task marked as complete", nil
}

func (s *TaskService) Delete(id int) error {
	return s.Store.Delete(id)
}
