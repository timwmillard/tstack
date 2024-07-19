package handler

import (
	"log/slog"
	"net/http"
	// "runtime/debug"
)

var (
	appError   = errorResponse("App Handler")
	adminError = errorResponse("Admin Handler")
)

type errorResponseFunc func(http.ResponseWriter, int, string, error)

func errorResponse(name string) errorResponseFunc {
	return func(wr http.ResponseWriter, httpCode int, message string, err error) {
		slog.Error(name+": "+message, "error", err)
		// debug.PrintStack()
		http.Error(wr, message, httpCode)
	}
}
