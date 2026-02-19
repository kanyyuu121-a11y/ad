package main

import (
	"advertiseproject/kitex_gen/advertiseproject/advertise"
	"advertiseproject/kitex_gen/advertiseproject/advertise/adservice"
	"context"
	"fmt"
	"os"
	"sync"
	"time"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/fallback"
)

func main() {

	cbs := circuitbreak.NewCBSuite(circuitbreak.RPCInfo2Key)

	cli, err := adservice.NewClient(
		"AdService",
		client.WithHostPorts("121.196.231.74:8888"),
		client.WithRPCTimeout(500*time.Millisecond),
		client.WithCircuitBreaker(cbs),
		client.WithFallback(fallback.ErrorFallback(fallback.UnwrapHelper(func(ctx context.Context, req, resp interface{}, err error) (interface{}, error) {
			r, _ := req.(*advertise.GetAdReq)
			return &advertise.GetAdRes{
				Ad: &advertise.Advertise{
					Id:          r.Id,
					Name:        "client-fallback-ad",
					Description: "fallback due to timeout/circuit-break",
					Stock:       0,
				},
			}, nil
		}))),
	)
	if err != nil {
		fmt.Println("创建 adservice 客户端错误:", err)
		return
	}

	scenario := os.Getenv("AD_SCENARIO")
	if scenario == "" {
		scenario = "breaker"
	}

	switch scenario {
	case "normal":
		callOnce(cli, &advertise.GetAdReq{Id: 10})
	case "timeout":
		callOnce(cli, &advertise.GetAdReq{Id: 999, Name: "slow"})
	case "breaker":
		for i := 0; i < 5; i++ {
			callOnce(cli, &advertise.GetAdReq{Id: -1})
		}
		callOnce(cli, &advertise.GetAdReq{Id: 2})
	case "limit":
		runBurst(cli, 100)
	default:
		fmt.Println("未知场景，支持: normal | timeout | breaker | limit")
	}
}

func callOnce(cli adservice.Client, req *advertise.GetAdReq) {
	ctx := context.Background()
	res, err := cli.GetAd(ctx, req)
	if err != nil {
		fmt.Printf("req=%+v err=%v\n", req, err)
		return
	}
	fmt.Printf("req=%+v res=%+v\n", req, res)
}

func runBurst(cli adservice.Client, n int) {
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			defer wg.Done()
			callOnce(cli, &advertise.GetAdReq{Id: int64(1000 + i)})
		}(i)
	}
	wg.Wait()
}
