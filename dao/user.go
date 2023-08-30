package dao

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"tiktok/model"
)

var (
	ErrorUserExit      = "用户已存在"
	ErrorUserNotExit   = "用户不存在"
	ErrorPasswordWrong = "密码错误"
	ErrorGenIDFailed   = errors.New("创建用户ID失败")
	ErrorInvalidID     = "无效的ID"
	ErrorQueryFailed   = "查询数据失败"
	ErrorInsertFailed  = errors.New("插入数据失败")
)

// 根据用户Id查找用户是否存在
func FindUserById(user_id int64) (*model.User, error) {
	var user *model.User
	err := DB.Where("user_id = ?", user_id).First(&user).Error
	if err != nil {
		log.Println(err.Error())
	}
	return user, err
}

// 根据用户名查找用户是否存在
func FindUserByName(username string) (*model.RegisterForm, error) {
	var user model.RegisterForm

	query := "SELECT * FROM register_forms WHERE username = ?"
	result := DB.Raw(query, username).First(&user)

	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Println("User not found:", username)
			return nil, errors.New(ErrorUserNotExit)
		}
		log.Println("Error querying database:", result.Error)
		return nil, result.Error
	}

	return &user, nil
}

// 添加新用户
func InsertUser(user *model.User) error {
	err := DB.Create(&user).Error
	return err
}

func InsertRegisterForm(user *model.RegisterForm) error {
	err := DB.Create(user).Error
	return err
}

func Login(username string) (*model.User, error) {
	var user model.User
	result := DB.Where(&model.User{UserName: username}).Find(&user)
	if result.RowsAffected > 0 {
		return &user, nil
	}
	return nil, errors.New(ErrorUserNotExit)
}
