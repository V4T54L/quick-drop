package service

import (
	"database/sql"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"quick-drop-be/internals/config"
	"quick-drop-be/internals/repo"

	"github.com/google/uuid"
)

const MAX_UPLOAD_SIZE = 1024 * 1524 // about 1.5 MB

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
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "error parsing the file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	if header.Size > MAX_UPLOAD_SIZE {
		http.Error(w, "error file size exceeds the limit", http.StatusBadRequest)
		return
	}

	// Create an uploads directory if it doesn't exist
	uploadsDir := "./uploads"
	err = os.MkdirAll(uploadsDir, os.ModePerm)
	if err != nil {
		http.Error(w, "Failed to create uploads directory", http.StatusInternalServerError)
		return
	}

	outFileName := uuid.New().String()
	filePath := filepath.Join(uploadsDir, outFileName)

	outFile, err := os.Create(filePath)
	if err != nil {
		http.Error(w, "Failed to create file on the server", http.StatusInternalServerError)
		return
	}
	defer outFile.Close()

	_, err = io.Copy(outFile, file)
	if err != nil {
		http.Error(w, "Failed to save file", http.StatusInternalServerError)
		return
	}

	err = s.repo.AddFileMetadata(r.Context(), header.Filename, outFileName)
	if err != nil {
		http.Error(w, "Error querying the database", http.StatusInternalServerError)
		return
	}

	fileURL := fmt.Sprintf("%s/files/%s", config.GetConfig().ServerURL, outFileName)
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"fileUrl": "%s"}`, fileURL)
}
