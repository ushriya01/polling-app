package middleware

import (
	"context"
	"net/http"
	"poll-app/internal/utils"

	"github.com/julienschmidt/httprouter"
)

// ParseTokenMiddleware is a middleware function to parse and validate JWT token
func ParseTokenMiddleware(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		claims, err := utils.ParseToken(tokenString)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", claims["user_id"])
		next(w, r.WithContext(ctx), ps)
	}
}
