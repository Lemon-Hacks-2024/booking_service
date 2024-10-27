package pkg

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

type JWT struct {
	Token string
}

type tokenClaims struct {
	jwt.StandardClaims `json:"jwt_._standard_claims,omitempty"`
	UserID             int `json:"user_id" json:"user_id,omitempty"`
}

// Secret key for signing the token
var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

// CreateJWT создает JWT токен с заданными пользовательскими данными.
func CreateJWT(userID int) (JWT, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(7 * 24 * time.Hour).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userID,
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return JWT{}, err
	}

	return JWT{Token: tokenString}, nil

}

// ValidateJWT проверяет, является ли переданный JWT токен действительным.
func ValidateJWT(tokenString string) (int, error) {
	token, err := jwt.ParseWithClaims(tokenString, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}

		return secretKey, nil
	})
	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}

	return claims.UserID, nil
}
