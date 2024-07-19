package middleware

import "net/http"

func JSON(next http.Handler) http.Handler {
	return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {

		wr.Header().Add("Content-Type", "application/json")

		next.ServeHTTP(wr, req)
	})
}
