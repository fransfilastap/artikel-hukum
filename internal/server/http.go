package server

import (
	pkgmiddlerware "bphn/artikel-hukum/internal/middleware"
	"bphn/artikel-hukum/pkg/log"
	"bphn/artikel-hukum/pkg/server"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
	"net/http"
)

type CustomValidator struct {
	validator *validator.Validate
}

func (v *CustomValidator) Validate(i interface{}) error {
	if err := v.validator.Struct(i); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return nil
}

func SetupValidator(e *echo.Echo) {
	e.Validator = &CustomValidator{validator: validator.New()}
}

func NewHttpServer(viper *viper.Viper, logger *log.Logger) *server.HttpServer {
	e := echo.New()
	e.Validator = &CustomValidator{validator: validator.New()}
	host := viper.GetString("http.host")
	port := viper.GetInt("http.port")

	// register middlewares
	jwtKey := viper.GetString("security.jwt.key")
	e.Use(middleware.RequestLoggerWithConfig(pkgmiddlerware.RequestLoggerMiddleware(logger)))
	e.Use(middleware.CORS())
	e.Use(pkgmiddlerware.JWTMiddleware(jwtKey))

	// create http server
	s := server.NewServer(e, logger, server.WithServerHost(host), server.WithServerPort(port))

	return s
}
