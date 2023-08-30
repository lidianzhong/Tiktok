package service

import (
	"errors"
	"log"
	"tiktok/dao"
	"tiktok/middleware"
	"tiktok/models"

	// "gorm.io/gorm"
)

func Register(req models.RegisterForm) (int64, string, error) {
    log.Println("调用了 service.register")
    
    // 首先检查用户是否已存在
    existingUser, err := dao.FindUserByName(req.UserName)
    // log.Println("打印一下这个错误",err)
    
    if existingUser != nil {
        return 0, "", errors.New("用户已存在")
    }
    log.Println("111")
    // 用户不存在，进行注册
    u := models.User{
        UserName: req.UserName,
        Password: req.Password,
    }

    err = dao.InsertUser(&u)
    if err != nil {
        return 0, "", errors.New("创建失败")
    }

    // 生成 token
    token, err := middleware.CreateToken(u.UserId, u.UserName, u.Password)
    if err != nil {
        return 0, "", err
    }

    // log.Println("调用了 service.register，成功返回用户 ID 和 token")
    // 在注册成功后，返回用户 ID 和 token
    return u.UserId, token, nil
}



func Login(form *models.LoginForm) (*models.User, error) {
	// 进行用户名密码验证
	user, err := dao.Login(form.UserName)
	if err != nil {
		return nil, err
	}
  // log.Println("这是service里面的user2：",user)
	token, err := middleware.CreateToken(user.UserId, user.UserName, user.Password)
	if err != nil {
		return nil, err
	}
	user.Token = token

	return user, nil
}





func GetUserInfo(userForm *models.UserForm) (*models.User, error) {
    userInfo, err := dao.FindUserById(userForm.UserId)
    if err != nil {
        return nil, err
    }
    return userInfo, nil
}

