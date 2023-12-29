package http

import (
	internalhttp "bphn/artikel-hukum/internal/handler/http"
	internal "bphn/artikel-hukum/internal/middleware"
	"bphn/artikel-hukum/internal/routes"
	"bphn/artikel-hukum/pkg/log"
	"bphn/artikel-hukum/pkg/server/http"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func NewHttpServer(viper *viper.Viper, logger *log.Logger) *http.Server {
	e := echo.New()
	e.Validator = &Validator{validator: validator.New()}
	host := viper.GetString("http.host")
	port := viper.GetInt("http.port")

	// register middlewares
	internal.SetupMiddleware(viper, logger, e)

	// routes
	routes.SetupRoutes(e, &internalhttp.Handler{Logger: logger})

	// create http server
	s := http.NewServer(e, logger, http.WithServerHost(host), http.WithServerPort(port))

	return s
}
