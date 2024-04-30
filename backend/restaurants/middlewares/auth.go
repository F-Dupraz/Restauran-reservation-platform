package middlewares

import (
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"

	"github.com/F-Dupraz/Restauran-reservation-platform.git/models"
	"github.com/F-Dupraz/Restauran-reservation-platform.git/server"
)

func CheckAuthMiddleware(s server.Server) func(h http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			tokenString := strings.TrimSpace(r.Header.Get("Authorization"))
			_, err := jwt.ParseWithClaims(tokenString, &models.UserToken{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(s.Config().JWTSecret), nil
			})
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
