package routers

import (
	"tiktok/authmiddleware"
	"tiktok/controller"

	"github.com/gin-gonic/gin"
)

func BaseRoutersInit(r *gin.Engine) {
	baseRouters := r.Group("/douyin")
	{
		baseRouters.GET("/feed/", controller.Feed)
		baseRouters.POST("/publish/action/", authmiddleware.JWTMiddleWare(), controller.PublishVideoController)
		baseRouters.POST("/user/register/", controller.Register)
		baseRouters.POST("/user/login/", controller.Login)
    baseRouters.GET("/user/", authmiddleware.JWTMiddleWare(), controller.UserInfo)
   
		//baseRouters.GET("/user", controller.UserInfo)
	}

}
