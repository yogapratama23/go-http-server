package middlewares

import (
	"net/http"
	"strings"

	"github.com/yogapratama23/go-http-server/internal/constants/message"
	"github.com/yogapratama23/go-http-server/internal/response"
)

type ExcludeRoutes struct {
	Url    string
	Method string
}

func Auth(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		found := false
		excludeRoutes := []ExcludeRoutes{
			{Url: "/", Method: "GET"},
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

			// logic of verifying user here
		}

		h.ServeHTTP(w, r)
	})
}
