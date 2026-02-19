package main

import (
	"advertiseproject/kitex_gen/advertiseproject/advertise"
	"advertiseproject/kitex_gen/advertiseproject/advertise/adservice"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/cloudwego/kitex/client"
	"github.com/gin-gonic/gin"
)

func main() {
	rpcHostPort := os.Getenv("AD_RPC_HOST_PORT")
	if rpcHostPort == "" {
		rpcHostPort = "127.0.0.1:8888"
	}

	cli, err := adservice.NewClient("AdService", client.WithHostPorts(rpcHostPort))
	if err != nil {
		log.Fatalf("创建 adservice 客户端错误: %v", err)
	}

	r := gin.Default()

	r.GET("/ad", func(c *gin.Context) {
		id := c.Query("id")
		strId, err := strconv.ParseInt(id, 10, 64)
		if err != nil {
			log.Printf("字段串id ：%s转int失败\n", id)
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "id必须为整数",
			})
			return
		}
		req := advertise.GetAdReq{
			Id: strId,
		}
		res, err := cli.GetAd(c.Request.Context(), &req)
		if err != nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": err.Error()})
			return
		}
		if res == nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "empty res"})
			return
		}
		c.JSON(http.StatusOK, res)

	})

	err = r.Run(":8080")
	if err != nil {
		log.Printf("gin 启动失败: %v", err)
		return
	}

}
