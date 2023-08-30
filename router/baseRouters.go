package router

import (
	"tiktok/controller"
	"tiktok/middleware"

	"github.com/gin-gonic/gin"
)

func BaseRoutersInit(r *gin.Engine) {
	baseRouters := r.Group("/douyin")
	{
		baseRouters.GET("/feed/", controller.Feed)
		baseRouters.POST("/publish/action/", middleware.JWTMiddleWare(), controller.PublishVideoController)
		baseRouters.POST("/user/register/", controller.Register)
		baseRouters.POST("/user/login/", controller.Login)
		baseRouters.GET("/user/", middleware.JWTMiddleWare(), controller.UserInfo)

		//baseRouters.GET("/user", controller.UserInfo)
	}

}
