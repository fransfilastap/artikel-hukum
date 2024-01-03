package http

import (
	"bphn/artikel-hukum/api/v1"
	errors2 "bphn/artikel-hukum/internal/errors"
	"bphn/artikel-hukum/internal/ito"
	"bphn/artikel-hukum/internal/middleware"
	"bphn/artikel-hukum/internal/server"
	mockservice "bphn/artikel-hukum/internal/service/mocks"
	"bphn/artikel-hukum/pkg/config"
	pkgjwt "bphn/artikel-hukum/pkg/jwt"
	"bphn/artikel-hukum/pkg/log"
	"bytes"
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var (
	logger  *log.Logger
	handler *Handler
	e       *echo.Echo
	jwt     *pkgjwt.JWT
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

	jwt = pkgjwt.NewJwt(conf)

	// register middlewares
	middleware.SetupMiddleware(conf, logger, e)

	if err != nil {
		return
	}

	code := m.Run()
	fmt.Println("test end")

	os.Exit(code)

}

func TestUserManagementHandler_ForgotPassword(t *testing.T) {

	t.Run("Success update password", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		forgotPasswordRequest := v1.ForgotPasswordRequest{Email: "mail@johndoe.com"}
		forgotPasswordRequestJson, _ := json.Marshal(forgotPasswordRequest)

		userServiceMock := mockservice.NewMockUserService(ctrl)
		userServiceMock.EXPECT().ForgotPassword(gomock.Any(), forgotPasswordRequest).Return(nil).Times(1)

		req := httptest.NewRequest(http.MethodPut, "/api/users/forgot-password", bytes.NewReader(forgotPasswordRequestJson))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+generateToken(t))
		rec := httptest.NewRecorder()

		handler := NewUserManagementHandler(handler, userServiceMock)

		e.PUT("/api/users/forgot-password", handler.ForgotPassword)
		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusOK, rec.Code)

	})

	t.Run("Failed to update password due to validation error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		forgotPasswordRequest := v1.ForgotPasswordRequest{Email: "mail"}
		forgotPasswordRequestJson, _ := json.Marshal(forgotPasswordRequest)

		userServiceMock := mockservice.NewMockUserService(ctrl)
		//userServiceMock.EXPECT().ForgotPassword(gomock.Any(), forgotPasswordRequest)

		req := httptest.NewRequest(http.MethodPut, "/api/users/forgot-password", bytes.NewReader(forgotPasswordRequestJson))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+generateToken(t))
		rec := httptest.NewRecorder()

		handler := NewUserManagementHandler(handler, userServiceMock)

		e.PUT("/api/users/forgot-password", handler.ForgotPassword)
		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusBadRequest, rec.Code)

	})

	t.Run("Failed to update password due to email doesn't exists", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		forgotPasswordRequest := v1.ForgotPasswordRequest{Email: "mail@johndoe.com"}
		forgotPasswordRequestJson, _ := json.Marshal(forgotPasswordRequest)

		userServiceMock := mockservice.NewMockUserService(ctrl)
		userServiceMock.EXPECT().ForgotPassword(gomock.Any(), forgotPasswordRequest).Return(errors2.ErrUserDoesNotExists).Times(1)

		req := httptest.NewRequest(http.MethodPut, "/api/users/forgot-password", bytes.NewReader(forgotPasswordRequestJson))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+generateToken(t))
		rec := httptest.NewRecorder()

		handler := NewUserManagementHandler(handler, userServiceMock)

		e.PUT("/api/users/forgot-password", handler.ForgotPassword)
		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusNotAcceptable, rec.Code)

	})

	t.Run("Failed to update password due to other error", func(t *testing.T) {
		ctrl := gomock.NewController(t)

		forgotPasswordRequest := v1.ForgotPasswordRequest{Email: "mail@johndoe.com"}
		forgotPasswordRequestJson, _ := json.Marshal(forgotPasswordRequest)

		userServiceMock := mockservice.NewMockUserService(ctrl)
		userServiceMock.EXPECT().ForgotPassword(gomock.Any(), forgotPasswordRequest).Return(errors.New("Database error")).Times(1)

		req := httptest.NewRequest(http.MethodPut, "/api/users/forgot-password", bytes.NewReader(forgotPasswordRequestJson))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		req.Header.Set(echo.HeaderAuthorization, "Bearer "+generateToken(t))
		rec := httptest.NewRecorder()

		handler := NewUserManagementHandler(handler, userServiceMock)

		e.PUT("/api/users/forgot-password", handler.ForgotPassword)
		e.ServeHTTP(rec, req)

		assert.Equal(t, http.StatusInternalServerError, rec.Code)

	})

}

