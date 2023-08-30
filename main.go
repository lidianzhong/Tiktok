package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"tiktok/dao"
	"tiktok/router"
)

func main() {
	if err := Init(); err != nil {
		os.Exit(-1)
	}

	r := gin.Default()

	// 基础接口路由
	router.BaseRoutersInit(r)

	err := r.Run()
	if err != nil {
		return
	}

}

func Init() error {
	if err := dao.Init(); err != nil {
		return err
	}
	return nil
}
