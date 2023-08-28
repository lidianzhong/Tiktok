package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"tiktok/dao"
	"tiktok/models"
	"tiktok/service"
)

var userIdSequence = int64(1)

type UserLoginResponse struct {
	Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	Response
	User models.User `json:"user"`
}

func Register(c *gin.Context) {
	var req models.RegisterForm
	if err := c.ShouldBind(&req); err != nil {
		log.Println("注册时请求参数错误", err)
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "请求参数错误"},
			})
			return
		}
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "请求参数错误"},
		})
		return
	}

	_ = service.Register(req)

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   userIdSequence,
		Token:    req.UserName + req.Password,
	})
}

func Login(c *gin.Context) {
	var u *models.LoginForm
	if err := c.ShouldBind(&u); err != nil {
		log.Println("登陆时请求参数错误", err)
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "请求参数错误"},
			})
			return
		}
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "请求参数错误"},
		})
		return
	}

	user, err := service.Login(u)
	if err != nil {
		log.Println("登录失败", err)
		if err.Error() == dao.ErrorUserNotExit {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 3, StatusMsg: "用户不存在"},
			})
			return
		}
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "请求参数错误"},
		})
		return
	}

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: Response{StatusCode: 0},
		UserId:   user.UserId,
		Token:    user.Token,
	})
}

func UserInfo(c *gin.Context) {
	var u *models.UserForm
	if err := c.ShouldBind(&u); err != nil {
		log.Println("获取信息时请求参数错误", err)
		_, ok := err.(validator.ValidationErrors)
		if !ok {
			c.JSON(http.StatusOK, UserLoginResponse{
				Response: Response{StatusCode: 1, StatusMsg: "请求参数错误"},
			})
			return
		}
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 1, StatusMsg: "请求参数错误"},
		})
		return
	}

	userInfo, err := service.GetUserInfo(u)
	if err != nil {
		log.Println("获取用户信息失败", err)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: Response{StatusCode: 3, StatusMsg: "用户不存在"},
		})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		Response: Response{StatusCode: 0},
		User:     *userInfo,
	})
}
