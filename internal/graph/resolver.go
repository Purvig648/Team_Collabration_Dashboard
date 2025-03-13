package graph

import "team_collabrative_dashboard/internal/service"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Svc service.ServicedInterface
}
