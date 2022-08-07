package middleware

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/martinjirku/zasobar/config"
)

type userKey int

const (
	UserKey userKey = iota
)

func JwtMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenCookie, err := r.Cookie("auth")
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		token := tokenCookie.Value
		jwtToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.GetJwtSecret()), nil
		})
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		if !jwtToken.Valid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserKey, jwtToken)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
