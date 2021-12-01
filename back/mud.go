package main

import (
	"context"
	"flag"
	"net/http"
	"taeho/mud/agent"
	"taeho/mud/router"
)

func main() {
	port := flag.String("p", "8080", "Port")
	flag.Parse()

	ctx := context.Background()
	agent.CreateClient(ctx)
	defer agent.DeleteClient(ctx)

	http.ListenAndServe(":"+*port, router.GetRouters())
}
