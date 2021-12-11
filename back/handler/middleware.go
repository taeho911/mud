package handler

import (
	"context"
	"net/http"
	"taeho/mud/errors"
)

type ctxKey string

const (
	usernameKey ctxKey = "username"
)

func Auth(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, err := getSession(r)
		if err != nil {
			deleteSession(w, r)
			writeError(w, errors.UNAUTHORIZED, err.Error(), http.StatusUnauthorized)
			return
		}
		if err := updateSession(r); err != nil {
			deleteSession(w, r)
			writeError(w, errors.UNAUTHORIZED, err.Error(), http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r.WithContext(context.WithValue(r.Context(), usernameKey, session.username)))
	})
}
