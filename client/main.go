package main

import (
	"advertiseproject/kitex_gen/advertiseproject/advertise"
	"advertiseproject/kitex_gen/advertiseproject/advertise/adservice"
	"context"
	"fmt"

	"github.com/cloudwego/kitex/client"
)

func main() {
	cli, err := adservice.NewClient("AdService", client.WithHostPorts("127.0.0.1:8888"))
	if err != nil {
		fmt.Println("创建 adservice 客户端错误:", err)
		return
	}
	req := advertise.GetAdReq{
		Id: 1,
	}
	ctx := context.Background()
	res, err := cli.GetAd(ctx, &req)
	if err != nil {
		fmt.Println("调用 GetAd 失败:", err)
		return
	}
	fmt.Println(res)
}
