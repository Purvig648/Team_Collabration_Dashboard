package service

import "team_collabrative_dashboard/internal/repository"

type Service struct {
	repo repository.RepoInterface
}

type ServicedInterface interface {
}

func NewService(repo repository.RepoInterface) ServicedInterface {
	return &Service{
		repo: repo,
	}
}
