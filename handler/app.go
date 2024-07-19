package handler

import (
	"app/app"
	"app/model"
	"log/slog"
	"net/http"
)

func index(_ model.Service) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {

		content := app.Index()
		err := app.Base("Home", content).Render(req.Context(), wr)
		if err != nil {
			slog.Error("App Handler: index, templ.Base", "error", err)
			http.Error(wr, "Something when wrong", http.StatusInternalServerError)
			return
		}
	}
}
