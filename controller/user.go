package controller

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"tiktok/dao"
	"tiktok/model"
	"tiktok/service"
	"tiktok/util"
)

// usersLoginInfo use map to store user info, and key is username+password for demo
// user data will be cleared every time the server starts
// test data: username=zhanglei, password=douyin
//
//	var usersLoginInfo = map[string]User{
//		"zhangleidouyin": {
//			Id:            1,
//			Name:          "zhanglei",
//			FollowCount:   10,
//			FollowerCount: 5,
//			IsFollow:      true,
//		},
//	}
var userIdSequence = int64(1)

type UserLoginResponse struct {
	util.Response
	UserId int64  `json:"user_id,omitempty"`
	Token  string `json:"token"`
}

type UserResponse struct {
	util.Response
	User model.User `json:"user"`
}

func Register(c *gin.Context) {
	var req model.RegisterForm

	// 从 URL 查询参数中获取用户名和密码
	username := c.Query("username")
	password := c.Query("password")

	// 将查询参数绑定到请求结构体
	req.UserName = username
	req.Password = password

	userId, token, err := service.Register(req)
	if err != nil {
		log.Println("注册失败", err)
		c.JSON(http.StatusOK, UserLoginResponse{
			Response: util.Response{StatusCode: 2, StatusMsg: "用户已存在"},
		})
		return
	}

	c.JSON(http.StatusOK, UserLoginResponse{
		Response: util.Response{StatusCode: 0},
		UserId:   userId,
		Token:    token,
	})
}

func Login(c *gin.Context) {
	log.Println("Login request received")

	u := &model.LoginForm{}

	log.Println("URL:", c.Request.URL.String())
	log.Println("Params:", c.Request.URL.Query())

	// 从 URL 参数中获取用户名和密码
	username := c.Query("username")
	password := c.Query("password")

	// 将用户名和密码绑定到登录表单
	u.UserName = username
	u.Password = password

	// 在这里可以进行参数验证逻辑
	if username == "" || password == "" {
		c.JSON(http.StatusOK, gin.H{
			"status_code": 400,
			"status_msg":  "用户名和密码不能为空",
		})
		return
	}

	user, err := service.Login(u)
	if err != nil {
		log.Println("service.Login failed", err)
		if err.Error() == dao.ErrorUserNotExit {
			c.JSON(http.StatusOK, gin.H{
				"status_code": 3,
				"status_msg":  "用户不存在",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status_code": 1,
			"status_msg":  "请求参数错误",
		})
		return
	}

	log.Println("User found:", user)

	c.JSON(http.StatusOK, gin.H{
		"status_code": 0,
		"user_id":     user.UserId,
		"token":       user.Token,
		"status_msg":  "登录成功",
	})
}

func UserInfo(c *gin.Context) {
	// var u model.UserForm

	log.Println("URL 参数 user_id:", c.Query("user_id"))
	log.Println("URL 参数 token:", c.Query("token"))

	// 手动解析参数
	token := c.Query("token")

	// 验证 token，您可以在此处添加验证逻辑
	claims, err := util.ParseToken(token)
	if err != nil {
		log.Println("Token 验证失败", err)
		c.JSON(http.StatusOK, UserResponse{
			Response: util.Response{StatusCode: 5, StatusMsg: "Token 验证失败"},
		})
		return
	}

	// 使用解析后的 claims 数据获取用户信息
	userInfo, err := service.GetUserInfo(&model.UserForm{
		UserId: claims.UserId, // 使用解析后的用户 ID
		Token:  token,
	})
	if err != nil {
		log.Println("获取用户信息失败", err)
		c.JSON(http.StatusOK, UserResponse{
			Response: util.Response{StatusCode: 3, StatusMsg: "用户不存在"},
		})
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		Response: util.Response{StatusCode: 0},
		User:     *userInfo,
	})
}
