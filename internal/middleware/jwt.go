package middleware

import (
	"bphn/artikel-hukum/internal/constants"
	util "bphn/artikel-hukum/pkg/jwt"
	"github.com/golang-jwt/jwt/v5"
	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func JWTMiddleware(secret string) echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		Skipper: func(c echo.Context) bool {
			return c.Request().URL.Path == "auth/login"
		},
		ContextKey: constants.JwtCtxKey,
		SigningKey: []byte(secret),
		NewClaimsFunc: func(c echo.Context) jwt.Claims {
			return new(util.CustomClaims)
		},
	})
}
