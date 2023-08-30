package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"tiktok/dao"
	"tiktok/model"
	"tiktok/service"
	"tiktok/util"
	"time"
)

type FeedResponse struct {
	util.Response
	VideoList []model.Video `json:"video_list,omitempty"`
	NextTime  int64         `json:"next_time,omitempty"`
}

func Feed(c *gin.Context) {

	// 返回参数
	currentTimeStr := strconv.FormatInt(time.Now().Unix(), 10)
	latestTime := c.DefaultQuery("latest_time", currentTimeStr)

	tokenStr := c.Query("token")

	//传入了token
	if tokenStr != "" {
		tokenStruck, err := util.ParseToken(tokenStr)
		if err != nil {
			c.JSON(http.StatusOK, util.Response{
				StatusCode: 403,
				StatusMsg:  "token不正确",
			})
			c.Abort() //阻止执行
			return
		}
		//token超时
		if tokenStruck.ExpiresAt.Time.Before(time.Now()) {
			c.JSON(http.StatusOK, util.Response{
				StatusCode: 402,
				StatusMsg:  "token过期",
			})
			c.Abort() //阻止执行
			return
		}
		c.Set("user_id", tokenStruck.UserId)
	}

	// FIXME 在第一次登录抖音时，会发回错误的 latest_time 数值，为了适应这个bug而做的改动
	if len(latestTime) > 10 {
		latestTime = currentTimeStr
	}

	// 参数转换
	postTime, err := util.ConvertTimestampStrToUnix(latestTime)
	if err != nil {
		FeedErrorResponse(c, err.Error())
	}

	// 调用service层获取videoList
	videoList, err := dao.NewVideoDaoInstance().QueryFeedVideoList(postTime)
	if err != nil {
		FeedErrorResponse(c, err.Error())
	}

	// 选出videoList中最早的post_time
	nextTime := service.FindEarliestPostTime(videoList)

	// 返回数据
	c.JSON(http.StatusOK, FeedResponse{
		Response:  util.Response{StatusCode: 0},
		VideoList: videoList,
		NextTime:  nextTime,
	})
}

func FeedErrorResponse(context *gin.Context, msg string) {
	context.JSON(http.StatusOK, util.Response{StatusCode: 1, StatusMsg: msg})
}
