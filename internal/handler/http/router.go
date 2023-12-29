package http

import (
	"bphn/artikel-hukum/internal/service"
	"bphn/artikel-hukum/pkg/log"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, logger *log.Logger, userService service.UserService) {

	handler := &Handler{logger}

	// default handler
	e.GET("/", Default)

	// api group handler
	apiRoute := e.Group("/api")

	userDataRequestHandler := NewUserRequestHandler(handler, userService)

	//user group handler
	apiRoute.GET("/users", userDataRequestHandler.List)
}
