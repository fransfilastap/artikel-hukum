package http

import (
	"bphn/artikel-hukum/internal/middleware"
	"bphn/artikel-hukum/internal/server"
	"bphn/artikel-hukum/pkg/config"
	jwt2 "bphn/artikel-hukum/pkg/jwt"
	"bphn/artikel-hukum/pkg/log"
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"os"
	"testing"
)

var (
	logger  *log.Logger
	handler *Handler
	e       *echo.Echo
	jwt     *jwt2.JWT
)

func TestMain(m *testing.M) {
	err := os.Setenv("APP_CONF", "../../../config/local.yml")
	if err != nil {
		fmt.Println("Setenv error", err)
	}
	var envConf = flag.String("conf", "config/local.yml", "config path, eg: -conf ./config/local.yml")
	flag.Parse()
	conf := config.NewConfig(*envConf)

	logger = log.NewLog(conf)
	handler = NewHandler(conf, logger)

	e = echo.New()
	server.SetupValidator(e)

	jwt = jwt2.NewJwt(conf)

	// register middlewares
	middleware.SetupMiddleware(conf, logger, e)

	if err != nil {
		return
	}

	code := m.Run()
	fmt.Println("test end")

	os.Exit(code)

}
