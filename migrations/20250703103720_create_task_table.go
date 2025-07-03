package migrations

import "gofr.dev/pkg/gofr/migration"

const createTableTaskSQL = `
CREATE TABLE IF NOT EXISTS tasks (
    id INTEGER PRIMARY KEY NOT NULL,
    description TEXT NOT NULL,
    completed BOOLEAN NOT NULL
);`

// CreateTableTask creates the 'tasks' table.
func CreateTableTask() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(createTableTaskSQL)
			if err != nil {
				return err
			}
			return nil
		},
	}
}
