package task

import (
	models "ThreeLayeredArchitecture/models/task"
	"gofr.dev/pkg/gofr"
)

type TaskStore interface {
	Add(ctx *gofr.Context, input models.Task) (models.Task, error)
	GetPending(ctx *gofr.Context) ([]models.Task, error)
	GetByID(ctx *gofr.Context, id int) (models.Task, error)
	Delete(ctx *gofr.Context, id int) error
	MarkComplete(ctx *gofr.Context, id int) error
}
