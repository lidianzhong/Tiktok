package dao

import (
	"fmt"
	"tiktok/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

func Init() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/tiktok?charset=utf8mb4&parseTime=True&loc=Local",
		config.MysqlUsername,
		config.MysqlPassword,
		config.MysqlHost,
		config.MysqlPort,
	)

	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})

	// 更新数据库中videoURL数据
	err = UpdateVideoURL()
	if err != nil {
		return err
	}

	return nil
}
