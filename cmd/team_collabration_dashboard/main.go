package main

import (
	"team_collabrative_dashboard/internal/database"
	"team_collabrative_dashboard/internal/repository"
	"team_collabrative_dashboard/internal/server"
	"team_collabrative_dashboard/internal/service"
)

func main() {
	db, err := database.ConnectToMongoDB()
	if err != nil {
		return
	}
	repo := repository.NewRepo(db)

	svc := service.NewService(repo)

	server.StartApp(svc)
}
