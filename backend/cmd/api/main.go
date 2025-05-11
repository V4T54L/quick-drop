package main

import (
	"log"
	"net/http"
	"quick-drop-be/internals/dbconnectors"
	"quick-drop-be/internals/router"
	"quick-drop-be/internals/server"

	"github.com/go-chi/chi/v5"
)

func registerRoutes(r *chi.Mux) {
	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("healthy"))
	})
}

func main() {
	postgresConn, err := dbconnectors.GetPostgresDb("postgres://user:password@localhost:5432/db?sslmode=disable")
	if err != nil {
		log.Println("Error connecting to postgreSQL db: ", err)
	} else {
		log.Println("PostgreSQL connected successfully! ")
	}

	defer func() {
		if err = dbconnectors.CloseDBConn(postgresConn); err != nil {
			log.Println("Error closing postgreSQL db: ", err)
		}
	}()

	r := router.NewChiRouter()

	// TODO: Use Database in the routers
	registerRoutes(r)

	s := http.Server{
		Addr:    ":8000",
		Handler: r,
	}

	server.RunWithGracefulShutdown(&s)
}
