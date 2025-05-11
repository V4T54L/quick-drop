package repo

import (
	"database/sql"
	"errors"
)

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
	filename, filepath string,
) (id int, err error) {
	return 0, nil
}

func (r *fileRepo) GetFileMetadata(id int) (
	filename, filepath string, err error,
) {
	return "", "", nil
}
