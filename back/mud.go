package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"taeho/mud/agent"
	"taeho/mud/router"
)

func main() {
	// API_PORT는 frontend와 공유하는 환경변수
	// frontend에서는 backend API endpoint를 지정하는데 사용
	port := os.Getenv("API_PORT")
	if len(port) == 0 {
		port = "8080"
	}

	ctx := context.Background()
	agent.CreateClient(ctx)
	defer agent.DeleteClient(ctx)

	// 각 collection의 index들을 생성
	// DB의 scalability를 위해 DB 설정을 전부 AP에서 처리
	agent.CreateIndexes()

	fmt.Println("Backend listen to", port)
	http.ListenAndServe(":"+port, router.GetRouters())
}
