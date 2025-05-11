package service

import (
	"database/sql"
	"net/http"
	"quick-drop-be/internals/repo"
)

type fileServiceImpl struct {
	repo repo.FileRepo
}

func NewFileService(db *sql.DB) (FileService, error) {
	repo, err := repo.NewFileRepo(db)
	if err != nil {
		return nil, err
	}
	return &fileServiceImpl{repo}, nil
}

func (s *fileServiceImpl) GetFile(
	w http.ResponseWriter, r *http.Request,
) {

}

func (s *fileServiceImpl) UploadFile(
	w http.ResponseWriter, r *http.Request,
) {
}
