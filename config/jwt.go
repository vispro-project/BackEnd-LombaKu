package config

import "github.com/golang-jwt/jwt/v4"

var JWT_KEY = []byte("as3dadad34jnri3ajd9834353")

type JWTclaim struct {
	Username string `json:"username"`
	UserId   int64  `json:"userId"`
	jwt.RegisteredClaims
}
