package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type UserRequestHandler struct {
	*Handler
}

func NewUserRequestHandler(handler *Handler) *UserRequestHandler {
	return &UserRequestHandler{handler}
}

func (h *UserRequestHandler) List(ctx echo.Context) error {
	err := ctx.JSON(http.StatusOK, "ok")
	if err != nil {
		return err
	}
	return nil
}
