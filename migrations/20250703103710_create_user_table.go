package migrations

import "gofr.dev/pkg/gofr/migration"

const createTableUserSQL = `
CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL
);`

// CreateTableUser creates the 'users' table.
func CreateTableUser() migration.Migrate {
	return migration.Migrate{
		UP: func(d migration.Datasource) error {
			_, err := d.SQL.Exec(createTableUserSQL)
			if err != nil {
				return err
			}
			return nil
		},
	}
}
