package server

import (
	internal "bphn/artikel-hukum/internal/middleware"
	"bphn/artikel-hukum/pkg/log"
	"bphn/artikel-hukum/pkg/server"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func NewHttpServer(viper *viper.Viper, logger *log.Logger) *server.HttpServer {
	e := echo.New()
	e.Validator = &CValidator{Validator: validator.New()}
	host := viper.GetString("http.host")
	port := viper.GetInt("http.port")

	// register middlewares
	internal.SetupMiddleware(viper, logger, e)

	// create http server
	s := server.NewServer(e, logger, server.WithServerHost(host), server.WithServerPort(port))

	return s
}
