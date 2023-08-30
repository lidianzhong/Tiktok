package service

import (
	"tiktok/util"
	"time"
)

func QueryFeedVideo(postTime time.Time) []util.Video {

	// 从数据库中取videoList数据
	var videoList []util.Video

	util.DB.Preload("Author").
		Where("post_time < ?", postTime).
		Order("post_time desc").
		Limit(30).
		Find(&videoList)

	return videoList
}

func FindEarliestPostTime(videoList []util.Video) int64 {
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
