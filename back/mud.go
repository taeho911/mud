package main

import (
	"flag"
	"io"
	"net/http"
	"taeho/mud/agent"
)

func main() {
	port := flag.String("p", "80", "Port")
	flag.Parse()

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}
	http.HandleFunc("/hello", helloHandler)

	agent.CreateClient()
	http.ListenAndServe(":"+*port, nil)
}
