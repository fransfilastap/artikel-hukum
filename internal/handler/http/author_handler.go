package http

import (
	v1 "bphn/artikel-hukum/api/v1"
	"bphn/artikel-hukum/internal/service"
	"errors"
	"github.com/labstack/echo/v4"
	"net/http"
)

type AuthorManagementHandler struct {
	*Handler
	authorService service.AuthorService
}

func NewAuthorManagementHandler(handler *Handler, authorService service.AuthorService) *AuthorManagementHandler {
	return &AuthorManagementHandler{
		Handler:       handler,
		authorService: authorService,
	}
}

func (h *AuthorManagementHandler) Register(ctx echo.Context) error {

	var registrationRequest v1.AuthorRegistrationRequest

	if err := ctx.Bind(&registrationRequest); err != nil {
		h.Logger.Info(err.Error())
		return echo.NewHTTPError(http.StatusBadRequest, err)
	}

	if err := ctx.Validate(registrationRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, err)
	}

	if err := h.authorService.Register(ctx.Request().Context(), registrationRequest); err != nil {
		if errors.Is(err, v1.ErrEmailAlreadyExists) {
			return ctx.JSON(http.StatusConflict, v1.ErrEmailAlreadyExists)
		}
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return ctx.JSON(http.StatusCreated, &v1.CommonResponse{
		Code:    0,
		Message: "Success",
	})
}

func (h *AuthorManagementHandler) ForgotPassword(ctx echo.Context) error {
	// 1. Check user/author with submitted email exist in database
	// 2. if exists send reset password link
	// 3. if doesn't exist return err
	return nil
}

func (h *AuthorManagementHandler) GetProfile(ctx echo.Context) error {
	panic("implement me")
}
