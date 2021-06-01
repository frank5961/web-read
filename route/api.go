package route

import (
	"github.com/gin-gonic/gin"
	"web-read/controller"
)

func LoadRoute(route *gin.Engine) *gin.Engine {
	rootGroup := route.Group("")
	{
		rootGroup.GET("/", func(context *gin.Context) {
			context.JSON(200, gin.H{
				"message": "hello",
			})
		})
	}

	// 微信
	apiGroup := route.Group("/wechat")
	{
		apiGroup.GET("/callback", controller.WechatController{}.Init)
		apiGroup.POST("/callback", controller.WechatController{}.Callback)
	}
	return route
}
