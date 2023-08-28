package controller

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
	"tiktok/util"
)

func TestPublishVideoController(t *testing.T) {
	fileName := "cat"
	fullName := "cat.mp4"
	staticSourcePath = "../static"
	//	写入static
	savePath := filepath.Join(staticSourcePath, fullName)
	//  获取视频封面并写入static
	snapshotPath := filepath.Join(staticSourcePath, fileName)
	coverName, _ := util.GetSnapshot(savePath, snapshotPath, 5)
	fmt.Println(savePath)
	fmt.Println(snapshotPath)
	fmt.Println(coverName)
	_, err := os.Stat("D:/GoProject/Tiktok/static/cat.mp4")
	fmt.Println(err)
	assert.ErrorIs(t, err, nil)
}
