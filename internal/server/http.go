package server

import (
	"bphn/artikel-hukum/internal/handler"
	mddlware "bphn/artikel-hukum/internal/middleware"
	"bphn/artikel-hukum/pkg/log"
	"bphn/artikel-hukum/pkg/server/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func NewHttpServer(viper *viper.Viper, logger *log.Logger) *http.Server {
	e := echo.New()
	host := viper.GetString("http.host")
	port := viper.GetInt("http.port")

	// register middlewares
	e.Use(middleware.RequestLoggerWithConfig(mddlware.RequestLoggerMiddleware(logger)))
	e.Use(middleware.CORS())

	e.GET("/", handler.Default)

	// create http server
	s := http.NewServer(e, logger, http.WithServerHost(host), http.WithServerPort(port))

	return s
}
