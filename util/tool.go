package util

import (
	"bytes"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const IP = "127.0.0.1"
const Port = "8080"

var supportedFormats = map[string]struct{}{
	".mp4": {},
	//	todo 添加新支持的视频格式 ~
}

// ConvertTimeToTimestampStr 接收一个time.Time的类型数据，返回string类型、以s为单位的时间戳
func ConvertTimeToTimestampStr(t time.Time) string {
	return strconv.FormatInt(t.Unix(), 10)
}

func ConvertTimestampStrToUnix(timestampStr string) (time.Time, error) {
	timestamp, err := strconv.ParseInt(timestampStr, 10, 64) // 将时间戳字符串解析为 int64 类型的时间戳值
	t := time.Unix(timestamp, 0)                             // 使用 time.Unix() 函数将时间戳值转换为 time.Time 类型
	if err != nil {
		return time.Time{}, err
	}

	return t, nil
}

// GetDataUrl 通过文件名获取对应url
func GetDataUrl(name string) string {
	url := fmt.Sprintf("http://%s:%d/static/%s", IP, Port, name)
	return url
}

// IsSupportedVideoFormat 判断文件是否为视频格式 可在supportedFormats增加新支持的视频格式
func IsSupportedVideoFormat(ext string) bool {
	_, ok := supportedFormats[ext]
	return ok
}

// GenerateUniqueFileName 通过用户名生成唯一的文件名 userid + timestamp
func GenerateUniqueFileName(userid int64) string {
	timestamp := time.Now().Unix() / int64(time.Millisecond)
	return fmt.Sprintf("%d_%d", userid, timestamp)
}

// GetSnapshot
// 传参 视频地址 封面保存地址 获取第几帧 eg static\cat.mp4 static\cat 5
// 返回 封面名 eg cat.png
func GetSnapshot(videoPath, snapshotPath string, frameNum int) (snapshotName string, err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		log.Fatal("生成缩略图失败：", err)
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		log.Fatal("解码缩略图失败：", err)
		return "", err
	}

	err = imaging.Save(img, snapshotPath+".png")
	if err != nil {
		log.Fatal("保存缩略图失败：", err)
		return "", err
	}

	fmt.Println("--snapshotPath--", snapshotPath)
	// --snapshotPath-- ./static/testImage

	names := strings.Split(snapshotPath, `\`)
	fmt.Println("----names----", names)
	// ----names---- [./static/testImage]
	// 这里把 snapshotPath 的 string 类型转换成 []string

	snapshotName = names[len(names)-1] + ".png"
	fmt.Println("----snapshotName----", snapshotName)
	// ----snapshotName---- ./static/testImage.png

	return snapshotName, nil
}
