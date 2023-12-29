package http

import (
	"bphn/artikel-hukum/internal/service"
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserRequestHandler struct {
	*Handler
	userService service.UserService
}

func NewUserRequestHandler(handler *Handler, userService service.UserService) *UserRequestHandler {
	return &UserRequestHandler{handler, userService}
}

func (h *UserRequestHandler) List(ctx echo.Context) error {
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

func (h *UserRequestHandler) Create(ctx echo.Context) error {
	if err := h.userService.Create(ctx.Request().Context()); err != nil {
		// TODO Error Struct
		return ctx.JSON(http.StatusInternalServerError, "Oppss")
	}

	return nil
}
