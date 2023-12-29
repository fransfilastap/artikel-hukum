package middleware

import (
	"bphn/artikel-hukum/pkg/log"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func SetupMiddleware(viper *viper.Viper, logger *log.Logger, e *echo.Echo) {
	e.Use(middleware.RequestLoggerWithConfig(RequestLoggerMiddleware(logger)))
	e.Use(middleware.CORS())
}
