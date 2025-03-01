package server

import (
	"errors"
	"net/http"
	"team_collabrative_dashboard/internal/service"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

// Initialize Logrus
var log = logrus.New()

func StartApp(svc service.ServicedInterface) {
	server := echo.New()
	server.GET("/", hello)

	// Start server with Logrus logging
	if err := server.Start(":8081"); err != nil && !errors.Is(err, http.ErrServerClosed) {
		log.Error("Failed to start server: ", err)
	}
}

// Handler function
func hello(c echo.Context) error {
	log.Info("Received request on /")
	return c.String(http.StatusOK, "Hello, World!")
}
