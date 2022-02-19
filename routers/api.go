package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterAPIRouters(r *gin.Engine) {

	// 测试一个 v1 的路由组，我们所有v1版本的路由都存放在这里
	v1 := r.Group("/v1")
	{
		// 注册一个路由
		v1.GET("/", func(c *gin.Context) {
			// 以 JSON 格式响应
			c.JSON(http.StatusOK, gin.H{
				"Hello": "v1版本",
			})
		})
	}

}
