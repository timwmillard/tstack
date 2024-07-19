package handler

import (
	"app/admin"
	"app/auth"
	"app/internal/htmx"
	"app/model"
	"errors"
	"log/slog"
	"net/http"

	"github.com/alexedwards/scs/v2"
)

func adminRegisterForm(s model.Service, msg string) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {

		adminUsers, err := s.ListAdminUsers(req.Context())
		if err != nil {
			appError(wr, http.StatusInternalServerError, "Something went wrong", err)
			return
		}
		if len(adminUsers) > 1 {
			http.Redirect(wr, req, "/admin/", http.StatusSeeOther)
			return
		}

		content := admin.Register(msg)
		err = content.Render(req.Context(), wr)
		if err != nil {
			slog.Error("AdminHandler: adminLogin, render", "error", err)
			http.Error(wr, "Something when wrong", http.StatusInternalServerError)
			return
		}
	}
}

func adminRegister(s model.Service, sess *scs.SessionManager) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {
		name := req.FormValue("name")
		email := req.FormValue("email")
		password := req.FormValue("password")

		adminUsers, err := s.ListAdminUsers(req.Context())
		if err != nil {
			appError(wr, http.StatusInternalServerError, "Something went wrong", err)
			return
		}
		if len(adminUsers) > 1 {
			http.Redirect(wr, req, "/admin/", http.StatusSeeOther)
			return
		}

		_, err = s.Register(req.Context(), name, email, password)
		if err != nil {
			if errors.Is(err, model.ErrInvalidLogin) {
				wr.WriteHeader(http.StatusUnauthorized)
				adminLoginForm("invalid login")(wr, req)
				return
			}
			appError(wr, http.StatusInternalServerError, "Something went wrong", err)
			return
		}

		sess.Put(req.Context(), "username", email)

		http.Redirect(wr, req, "/admin/", http.StatusSeeOther)
	}
}

func adminLoginForm(msg string) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {

		content := admin.Login(msg)
		err := content.Render(req.Context(), wr)
		if err != nil {
			slog.Error("AdminHandler: adminLogin, render", "error", err)
			http.Error(wr, "Something when wrong", http.StatusInternalServerError)
			return
		}
	}
}

func adminLogin(s model.Service, sess *scs.SessionManager) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {
		username := req.FormValue("username")
		password := req.FormValue("password")

		user, err := s.Login(req.Context(), username, password)
		if err != nil {
			if errors.Is(err, model.ErrInvalidLogin) {
				wr.WriteHeader(http.StatusUnauthorized)
				adminLoginForm("invalid login")(wr, req)
				return
			}
			appError(wr, http.StatusInternalServerError, "Something went wrong", err)
			return
		}
		// First renew the session token...
		err = sess.RenewToken(req.Context())
		if err != nil {
			slog.Error("Renew token", "error", err)
			wr.WriteHeader(http.StatusInternalServerError)
			return
		}
		sess.Put(req.Context(), auth.SessionUserID, user.ID.String())

		http.Redirect(wr, req, "/admin/", http.StatusSeeOther)
	}
}

func adminLogout(sess *scs.SessionManager) http.HandlerFunc {
	return func(wr http.ResponseWriter, req *http.Request) {
		// First renew the session token...
		err := sess.RenewToken(req.Context())
		if err != nil {
			slog.Error("Renew token", "error", err)
			wr.WriteHeader(http.StatusInternalServerError)
			return
		}
		sess.Remove(req.Context(), auth.SessionUserID)

		if htmx.Request(req) {
			htmx.Redirect(wr, "/admin/login")
			return
		}
		http.Redirect(wr, req, "/admin/login", http.StatusSeeOther)
	}
}
