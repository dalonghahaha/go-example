package services

import (
	"cobra-demo/models/comjia"
)

type VideoService struct {
}

func NewVideoService() *VideoService {
	return &VideoService{}
}

func (s *VideoService) GetVideo(id int64) *comjia.CjVideoInfo {
	video := new(comjia.CjVideoInfo)
	err := video.GetByID(id)
	if err != nil {
		return nil
	} else {
		return video
	}
}
