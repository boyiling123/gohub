package routers

import (
	"gohub/app/http/controllers/api/v1/auth"

	"github.com/gin-gonic/gin"
)

// 注册网页相关路由
func RegisterAPIRouters(r *gin.Engine) {

	// 测试一个 v1 的路由组，我们所有v1版本的路由都存放在这里
	v1 := r.Group("/v1")
	{
		authGroup := v1.Group("/auth")
		{
			suc := new(auth.SignupController)
			// 判断手机是否已注册
			authGroup.POST("/signup/phone/exist", suc.IsPhoneExist)
		}
	}

}
