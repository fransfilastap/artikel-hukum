package http

import (
	v1 "bphn/artikel-hukum/api/v1"
	"bphn/artikel-hukum/internal/errors"
	"bphn/artikel-hukum/internal/ito"
	mockservice "bphn/artikel-hukum/internal/service/mocks"
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

func TestAuthorManagementHandler_Register(t *testing.T) {
	ctrl := gomock.NewController(t)

	registrationRequest := v1.AuthorRegistrationRequest{
		FullName:   "Frans Filasta Pratama",
		Email:      "mail@fransfp.dev",
		Occupation: "PNS",
		Company:    "BPHN",
		Password:   "12345678",
	}
	registrationJSON, _ := json.Marshal(registrationRequest)
	req := httptest.NewRequest(http.MethodPost, "/api/author", bytes.NewReader(registrationJSON))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	authorService := mockservice.NewMockAuthorService(ctrl)
	authorService.EXPECT().Register(gomock.Any(), registrationRequest).Return(nil).Times(1)

	handler := NewAuthorManagementHandler(handler, authorService)

	if assert.NoError(t, handler.Register(c)) {
		assert.Equal(t, http.StatusCreated, rec.Code)
	}

}

func TestAuthorManagementHandler_RegisterFailed(t *testing.T) {
	ctrl := gomock.NewController(t)

	registrationRequest := v1.AuthorRegistrationRequest{
		FullName:   "John Doe",
		Email:      "mail@johndoe.com",
		Occupation: "Lawyer",
		Company:    "Github",
		Password:   "12345678",
	}
	registrationJSON, _ := json.Marshal(registrationRequest)
	req := httptest.NewRequest(http.MethodPost, "/api/author", bytes.NewReader(registrationJSON))
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	authorService := mockservice.NewMockAuthorService(ctrl)
	authorService.EXPECT().Register(gomock.Any(), registrationRequest).Return(errors.ErrEmailAlreadyExists).AnyTimes()

	handler := NewAuthorManagementHandler(handler, authorService)

	if assert.NoError(t, handler.Register(c)) {
		assert.Equal(t, http.StatusConflict, rec.Code)
	}
}

func TestAuthorManagementHandler_GetProfile(t *testing.T) {
	ctrl := gomock.NewController(t)

	want := v1.AuthorProfileDataResponse{
		Id:         12,
		FullName:   "John Doe",
		Email:      "mail@johndoe.com",
		Occupation: "Lawyer",
		Company:    "Github",
	}

	resJSON, _ := json.Marshal(want)

	authorService := mockservice.NewMockAuthorService(ctrl)
	authorService.EXPECT().Profile(gomock.Any(), uint(12)).Return(want, nil).AnyTimes()

	handler := NewAuthorManagementHandler(handler, authorService)

	e.GET("/api/profile", handler.GetProfile)

	req := httptest.NewRequest(http.MethodGet, "/api/profile", nil)
	req.Header.Set("Authorization", "Bearer "+generateToken(t))
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)
	assert.Equal(t, string(resJSON), strings.Replace(rec.Body.String(), "\n", "", 1))
}

func TestAuthorManagementHandler_UpdateProfile(t *testing.T) {
	ctrl := gomock.NewController(t)

	updateRequest := v1.UpdateAuthorProfileRequest{
		Id:         12,
		FullName:   "John Doe",
		Email:      "mail@johndoe.com",
		Occupation: "Lawyer",
		Company:    "Microsoft",
	}

	updatePayload, _ := json.Marshal(updateRequest)

	authorService := mockservice.NewMockAuthorService(ctrl)
	authorService.EXPECT().UpdateProfile(gomock.Any(), updateRequest).Return(nil).AnyTimes()

	handler := NewAuthorManagementHandler(handler, authorService)

	e.PUT("/api/author", handler.UpdateProfile)

	req := httptest.NewRequest(http.MethodPut, "/api/author", bytes.NewReader(updatePayload))
	req.Header.Set("Authorization", "Bearer "+generateToken(t))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	rec := httptest.NewRecorder()

	e.ServeHTTP(rec, req)

	assert.Equal(t, http.StatusOK, rec.Code)

}

func generateToken(t *testing.T) string {
	token, err := jwt.GenerateToken(12, "author", time.Now().Add(time.Hour*24))
	if err != nil {
		t.Error(err)
		return token
	}

	return token
}

func TestAuthorManagementHandler_List(t *testing.T) {

	ctrl := gomock.NewController(t)

	req := httptest.NewRequest(http.MethodGet, "/api/authors?page=1&size=2", nil)
	req.Header.Set(echo.HeaderAuthorization, "Bearer "+generateToken(t))
	rec := httptest.NewRecorder()

	authorService := mockservice.NewMockAuthorService(ctrl)
	authorService.EXPECT().List(gomock.Any(), gomock.Any()).Return(ito.ListQueryResult[v1.AuthorProfileDataResponse]{
		TotalPage: 2,
		Page:      1,
		Items: []v1.AuthorProfileDataResponse{{
			Id:         12,
			FullName:   "John Doe",
			Email:      "mail@johndoe.com",
			Occupation: "Lawyer",
			Company:    "Microsoft",
		}},
	}, nil)

	handler := NewAuthorManagementHandler(handler, authorService)

	e.GET("/api/authors", handler.List)
	e.ServeHTTP(rec, req)

	var users ito.ListQueryResult[ito.UserDataResponse]
	unmarshalErr := json.Unmarshal(rec.Body.Bytes(), &users)

	if unmarshalErr != nil {
		t.Fail()
	}
	assert.Equal(t, http.StatusOK, rec.Code)

}
