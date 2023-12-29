package http

import (
	custommiddlerware "bphn/artikel-hukum/internal/middleware"
	"bphn/artikel-hukum/pkg/config"
	"bphn/artikel-hukum/pkg/log"
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	logger  *log.Logger
	handler *Handler
	e       *echo.Echo
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
	handler = &Handler{logger}

	e = echo.New()
	// register middlewares
	e.Use(middleware.RequestLoggerWithConfig(custommiddlerware.RequestLoggerMiddleware(logger)))
	e.Use(middleware.CORS())

	code := m.Run()
	fmt.Println("test end")

	os.Exit(code)

}

func TestUserHandler(t *testing.T) {
	t.Run("should return 200 http code", func(t *testing.T) {

		req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		userHandler := NewUserRequestHandler(handler)

		err := userHandler.List(c)
		if err != nil {
			panic(err)
		}

		assert.Equal(t, http.StatusOK, rec.Code)

	})
}
