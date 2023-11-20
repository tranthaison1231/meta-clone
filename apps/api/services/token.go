package services

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/pkg/errors"
)

var SecretKey []byte

type UserClaims struct {
	Email string `json:"email"`
	PwdTS int64  `json:"pwd_ts"`
	jwt.RegisteredClaims
}

func ParseToken(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, errors.New("Token is malformed")
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, errors.New("Token is expired")
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, errors.New("Token not active yet")
			} else {
				return nil, errors.New("Couldn't handle this token")
			}
		}
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("Couldn't handle this token")
}
