package taskstore

import (
	"ThreeLayeredArchitecture/models/task"
	"fmt"
	"gofr.dev/pkg/gofr"
)

type TaskStore struct{}

func NewTaskStore() *TaskStore {
	return &TaskStore{}
}

func (*TaskStore) Add(ctx *gofr.Context, input models.Task) (models.Task, error) {
	res, err := ctx.SQL.Exec("INSERT INTO task (description, completed) VALUES (?, ?)", input.Description, input.Completed)
	if err != nil {
		return models.Task{}, fmt.Errorf("failed to add task: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return models.Task{}, err
	}

	input.ID = int(id)

	return input, nil
}

func (*TaskStore) GetPending(ctx *gofr.Context) ([]models.Task, error) {
	query := "SELECT id, description, completed FROM task WHERE completed = FALSE ORDER BY id"
	rows, err := ctx.SQL.Query(query)

	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, rows.Err()
	}

	var tasks []models.Task

	for rows.Next() {
		var task models.Task
		err := rows.Scan(&task.ID, &task.Description, &task.Completed)

		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (*TaskStore) GetByID(ctx *gofr.Context, id int) (models.Task, error) {
	query := "SELECT id, description, completed FROM task WHERE id = ?"
	row := ctx.SQL.QueryRow(query, id)

	var task models.Task
	err := row.Scan(&task.ID, &task.Description, &task.Completed)

	return task, err
}

func (*TaskStore) MarkComplete(ctx *gofr.Context, id int) error {
	query := "UPDATE task SET completed = TRUE WHERE id = ?"
	_, err := ctx.SQL.Exec(query, id)

	return err
}

func (*TaskStore) Delete(ctx *gofr.Context, id int) error {
	query := "DELETE FROM task WHERE id = ?"
	_, err := ctx.SQL.Exec(query, id)

	return err
}
