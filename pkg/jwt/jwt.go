package jwt

import (
	"bphn/artikel-hukum/internal/model"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"regexp"
	"time"
)

type CustomClaims struct {
	UserID uint   `json:"user_id"`
	Role   string `json:"role"`
	jwt.RegisteredClaims
}

type JWT struct {
	key []byte
}

func NewJwt(conf *viper.Viper) *JWT {
	return &JWT{key: []byte(conf.GetString("security.jwt.key"))}
}

func (j *JWT) GenerateToken(userId uint, role model.Role, expiresAt time.Time) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, CustomClaims{
		UserID: userId,
		Role:   string(role),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "",
			Subject:   "",
			ID:        "",
			Audience:  []string{},
		},
	})

	tokenStr, err := token.SignedString(j.key)

	if err != nil {
		return "", err
	}

	return tokenStr, err

}

func (j *JWT) ParseToken(tokenStr string) (*CustomClaims, error) {
	re := regexp.MustCompile(`(?i)Bearer`)
	tokenStr = re.ReplaceAllString(tokenStr, "")

	if tokenStr == "" {
		return nil, errors.New("token is empty")
	}

	token, err := jwt.ParseWithClaims(tokenStr, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.key, nil
	})

	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	} else {
		return nil, err
	}
}
