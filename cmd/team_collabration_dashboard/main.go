package main

import (
	"team_collabrative_dashboard/internal/database"
	"team_collabrative_dashboard/internal/repository"
	"team_collabrative_dashboard/internal/server"
	"team_collabrative_dashboard/internal/service"
	"team_collabrative_dashboard/internal/util"

	"github.com/joho/godotenv"
	log "github.com/sirupsen/logrus"
)

func main() {
	err := godotenv.Load(util.Envpath)
	if err != nil {
		// log.WithFields(log.Fields{
		// 	"error": err.Error(),
		// }).Info("unable to load the env file")
		log.Error("unable to load the env")
	}

	db, err := database.ConnectToMongoDB()
	if err != nil {
		return
	}
	repo := repository.NewRepo(db)

	svc := service.NewService(repo)

	server.StartApp(svc)
}
