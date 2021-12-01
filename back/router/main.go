package router

import (
	"io"
	"net/http"
)

func GetRouters() *http.ServeMux {
	rootMux := http.NewServeMux()
	testMux := http.NewServeMux()
	testMux.HandleFunc("/hello", helloFunc)
	rootMux.Handle("/test/", http.StripPrefix("/test", testMux))
	return rootMux
}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}
