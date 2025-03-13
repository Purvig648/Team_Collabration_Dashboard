package service

import (
	"context"
	"team_collabrative_dashboard/internal/graph/model"
	"team_collabrative_dashboard/internal/repository"
)

type Service struct {
	repo repository.RepoInterface
}

type ServicedInterface interface {
	AdminSignUpService(ctx context.Context, input model.AdminInput) (string, error)
}

func NewService(repo repository.RepoInterface) ServicedInterface {
	return &Service{
		repo: repo,
	}
}
