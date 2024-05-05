package models

import "github.com/golang-jwt/jwt"

//TODO: на будущее

type JwtClaims struct {
	UserID uint `json:"user_id"`
	jwt.MapClaims
}
