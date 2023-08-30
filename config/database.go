package config

import (
	"gorm.io/gorm"
)

var DB *gorm.DB

// 1024code 中从环境变量获取值

//var MysqlUsername = os.Getenv("MYSQL_USER")
//var MysqlPassword = "OIcTDrpm"
//var MysqlHost = os.Getenv("MYSQL_HOST")
//var MysqlPort = os.Getenv("MYSQL_PORT")

var MysqlUsername = "root"
var MysqlPassword = "1234"
var MysqlHost = "127.0.0.1"
var MysqlPort = "3306"
