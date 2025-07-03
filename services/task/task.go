package task

import (
	"ThreeLayeredArchitecture/models/task"
	"gofr.dev/pkg/gofr"
)

type TaskService struct {
	Store TaskStore
}

func NewTaskService(store TaskStore) *TaskService {
	return &TaskService{Store: store}
}

func (s *TaskService) Add(ctx *gofr.Context, input models.Task) (models.Task, error) {
	return s.Store.Add(ctx, input)
}

func (s *TaskService) GetPending(ctx *gofr.Context) ([]models.Task, error) {
	return s.Store.GetPending(ctx)
}

func (s *TaskService) GetByID(ctx *gofr.Context, id int) (models.Task, error) {
	return s.Store.GetByID(ctx, id)
}

func (s *TaskService) MarkComplete(ctx *gofr.Context, id int) (string, error) {
	task, err := s.Store.GetByID(ctx, id)
	if err != nil {
		return "", err
	}

	if task.Completed {
		return "Task already completed", nil
	}

	err = s.Store.MarkComplete(ctx, id)
	if err != nil {
		return "", err
	}

	return "Task marked as complete", nil
}

func (s *TaskService) Delete(ctx *gofr.Context, id int) (string, error) {
	_, err := s.Store.GetByID(ctx, id)
	if err != nil {
		return "No task Found", err
	}

	err = s.Store.Delete(ctx, id)

	if err != nil {
		return "No task Found", err
	}

	return "Task is Successfully Deleted", nil
}
