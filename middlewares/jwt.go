package middlewares

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/vispro-project/BackEnd-LombaKu.git/config"
	"github.com/vispro-project/BackEnd-LombaKu.git/helper"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				helper.ResponseError(w, http.StatusUnauthorized, "Unauthorized: No token found")
				return
			}
		}

		tokenString := c.Value
		claims := &config.JWTclaim{}

		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				helper.ResponseError(w, http.StatusUnauthorized, "Unauthorized: Invalid token signature")
				return
			case jwt.ValidationErrorExpired:
				helper.ResponseError(w, http.StatusUnauthorized, "Unauthorized: Token expired")
				return
			default:
				helper.ResponseError(w, http.StatusUnauthorized, "Unauthorized: Invalid token")
				return
			}
		}

		if !token.Valid {
			helper.ResponseError(w, http.StatusUnauthorized, "Unauthorized: Token is not valid")
			return
		}

		next.ServeHTTP(w, r)
	})
}
