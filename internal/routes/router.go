package routes

import (
	"bphn/artikel-hukum/internal/handler/http"
	"github.com/labstack/echo/v4"
)

func SetupRoutes(e *echo.Echo, handler *http.Handler) {

	// default handler
	e.GET("/", http.Default)

	// api group handler
	apiRoute := e.Group("/api")

	userRequestHandler := http.NewUserRequestHandler(handler)

	//user group handler
	apiRoute.GET("/users", userRequestHandler.List)
}
