package http

import (
	"bphn/artikel-hukum/internal/service"
	"bphn/artikel-hukum/pkg/log"
	"bphn/artikel-hukum/pkg/server"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
	"net/http"
)

func RegisterRoutes(server *server.HttpServer, viper *viper.Viper, logger *log.Logger, userService service.UserService, authorService service.AuthorService) {
	e := server.Engine()
	handler := NewHandler(viper, logger)

	// default handler
	e.GET("/", func(c echo.Context) error {
		err := c.JSON(http.StatusOK, "Oh, Hey!")

		if err != nil {
			return err
		}
		return nil
	})

	// api group handler
	apiRoute := e.Group("/api")

	userDataRequestHandler := NewUserManagementHandler(handler, userService)

	//user group handler
	apiRoute.GET("/users", userDataRequestHandler.List)
	apiRoute.POST("/users", userDataRequestHandler.Create)
	apiRoute.PUT("/users/:id", userDataRequestHandler.Update)
	apiRoute.DELETE("/users/:id", userDataRequestHandler.Delete)
	apiRoute.POST("/users/forgot-password", userDataRequestHandler.ForgotPassword)

	authorManagementHandler := NewAuthorManagementHandler(handler, authorService)

	apiRoute.GET("/author/profile", authorManagementHandler.GetProfile)
	apiRoute.POST("/author", authorManagementHandler.Register)
	apiRoute.PUT("/author/:id", authorManagementHandler.UpdateProfile)

}
