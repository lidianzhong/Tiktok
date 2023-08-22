package service

import (
	"errors"
	"gorm.io/gorm"
	"tiktok/dao"
	"tiktok/middleware"
	"tiktok/models"
)

func Register(req models.RegisterForm) (error error) {
	_, err := dao.FindUserByName(req.UserName)
	if gorm.ErrRecordNotFound != err {
		return errors.New("用户已存在")
	}

	u := models.User{
		UserName: req.UserName,
		Password: req.Password,
	}

	err = dao.InsertUser(&u)
	if err != nil {
		return errors.New("创建失败")
	}
	return err
}

func Login(form *models.LoginForm) (user *models.User, error error) {
	user = &models.User{
		UserName: form.UserName,
		Password: form.Password,
	}
	if err := dao.Login(user); err != nil {
		return nil, err
	}

	token, err := middleware.CreateToken(user.UserId, user.UserName, user.Password)
	if err != nil {
		return
	}
	user.Token = token
	return
}
