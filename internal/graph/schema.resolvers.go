package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"fmt"
	"team_collabrative_dashboard/internal/graph/model"
)

// AdminSignUp is the resolver for the AdminSignUp field.
func (r *mutationResolver) AdminSignUp(ctx context.Context, input model.AdminInput) (*model.CreateAdminResponse, error) {
	adminID, err := r.Svc.AdminSignUpService(ctx, input)
	if err != nil {
		return &model.CreateAdminResponse{
			Success: false,
			Error:   err.Error(),
			Message: "Admin registration failed",
			Admin:   "",
		}, nil
	}

	return &model.CreateAdminResponse{
		Success: true,
		Message: "Admin registered successfully",
		Admin:   adminID,
	}, nil
}

// AdminSignIn is the resolver for the AdminSignIn field.
func (r *mutationResolver) AdminSignIn(ctx context.Context, input model.AdminSignIn) (*model.AdminSignInResponse, error) {
	panic(fmt.Errorf("not implemented: AdminSignIn - AdminSignIn"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
