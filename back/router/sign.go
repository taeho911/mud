package router

import (
	"net/http"
	"taeho/mud/handler"
)

func signRouter() *http.ServeMux {
	signMux := http.NewServeMux()
	signMux.HandleFunc("/up", handler.SignUpHandler)
	signMux.HandleFunc("/in", handler.SignInHandler)
	return signMux
}
