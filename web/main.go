package main

import (
	"advertiseproject/kitex_gen/advertiseproject/advertise"
	"advertiseproject/kitex_gen/advertiseproject/advertise/adservice"
	"log"
	"net/http"
	"strconv"

	"github.com/cloudwego/kitex/client"
	"github.com/gin-gonic/gin"
)

func main() {
	cli, err := adservice.NewClient("AdService", client.WithHostPorts("47.98.182.220:8888"))
	if err != nil {
		log.Fatalln("创建 adservice 客户端错误")
	}

	r := gin.Default()

	r.GET("/ad", func(c *gin.Context) {
		id := c.Query("id")
		strId, err := strconv.ParseInt(id, 0, 64)
		if err != nil {
			log.Printf("字段串id ：%s转int失败\n", id)
			c.JSON(http.StatusBadGateway, gin.H{
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
		}
		if res == nil {
			c.JSON(http.StatusBadGateway, gin.H{"error": "empty res"})
		}
		c.JSON(200, res)

	})

	err = r.Run(":8080")
	if err != nil {
		log.Printf("gin 启动失败: %v", err)
		return
	}

}
