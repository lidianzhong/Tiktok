package routers

import (
	"github.com/gin-gonic/gin"
	"tiktok/controller"
	"tiktok/middleware"
)

func BaseRoutersInit(r *gin.Engine) {
	baseRouters := r.Group("/douyin")
	{
		baseRouters.GET("/feed", controller.Feed)
		baseRouters.POST("/publish/action/", middleware.JWTMiddleWare(), controller.PublishVideoController)
		baseRouters.POST("/user/register", controller.Register)
		baseRouters.POST("/user/login", controller.Login)
		baseRouters.GET("/user", middleware.JWTMiddleWare(), controller.UserInfo)
	}

}
