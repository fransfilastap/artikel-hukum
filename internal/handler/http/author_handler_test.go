package http

import (
	v1 "bphn/artikel-hukum/api/v1"
	mockservice "bphn/artikel-hukum/internal/service/mocks"
	"bytes"
	"encoding/json"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthorManagementHandler_Register(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

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
	defer ctrl.Finish()

	registrationRequest := v1.AuthorRegistrationRequest{
		FullName:   "Frans Filasta Pratamax",
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
	authorService.EXPECT().Register(gomock.Any(), registrationRequest).Return(echo.ErrConflict).Times(1)

	handler := NewAuthorManagementHandler(handler, authorService)

	if assert.Error(t, handler.Register(c)) {
		assert.Equal(t, http.StatusConflict, rec.Code)
	}
}
