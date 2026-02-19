package main

import (
	advertise "advertiseproject/kitex_gen/advertiseproject/advertise/adservice"
	"log"
	"time"

	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/server"
)

func main() {
	maxQPS := envInt("AD_MAX_QPS", 50)
	maxConn := envInt("AD_MAX_CONN", 1000)
	breakerThreshold := envInt("AD_BREAKER_FAIL_THRESHOLD", 3)
	breakerOpenSeconds := envInt("AD_BREAKER_OPEN_SECONDS", 10)

	cb := newServerCircuitBreaker(breakerThreshold, time.Duration(breakerOpenSeconds)*time.Second)

	svr := advertise.NewServer(
		new(AdServiceImpl),
		server.WithLimit(&limit.Option{
			MaxConnections: maxConn,
			MaxQPS:         maxQPS,
		}),
		server.WithMiddleware(cb.middleware),
	)

	err := svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
