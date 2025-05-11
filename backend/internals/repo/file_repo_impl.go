package repo

import (
	"context"
	"database/sql"
	"errors"
	"time"
)

const queryTimeout = time.Second * 5

type fileRepo struct {
	db *sql.DB
}

func NewFileRepo(db *sql.DB) (FileRepo, error) {
	if db == nil {
		return nil, errors.New("nil value provided for db")
	}

	return &fileRepo{db}, nil
}

func (r *fileRepo) AddFileMetadata(
	ctx context.Context, filename, outFilename string,
) (err error) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	query := `INSERT INTO file_data (filename, out_filename)
	VALUES ($1, $2);`

	_, err = r.db.ExecContext(ctx, query, filename, outFilename)

	return
}

func (r *fileRepo) GetFileMetadata(ctx context.Context, outFilename string) (
	filename string, err error,
) {
	ctx, cancel := context.WithTimeout(ctx, queryTimeout)
	defer cancel()

	query := `SELECT filename FROM file_data WHERE out_filename = $1;`
	err = r.db.QueryRowContext(ctx, query, outFilename).Scan(&filename)

	return
}
