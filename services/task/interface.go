package task

import models "ThreeLayeredArchitecture/models/task"

type TaskStore interface {
	Add(desc string) (models.Task, error)
	GetPending() ([]models.Task, error)
	GetByID(id int) (models.Task, error)
	Delete(id int) error
	MarkComplete(id int) error
}
