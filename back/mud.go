package main

import (
	"context"
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

	ctx := context.Background()
	agent.CreateClient(ctx)
	defer agent.DeleteClient(ctx)

	http.ListenAndServe(":"+*port, nil)
}
