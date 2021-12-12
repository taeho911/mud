package router

import (
	"net/http"
	"taeho/mud/handler"
)

func signRouter() *http.ServeMux {
	signMux := http.NewServeMux()
	signMux.HandleFunc("/up", handler.SignUpHandler)
	signMux.HandleFunc("/in", handler.SignInHandler)
	signMux.HandleFunc("/confirm", handler.SignConfirmHandler)
	signMux.HandleFunc("/out", handler.Auth(handler.SignOutHandler))
	signMux.HandleFunc("/delete", handler.Auth(handler.SignDeleteHandler))
	return signMux
}
