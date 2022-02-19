package main

import (
	"fmt"
	"gohub/bootstrap"

	"github.com/gin-gonic/gin"
)

func main() {
	// new 一个 Gin Engine 实例
	r := gin.New()

	// 初始化路由的绑定
	bootstrap.SetupRoute(r)

	// 运行服务，默认为 8080，我们指定宽口为 8000
	err := r.Run(":8000")
	if err != nil {
		// 错误处理，端口被占用了或其他错误
		fmt.Println(err.Error())
	}

}