func TestAuthorManagementHandler_ChangePasswordSuccess(t *testing.T) {

	controller := gomock.NewController(t)

	password := v1.ChangePasswordRequest{
		CurrentPassword: "12345678",
		NewPassword:     "87654321",
	}

	requestJSON, _ := json.Marshal(password)

	req := httptest.NewRequest(http.MethodPut, "/api/users/change-password", bytes.NewReader(requestJSON))
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+generateToken(t))
	rec := httptest.NewRecorder()
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	userServiceMock := mockservice.NewMockUserService(controller)
	userServiceMock.EXPECT().ChangePasswordByNonAdmin(gomock.Any(), gomock.Any()).Return(nil).Times(1)

	handler := NewUserManagementHandler(handler, userServiceMock)

	e.PUT("/api/users/change-password", handler.ChangePasswordByNonAdmin)
	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

}

func TestUserRequestHandler_Create(t *testing.T) {

	t.Run("Success  create user, return http code 200, no error", func(t *testing.T) {

		userRequest := v1.CreateUserRequest{
			FullName: "John Snow",
			Email:    "snow@mail.com",
			Password: "12345678",
			Role:     "editor",
		}

		userRequestJSON, _ := json.Marshal(userRequest)

		req := httptest.NewRequest(http.MethodPost, "/api/users", bytes.NewReader(userRequestJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		ctrl := gomock.NewController(t)

		mockService := mockservice.NewMockUserService(ctrl)
		mockService.EXPECT().Create(c.Request().Context(), gomock.Any()).Return(nil).AnyTimes()

		userHandler := NewUserManagementHandler(handler, mockService)

		if assert.NoError(t, userHandler.Create(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
		}

	})

	t.Run("Validation error while creating user", func(t *testing.T) {
		userRequest := v1.CreateUserRequest{
			FullName: "John Snow",
			Email:    "mail@johnsnow.techx",
			Password: "123456",
			Role:     "editor",
		}

		userRequestJSON, _ := json.Marshal(userRequest)

		req := httptest.NewRequest(http.MethodPost, "/api/users", bytes.NewReader(userRequestJSON))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		ctrl := gomock.NewController(t)

		mockService := mockservice.NewMockUserService(ctrl)
		mockService.EXPECT().Create(gomock.Any(), userRequest).Return(errors.New("validation error")).AnyTimes()

		userHandler := NewUserManagementHandler(handler, mockService)

		if assert.NoError(t, userHandler.Create(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}

	})

}

func TestUserRequestHandler_List(t *testing.T) {

	t.Run("basic list query", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		ctrl := gomock.NewController(t)

		mockUserService := mockservice.NewMockUserService(ctrl)
		mockUserService.EXPECT().List(c.Request().Context(), gomock.Any()).Return(&ito.ListQueryResult[ito.UserDataResponse]{
			TotalPage: 1,
			Page:      1,
			Items: []ito.UserDataResponse{
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
			},
		}, nil).AnyTimes()

		userHandler := NewUserManagementHandler(handler, mockUserService)

		err := userHandler.List(c)
		if err != nil {
			panic(err)
		}

		var users ito.ListQueryResult[ito.UserDataResponse]
		unmarshalErr := json.Unmarshal(rec.Body.Bytes(), &users)

		if unmarshalErr != nil {
			panic(unmarshalErr)
		}

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, 3, len(users.Items))
	})

	t.Run("List with filter and options", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/api/users?page=1&size=1&sort=email", nil)
		rec := httptest.NewRecorder()
		c := e.NewContext(req, rec)

		ctrl := gomock.NewController(t)

		mockUserService := mockservice.NewMockUserService(ctrl)
		mockUserService.EXPECT().List(c.Request().Context(), gomock.Any()).Return(&ito.ListQueryResult[ito.UserDataResponse]{
			TotalPage: 1,
			Page:      1,
			Items: []ito.UserDataResponse{
				{
					Id:       1,
					FullName: "Frans Filasta Pratama",
					Email:    "mail@fransfp.dev",
					Avatar:   "https://ui-avatars.com/api/?uppercase=false&name=frans",
					Role:     "admin",
				},
			},
		}, nil).AnyTimes()

		userHandler := NewUserManagementHandler(handler, mockUserService)

		err := userHandler.List(c)
		if err != nil {
			panic(err)
		}

		var users ito.ListQueryResult[ito.UserDataResponse]
		unmarshalErr := json.Unmarshal(rec.Body.Bytes(), &users)

		if unmarshalErr != nil {
			panic(unmarshalErr)
		}

		assert.Equal(t, http.StatusOK, rec.Code)
		assert.Equal(t, 1, len(users.Items))
	})
}

func TestUserManagementHandler_Delete(t *testing.T) {
	req := httptest.NewRequest(http.MethodDelete, "/api/users/12", nil)
	res := httptest.NewRecorder()

	ctx := e.NewContext(req, res)
	ctx.SetPath("/api/users/:id")
	ctx.SetParamNames("id")
	ctx.SetParamValues("12")

	ctrl := gomock.NewController(t)

	mockUserService := mockservice.NewMockUserService(ctrl)
	mockUserService.EXPECT().Delete(gomock.Any(), uint(12)).Return(nil).AnyTimes()

	userHandler := NewUserManagementHandler(handler, mockUserService)

	if assert.NoError(t, userHandler.Delete(ctx)) {
		assert.Equal(t, http.StatusNoContent, res.Code)
		assert.Equal(t, "", res.Body.String())
	}
}

func TestUserManagementHandler_UpdateSuccess(t *testing.T) {
	ctrl := gomock.NewController(t)

	userUpdateRequest := v1.UpdateUserRequest{
		Id:       1,
		FullName: "John Snow",
		Email:    "mail@johnsnow.techx",
		Role:     "editor",
		Password: "Password",
	}

	userJSON, _ := json.Marshal(userUpdateRequest)

	mockUserService := mockservice.NewMockUserService(ctrl)
	mockUserService.EXPECT().Update(gomock.Any(), &userUpdateRequest).Return(nil).Times(1)

	userHandler := NewUserManagementHandler(handler, mockUserService)

	req := httptest.NewRequest(http.MethodPut, "/api/users/1", bytes.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	if assert.NoError(t, userHandler.Update(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
	}

}

func TestUserManagementHandler_UpdateFailedDueToUserNotFound(t *testing.T) {

	ctrl := gomock.NewController(t)

	userUpdateRequest := v1.UpdateUserRequest{
		FullName: "John Snow",
		Email:    "mail@johnsnow.techx",
		Role:     "editor",
	}

	userJSON, _ := json.Marshal(userUpdateRequest)

	req := httptest.NewRequest(http.MethodPut, "/api/users/1", bytes.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)
	c.SetParamNames("id")
	c.SetParamValues("1")

	mockUserService := mockservice.NewMockUserService(ctrl)
	mockUserService.EXPECT().Update(gomock.Any(), gomock.Any()).Return(errors2.ErrUserDoesNotExists).AnyTimes()

	userHandler := NewUserManagementHandler(handler, mockUserService)

	err := userHandler.Update(c)
	if assert.Error(t, err) {
		assert.Equal(t, echo.ErrNotFound, err)
	}

}
