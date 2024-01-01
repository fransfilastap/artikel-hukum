package http

import (
	"bphn/artikel-hukum/api/v1"
	mockservice "bphn/artikel-hukum/internal/service/mocks"
	"bytes"
	"encoding/json"
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

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
		defer ctrl.Finish()

		mockService := mockservice.NewMockUserService(ctrl)
		mockService.EXPECT().Create(c.Request().Context(), gomock.Any()).Return(nil).AnyTimes()

		userHandler := NewUserManagementHandler(handler, mockService)

		if assert.NoError(t, userHandler.Create(c)) {
			assert.Equal(t, http.StatusCreated, rec.Code)
		}

	})

	t.Run("Validation error while create user", func(t *testing.T) {
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
		defer ctrl.Finish()

		mockService := mockservice.NewMockUserService(ctrl)
		mockService.EXPECT().Create(gomock.Any(), userRequest).Return(errors.New("validation error")).AnyTimes()

		userHandler := NewUserManagementHandler(handler, mockService)

		if assert.NoError(t, userHandler.Create(c)) {
			assert.Equal(t, http.StatusBadRequest, rec.Code)
		}

	})

}

func TestUserRequestHandler_List(t *testing.T) {

	req := httptest.NewRequest(http.MethodGet, "/api/users", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockUserService := mockservice.NewMockUserService(ctrl)
	mockUserService.EXPECT().List(c.Request().Context()).Return([]v1.UserDataResponse{
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

	userHandler := NewUserManagementHandler(handler, mockUserService)

	err := userHandler.List(c)
	if err != nil {
		panic(err)
	}

	var users []v1.UserDataResponse
	unmarshalErr := json.Unmarshal(rec.Body.Bytes(), &users)

	if unmarshalErr != nil {
		panic(unmarshalErr)
	}

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, len(users), 3)
}

func TestUserManagementHandler_Update(t *testing.T) {
	userUpdateRequest := v1.UpdateUserRequest{
		FullName: "John Snow",
		Email:    "mail@johnsnow.techx",
		Role:     "editor",
	}

	userJSON, _ := json.Marshal(userUpdateRequest)

	req := httptest.NewRequest(http.MethodPut, "/api/users/1", bytes.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	res := httptest.NewRecorder()
	c := e.NewContext(req, res)

	ctrl := gomock.NewController(t)

	mockUserService := mockservice.NewMockUserService(ctrl)
	mockUserService.EXPECT().Update(gomock.Any(), gomock.Any()).Return(nil).AnyTimes()

	userHandler := NewUserManagementHandler(handler, mockUserService)
	err := userHandler.Update(c)

	if assert.NoError(t, err) {
		assert.Equal(t, res.Code, http.StatusOK)
	}

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
