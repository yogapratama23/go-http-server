package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/yogapratama23/go-http-server/internal/constants/message"
	"github.com/yogapratama23/go-http-server/internal/features/auth"
	"github.com/yogapratama23/go-http-server/internal/response"
)

type ExcludeRoutes struct {
	Url    string
	Method string
}

func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authRepo := new(auth.AuthRepository)
		found := false
		excludeRoutes := []ExcludeRoutes{
			{Url: "/", Method: "GET"},
			{Url: "/signin", Method: "POST"},
			{Url: "/signup", Method: "POST"},
		}

		route := ExcludeRoutes{
			Url:    r.URL.String(),
			Method: r.Method,
		}

		for i := 0; i < len(excludeRoutes); i++ {
			if (excludeRoutes[i].Url == route.Url) && (excludeRoutes[i].Method == route.Method) {
				found = true
				break
			}
		}

		if !found {
			authorization := r.Header.Get("Authorization")
			if authorization == "" {
				response.Error(w, message.Unauthorized, http.StatusForbidden)
				return
			}
			bearer := strings.Split(authorization, " ")
			if (len(bearer) < 2) || (bearer[1] == "") {
				response.Error(w, message.Unauthorized, http.StatusForbidden)
				return
			}

			token := bearer[1]
			newToken, err := authRepo.FindToken(token)
			if err != nil {
				response.Error(w, err.Error(), http.StatusForbidden)
				return
			}

			user, err := authRepo.FindById(&newToken.UserId)
			if err != nil {
				response.Error(w, err.Error(), http.StatusForbidden)
				return
			}

			newUser := auth.UserInfo{
				ID:       user.ID,
				Username: user.Username,
			}

			// get context data in controller user := r.Context().Value(middlewares.ContextKey("user"))
			ctxKey := auth.ContextKey("user")
			newCtx := context.WithValue(r.Context(), ctxKey, newUser)

			h.ServeHTTP(w, r.WithContext(newCtx))
			return
		}

		h.ServeHTTP(w, r)
	})
}
