package util

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB
var err error

// 1024code 中从环境变量获取值
//var MYSQL_USERNAME = os.Getenv("MYSQL_USER")
//var MYSQL_PASSWORD = "OVxQmSxf"
//var MYSQL_HOST = os.Getenv("MYSQL_HOST")
//var MYSQL_PORT = os.Getenv("MYSQL_PORT")

var MYSQL_USERNAME = "root"
var MYSQL_PASSWORD = 1234
var MYSQL_HOST = "localhost"
var MYSQL_PORT = 3306

func init() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/tiktok?charset=utf8mb4&parseTime=True&loc=Local", MYSQL_USERNAME, MYSQL_PASSWORD, MYSQL_HOST, MYSQL_PORT)
	// dsn := "root:1234@tcp(localhost:3306)/tiktok?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		fmt.Println(err)
	}
}
