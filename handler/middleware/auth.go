package middleware

import (
	"app/auth"
	"app/model"
	"log/slog"
	"net/http"

	"github.com/alexedwards/scs/v2"
	"github.com/gofrs/uuid/v5"
	"github.com/gorilla/mux"
)

func Auth(service model.Service, sess *scs.SessionManager) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(wr http.ResponseWriter, req *http.Request) {

			userID := sess.GetString(req.Context(), auth.SessionUserID)
			if userID == "" {
				slog.Debug("not logged in")
				http.Redirect(wr, req, "/admin/login", http.StatusSeeOther)
				return
			}
			user, err := service.GetUser(req.Context(), uuid.FromStringOrNil(userID))
			if err != nil {
				slog.Debug("user does not exist", "user_id", userID)
				http.Redirect(wr, req, "/admin/login", http.StatusSeeOther)
			}
			slog.Debug("Adding user to context", "user_id", user.ID)
			req = req.WithContext(auth.WithContextUser(req.Context(), user))
			next.ServeHTTP(wr, req)
		})
	}
}
