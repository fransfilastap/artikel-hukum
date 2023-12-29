package http

import (
	"bphn/artikel-hukum/api"
	"bphn/artikel-hukum/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserDataRequestHandler struct {
	*Handler
	userService service.UserService
}

func NewUserRequestHandler(handler *Handler, userService service.UserService) *UserDataRequestHandler {
	return &UserDataRequestHandler{handler, userService}
}

func (h *UserDataRequestHandler) List(ctx echo.Context) error {
	users, err := h.userService.List(ctx.Request().Context())

	if err != nil {
		return err
	}

	err = ctx.JSON(http.StatusOK, users)
	if err != nil {
		return err
	}
	return nil
}

func (h *UserDataRequestHandler) Create(ctx echo.Context) error {
	var createUserRequest api.AdminCreateUserRequest

	if err := ctx.Bind(&createUserRequest); err != nil {
		h.Logger.Debug(err.Error())

		return err
	}

	if err := h.userService.Create(ctx.Request().Context(), &createUserRequest); err != nil {
		// TODO Error Struct
		return ctx.JSON(http.StatusInternalServerError, "Oppss")
	}

	return nil
}
