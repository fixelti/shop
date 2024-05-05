package models

import "github.com/golang-jwt/jwt"

type JwtClaims struct {
	UserID uint `json:"user_id"`
	jwt.MapClaims
}
