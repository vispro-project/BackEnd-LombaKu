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
				response := map[string]string{"message": "unauthorized"}
				helper.ResponseJson(w, http.StatusUnauthorized, response)
				return
			}
		}
		tokeString := c.Value

		claims := &config.JWTclaim{}

		token, err := jwt.ParseWithClaims(tokeString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil {
			v, _ := err.(*jwt.ValidationError)
			switch v.Errors {
			case jwt.ValidationErrorSignatureInvalid:
				response := map[string]string{"message": "unauthorized"}
				helper.ResponseJson(w, http.StatusUnauthorized, response)
				return
			case jwt.ValidationErrorExpired:
				response := map[string]string{"message": "Token Expired"}
				helper.ResponseJson(w, http.StatusUnauthorized, response)
				return
			default:
				response := map[string]string{"message": "unauthorized"}
				helper.ResponseJson(w, http.StatusUnauthorized, response)
				return

			}
		}

		if !token.Valid {
			response := map[string]string{"message": "unauthorized"}
			helper.ResponseJson(w, http.StatusUnauthorized, response)
			return
		}
		next.ServeHTTP(w, r)
	})
}
