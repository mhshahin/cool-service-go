package model

import "github.com/golang-jwt/jwt/v5"

type JwtCustomClaims struct {
	Name  string
	Email string
	Role  string

	jwt.RegisteredClaims
}

type JwtRequest struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}
