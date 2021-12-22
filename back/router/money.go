package router

import (
	"net/http"
	"taeho/mud/handler"
)

func moneyRouter() *http.ServeMux {
	moneyMux := http.NewServeMux()
	moneyMux.HandleFunc("/get", handler.Auth(handler.MoneyGetAll))
	moneyMux.HandleFunc("/post", handler.Auth(handler.MoneyPostHandler))
	moneyMux.HandleFunc("/delete", handler.Auth(handler.MoneyDeleteHandler))
	return moneyMux
}
