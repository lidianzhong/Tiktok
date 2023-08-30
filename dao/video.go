package dao

import (
	"sync"
	"tiktok/model"
	"time"
)

type VideoDao struct{}

var videoDao *VideoDao
var videoOnce sync.Once

func NewVideoDaoInstance() *VideoDao {
	videoOnce.Do(
		func() {
			videoDao = &VideoDao{}
		})
	return videoDao
}

func (*VideoDao) CreateVideo(video *model.Video) error {
	if err := DB.Create(video).Error; err != nil {
		return err
	}
	return nil
}

func (*VideoDao) QueryVideoCountByUserId(userId int64, count *int64) error {
	if err := DB.Model(&model.Video{}).Where("author_id = ?", userId).Count(count).Error; err != nil {
		return err
	}
	return nil
}

func (*VideoDao) QueryFeedVideoList(postTime time.Time) ([]model.Video, error) {
	// 从数据库中取videoList数据
	var videoList []model.Video
	if err := DB.Preload("Author").Where("post_time < ?", postTime).Order("post_time desc").Limit(30).Find(&videoList).Error; err != nil {
		return nil, err
	}

	return videoList, nil
}
