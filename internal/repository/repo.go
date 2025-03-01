package repository

import "go.mongodb.org/mongo-driver/mongo"

type Repo struct {
	client *mongo.Client
}

type RepoInterface interface {
}

func NewRepo(repo *mongo.Client) RepoInterface {
	return &Repo{
		client: repo,
	}
}
