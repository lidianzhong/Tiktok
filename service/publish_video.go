package service

import (
	"tiktok/util"
  "time"
)

type PublishVideoFlow struct {
	videoName string
	coverName string
	title     string
	authorId  int64
  PostTime time.Time
	video *util.Video
}

func (f *PublishVideoFlow) Do() error {
	if err := f.checkParam(); err != nil {
		return err
	}
	if err := f.publish(); err != nil {
		return err
	}
	return nil
}

func (f *PublishVideoFlow) checkParam() error {
	f.videoName = util.GetDataUrl(f.videoName)
	f.coverName = util.GetDataUrl(f.coverName)
	return nil
}

func (f *PublishVideoFlow) publish() error {
	video := &util.Video{
		AuthorId: f.authorId,
		PlayUrl:  f.videoName,
		CoverUrl: f.coverName,
		Title:    f.title,
    PostTime:  time.Now(),
	}
	return util.NewVideoDaoInstance().CreateVideo(video)
}

func PublishVideo(userId int64, videoName, coverName, title string) error {
	return NewPublishVideoFlow(userId, videoName, coverName, title).Do()
}

func NewPublishVideoFlow(userId int64, videoName, coverName, title string) *PublishVideoFlow {
	return &PublishVideoFlow{
		authorId:  userId,
		videoName: videoName,
		coverName: coverName,
		title:     title,
    
	}
}
