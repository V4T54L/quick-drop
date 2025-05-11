package router

import (
	"net/http"
	"quick-drop-be/internals/service"

	"github.com/go-chi/chi/v5"
)

func RegisterRoutes(r *chi.Mux, fileService service.FileService) {
	r.Get("/health", healthHandler)

	r.Get("/files/{fileId}", fileService.GetFile)
	r.Post("/files", fileService.UploadFile)
}

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("healthy"))
}
