package handler

import (
	"app/admin"
	"app/model"
	"net/http"

	"log/slog"
)

func adminIndex(s model.Service) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {

		content := admin.Index()
		err := admin.Dashboard("Dashboard", content).Render(req.Context(), wr)
		if err != nil {
			slog.Error("AdminHandler: adminIndex, render", "error", err)
			http.Error(wr, "Something when wrong", http.StatusInternalServerError)
			return
		}
	}
}
