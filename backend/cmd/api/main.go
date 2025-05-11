package main

import (
	"log"
	"net/http"
	"quick-drop-be/internals/config"
	"quick-drop-be/internals/dbconnectors"
	"quick-drop-be/internals/router"
	"quick-drop-be/internals/server"
	"quick-drop-be/internals/service"
)

func main() {
	postgresConn, err := dbconnectors.GetPostgresDb(config.GetConfig().DBURI)
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

	fileService, err := service.NewFileService(postgresConn)
	if err != nil {
		log.Fatal("Error initializing file service: ", err.Error())
	}

	r := router.NewChiRouter()

	router.RegisterRoutes(r, fileService)

	s := http.Server{
		Addr:    config.GetConfig().PORT,
		Handler: r,
	}

	server.RunWithGracefulShutdown(&s)
}
