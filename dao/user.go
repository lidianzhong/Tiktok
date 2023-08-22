package dao

import (
	"errors"
	"log"
	"tiktok/models"
	"tiktok/util"
)

var (
	ErrorUserExit      = "用户已存在"
	ErrorUserNotExit   = "用户不已存在"
	ErrorPasswordWrong = "密码错误"
	ErrorGenIDFailed   = errors.New("创建用户ID失败")
	ErrorInvalidID     = "无效的ID"
	ErrorQueryFailed   = "查询数据失败"
	ErrorInsertFailed  = errors.New("插入数据失败")
)

// 根据用户名查找用户是否存在
func FindUserByName(username string) (models.RegisterForm, error) {
	var user models.RegisterForm
	err := util.DB.Where("username = ?", username).First(&user).Error
	if err != nil {
		log.Println(err.Error())
	}
	return user, err
}

// 添加新用户
func InsertUser(user *models.User) error {
	err := util.DB.Create(&user).Error
	return err
}

func Login(user *models.User) (err error) {
	var u models.User
	result := util.DB.Where(&models.User{UserName: user.UserName}).Find(&u)
	if result.RowsAffected > 0 {
		return nil
	}
	return errors.New(ErrorUserNotExit)
}
