package migrate

import (
	"context"
	"database/sql"
	"errors"
)

func MigrateDatabase(ctx context.Context, db *sql.DB) error {
	if db == nil {
		return errors.New("invalid value provided for db")
	}

	query := `CREATE TABLE IF NOT EXISTS file_data (
		id SERIAL PRIMARY KEY,
		filename VARCHAR(255) NOT NULL,
		out_filename VARCHAR(255) NOT NULL
	);`

	_, err := db.Exec(query)

	return err
}
