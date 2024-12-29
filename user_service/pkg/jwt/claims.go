package jwt

import "github.com/golang-jwt/jwt/v5"

type UserClaims struct {
	jwt.RegisteredClaims
	UserID   uint
	Role     uint
	Sections []string
}
