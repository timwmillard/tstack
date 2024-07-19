package handler

import (
	"app/admin"
	"context"
	"fmt"
	"log/slog"
	"net/http"
	"runtime"
)

func adminNotFound() http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {
		httpCode := http.StatusNotFound
		message := "Sorry, we couldn't find your page."

		wr.WriteHeader(httpCode)

		content := admin.ErrorPage(httpCode, http.StatusText(httpCode), message)
		err := content.Render(context.Background(), wr)
		if err != nil {
			slog.Error("Admin Handler: NotFound, ErrorPage", "error", err)
			http.Error(wr, "Something when wrong", http.StatusInternalServerError)
			return
		}
	}
}

type adminErrorResponseFunc func(http.ResponseWriter, string, error)

var (
	adminErrorNotFound    = adminErrorResponse(http.StatusNotFound)
	adminErrorBadRequest  = adminErrorResponse(http.StatusBadRequest)
	adminErrorServerError = adminErrorResponse(http.StatusInternalServerError)
)

func adminErrorResponse(httpCode int) adminErrorResponseFunc {
	return func(wr http.ResponseWriter, message string, err error) {
		slog.Error("Admin: "+message, "error", err)
		// debug.PrintStack()
		content := admin.ErrorPage(httpCode, http.StatusText(httpCode), message)
		err = content.Render(context.Background(), wr)
		if err != nil {
			slog.Error("Admin Handler: adminErrorResonse render", "error", err)
			http.Error(wr, "Something when wrong", http.StatusInternalServerError)
			return
		}
		wr.WriteHeader(httpCode)
	}
}

func adminErrorRender(wr http.ResponseWriter, err error) {
	slog.Error("Admin Handler: adminIndex, render", "error", err)
	http.Error(wr, "Something when wrong", http.StatusInternalServerError)
}

func trace() {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	fmt.Printf("%s:%d %s\n", file, line, f.Name())
}
