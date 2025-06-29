package taskhandler

import models "ThreeLayeredArchitecture/models/task"

type Taskservice interface {
	GetPendingTasks() ([]models.Task, error)
	AddTask(desc string) (models.Task, error)
	DeleteTask(id int) error
	CompleteTask(id int) (string, error)
	GetTaskByID(id int) (models.Task, error)
}
