package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"runtime"
	"taeho/mud/agent"
	"taeho/mud/handler"
	"taeho/mud/router"
	"time"
)

const (
	DEFAULT_PORT    string        = "9011"
	SES_GC_INTERVAL time.Duration = 3 * time.Minute
	SES_GC_PERCENT  float64       = float64(50)
)

func main() {
	// API_PORT는 frontend와 공유하는 환경변수
	// frontend에서는 backend API endpoint를 지정하는데 사용
	port := os.Getenv("API_PORT")
	if len(port) == 0 {
		port = DEFAULT_PORT
	}

	ctx := context.Background()
	agent.CreateClient(ctx)
	defer agent.DeleteClient(ctx)

	// 각 collection의 index들을 생성
	// DB의 scalability를 위해 DB 설정을 전부 AP에서 처리
	agent.CreateIndexes()

	// 만료 세션을 주기적으로 청소하는 GC 가동
	go sessionGC()

	fmt.Println("MUD api server is listening to", port)
	http.ListenAndServe(":"+port, router.GetRouters())
}

func sessionGC() {
	for {
		var memstat runtime.MemStats
		runtime.ReadMemStats(&memstat)

		if float64(memstat.Alloc)/float64(memstat.Sys)*100 > SES_GC_PERCENT {
			handler.SessionManager.GC()
		}

		time.Sleep(SES_GC_INTERVAL)
	}
}
