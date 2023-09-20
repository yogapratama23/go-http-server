package middlewares

import (
	"log"
	"net/http"
	"time"
)

func Logging(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		method := r.Method
		uri := r.URL

		h.ServeHTTP(w, r)

		log.Println("INFO >>>", "latency:", time.Since(start), "method:", method, "endpoint:", uri)
	})
}
