package http

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Default(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Oh, Hey!")
}
