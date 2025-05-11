package service

import "net/http"

type FileService interface {
	UploadFile(w http.ResponseWriter, r *http.Request)
	GetFile(w http.ResponseWriter, r *http.Request)
}
