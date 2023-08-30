package util

import (
	"sync"
	"time"
)

type Video struct {
	Id            int64     `json:"id"`
	AuthorId      int64     `json:"author_id"`
	PlayUrl       string    `json:"play_url"`
	CoverUrl      string    `json:"cover_url"`
	FavoriteCount int64     `json:"favorite_count"`
	CommentCount  int64     `json:"comment_count"`
	IsFavorite    bool      `json:"is_favorite"`
	Title         string    `json:"title"`
	Author        Author    `json:"author"`
	PostTime      time.Time `json:"post_time"`
}

// TableName 表示配置操作数据库的表名称
func (Video) TableName() string {
	return "video"
}

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

func (*VideoDao) CreateVideo(video *Video) error {
	if err := DB.Create(video).Error; err != nil {
		return err
	}
	return nil
}

func (*VideoDao) QueryVideoCountByUserId(userId int64, count *int64) error {
	if err := DB.Model(&Video{}).Where("author_id = ?", userId).Count(count).Error; err != nil {
		return err
	}
	return nil
}



