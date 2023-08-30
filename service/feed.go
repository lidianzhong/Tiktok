package service

import (
	"tiktok/model"
	"time"
)

func FindEarliestPostTime(videoList []model.Video) int64 {
	var nextTime int64 = time.Now().Unix()
	if len(videoList) > 0 {
		for _, video := range videoList {
			videoTime := video.PostTime.Unix()
			if videoTime < nextTime {
				nextTime = video.PostTime.Unix()
			}
		}
	}
	return nextTime
}
