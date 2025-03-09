package middlewares

import (
	"context"
	"log"
	"net/http"
	"strings"

	"github.com/tejashwinn/splitwise/constants"
	"github.com/tejashwinn/splitwise/util"
)

type AuthMiddleware struct {
	JwtUtil util.JwtUtil
}

func NewAuthMiddleware(
	jwtUtil *util.JwtUtil,
) *AuthMiddleware {
	return &AuthMiddleware{JwtUtil: *jwtUtil}
}

func (a *AuthMiddleware) AuthenticateAndSetUserId(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}

		parts := strings.Split(authHeader, " ")
		if len(parts) != 2 || parts[0] != "Bearer" {
			http.Error(w, "invalid token format", http.StatusUnauthorized)
			return
		}

		userId, err := a.JwtUtil.VerifyToken(parts[1])

		if err != nil {
			log.Println(err)
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}
		ctx := context.WithValue(r.Context(), constants.UserId, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
