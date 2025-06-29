package taskhandler

import models "ThreeLayeredArchitecture/models/task"

type Taskservice interface {
	GetPending() ([]models.Task, error)
	Add(desc string) (models.Task, error)
	Delete(id int) error
	MarkComplete(id int) (string, error)
	GetByID(id int) (models.Task, error)
}
