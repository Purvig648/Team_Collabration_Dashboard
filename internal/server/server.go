package server

import (
	"log"
	"team_collabrative_dashboard/internal/graph"
	"team_collabrative_dashboard/internal/service"

	// Update with actual module path

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/labstack/echo/v4"
)

func StartApp(svc service.ServicedInterface) {
	e := echo.New()
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Svc: svc,
	}}))

	e.POST("/query", func(c echo.Context) error {
		srv.ServeHTTP(c.Response(), c.Request())
		return nil
	})

	e.GET("/", func(c echo.Context) error {
		playground.Handler("GraphQL Playground", "/query").ServeHTTP(c.Response(), c.Request())
		return nil
	})

	log.Println("Server is running on http://localhost:8080")
	if err := e.Start(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
