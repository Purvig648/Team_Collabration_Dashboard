package repository

import (
	"context"
	"errors"
	"log"
	dbmodel "team_collabrative_dashboard/internal/model/db_Model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (r *Repo) EnsureIndexes() {
	collection := r.db.Collection("admin")
	ctx := context.Background()

	indexModel := mongo.IndexModel{
		Keys:    bson.M{"email": 1},
		Options: options.Index().SetUnique(true),
	}

	_, err := collection.Indexes().CreateOne(ctx, indexModel)
	if err != nil {
		log.Println("Error creating unique index on email:", err)
	}
}

func (r *Repo) AdminSignUp(ctx context.Context, data dbmodel.Admin) (primitive.ObjectID, error) {
	collection := r.db.Collection("admin")

	result, err := collection.InsertOne(ctx, data)
	if err != nil {
		if mongo.IsDuplicateKeyError(err) {
			return primitive.NilObjectID, errors.New("email already exists")
		}
		return primitive.NilObjectID, err
	}

	oid, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, errors.New("failed to convert InsertedID to ObjectID")
	}
	return oid, nil
}
