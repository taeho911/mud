package router

import (
	"net/http"
)

func GetRouters() *http.ServeMux {
	rootMux := http.NewServeMux()
	rootMux.Handle("/api/", http.StripPrefix("/api", rootMux))
	rootMux.Handle("/sign/", http.StripPrefix("/sign", signRouter()))
	rootMux.Handle("/money/", http.StripPrefix("/money", moneyRouter()))
	rootMux.Handle("/test/", http.StripPrefix("/test", testRouter()))
	return rootMux
}
