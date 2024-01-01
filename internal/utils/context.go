package utils

import (
	jwt2 "bphn/artikel-hukum/pkg/jwt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

func GetUserIdFromCtx(c echo.Context) uint {
	user := c.Get("user").(*jwt.Token)
	claims := user.Claims.(*jwt2.CustomClaims)
	return claims.UserID
}
