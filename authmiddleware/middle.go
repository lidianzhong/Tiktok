package authmiddleware



import (

	"net/http"

	"tiktok/controller"

	"tiktok/middleware"

	"time"

  "log"

	"github.com/gin-gonic/gin"

)



func JWTMiddleWare() gin.HandlerFunc {

	return func(c *gin.Context) {

		tokenStr := c.Query("token")

		if tokenStr == "" {

			tokenStr = c.PostForm("token")

		}

		//用户不存在

		if tokenStr == "" {

			c.JSON(http.StatusOK, controller.Response{StatusCode: 401, StatusMsg: "该用户不存在"})

			c.Abort() //阻止执行

			return

		}

		//验证token

		tokenStruck, err := middleware.ParseToken(tokenStr)

		if err != nil {
       log.Println("Token 解析错误:", err)

			c.JSON(http.StatusOK, controller.Response{

				StatusCode: 403,

				StatusMsg:  "token不正确",

			})

			c.Abort() //阻止执行

			return

		}
    // log.Println("解析后的 Token:", tokenStruck)

		//token超时

		if tokenStruck.ExpiresAt.Time.Before(time.Now()) {

			c.JSON(http.StatusOK, controller.Response{

				StatusCode: 402,

				StatusMsg:  "token过期",

			})

			c.Abort() //阻止执行

			return

		}

		c.Set("user_id", tokenStruck.UserId)

		c.Next()

	}

}

