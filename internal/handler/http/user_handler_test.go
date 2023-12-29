package http

import (
	"bphn/artikel-hukum/api"
	"bphn/artikel-hukum/internal/handler/http/fakes"
	custommiddlerware "bphn/artikel-hukum/internal/middleware"
	"bphn/artikel-hukum/pkg/config"
	"bphn/artikel-hukum/pkg/log"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
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
	custommiddlerware.SetupMiddleware(conf, logger, e)

	code := m.Run()
	fmt.Println("test end")

	os.Exit(code)

}

func TestUserRequestHandler_List(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userHandler := NewUserRequestHandler(handler, &fakes.FakeUserService{})

	err := userHandler.List(c)
	if err != nil {
		panic(err)
	}

	var users []api.UserDataResponse
	unmarshalErr := json.Unmarshal(rec.Body.Bytes(), &users)

	if unmarshalErr != nil {
		panic(unmarshalErr)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, len(users), 3)
}

func TestUserRequestHandler_Create(t *testing.T) {
	userJSON := `{"full_name":"Jon Snow","email":"jon@labstack.com","password":"password","role":"editor"}`
	req := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(userJSON))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	userHandler := NewUserRequestHandler(handler, &fakes.FakeUserService{})

	if err := userHandler.Create(c); err != nil {
		panic(err)
	}
}
