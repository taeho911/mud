package router

import (
	"io"
	"net/http"
	"taeho/mud/handler"
)

func testRouter() *http.ServeMux {
	testMux := http.NewServeMux()
	testMux.HandleFunc("/hello", helloFunc)
	testMux.HandleFunc("/session", handler.SessionTestHandler)
	return testMux
}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, world!\n")
}
