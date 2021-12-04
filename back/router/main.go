package router

import (
	"net/http"
)

func GetRouters() *http.ServeMux {
	rootMux := http.NewServeMux()
	rootMux.Handle("/test/", http.StripPrefix("/test", testRouter()))
	return rootMux
}
