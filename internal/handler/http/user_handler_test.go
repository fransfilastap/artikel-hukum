package http

import (
	"bphn/artikel-hukum/api"
	"bphn/artikel-hukum/internal/middleware"
	internalhttp "bphn/artikel-hukum/internal/server"
	mockservice "bphn/artikel-hukum/internal/service/mocks"
	"bphn/artikel-hukum/pkg/config"
	"bphn/artikel-hukum/pkg/log"
	"encoding/json"
	"flag"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
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
	e.Validator = &internalhttp.CValidator{Validator: validator.New()}
	// register middlewares
	middleware.SetupMiddleware(conf, logger, e)

	if err != nil {
		return
	}

	code := m.Run()
	fmt.Println("test end")

	os.Exit(code)

}

func TestUserRequestHandler_List(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mockservice.NewMockUserService(ctrl)
	mockUserService.EXPECT().List(c.Request().Context()).Return([]api.UserDataResponse{
		{
			Id:       1,
			FullName: "Frans Filasta Pratama",
			Email:    "mail@fransfp.dev",
			Avatar:   "https://ui-avatars.com/api/?uppercase=false&name=frans",
			Role:     "admin",
		},
		{
			Id:       2,
			FullName: "Rahma Fitri",
			Email:    "rahmafitri92@gmail.com",
			Avatar:   "https://ui-avatars.com/api/?uppercase=false&name=rahma+fitri",
			Role:     "author",
		},
		{
			Id:       3,
			FullName: "Ibrahim Finra Achernar",
			Email:    "finn@fransfp.dev",
			Avatar:   "https://ui-avatars.com/api/?uppercase=false&name=finn",
			Role:     "author",
		},
	}, nil).AnyTimes()

	userHandler := NewUserRequestHandler(handler, mockUserService)

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

	t.Run("Validation error while create user", func(t *testing.T) {
		userJSON := `{"full_name":"Jon Snow","email":"jon@labstack.com","password":"123456","role":"editor"}`
		req := httptest.NewRequest(http.MethodPost, "/api/users", strings.NewReader(userJSON))
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		ctrl := gomock.NewController(t)
		defer ctrl.Finish()

		mockService := mockservice.NewMockUserService(ctrl)
		mockService.EXPECT().Create(c.Request().Context(), gomock.Any()).Return(nil).AnyTimes()

		userHandler := NewUserRequestHandler(handler, mockService)

		err := userHandler.Create(c)

		assert.Error(t, err)

	})

}
