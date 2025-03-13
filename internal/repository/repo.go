package repository

import (
	"context"
	dbmodel "team_collabrative_dashboard/internal/model/db_Model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Repo struct {
	db *mongo.Database
}

type RepoInterface interface {
	AdminSignUp(ctx context.Context, data dbmodel.Admin) (primitive.ObjectID, error)
}

func NewRepo(repo *mongo.Database) RepoInterface {
	r := &Repo{
		db: repo,
	}
	r.EnsureIndexes()
	return r
}
