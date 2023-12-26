package service

import (
	"intro-gin/entity"
)

type ServiceVideo interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type serviceVideo struct {
	videos []entity.Video
}

func New() ServiceVideo {
	return &serviceVideo{}
}

func (s *serviceVideo) Save(video entity.Video) entity.Video {
	s.videos = append(s.videos, video)
	return video
}

func (s *serviceVideo) FindAll() []entity.Video {
	return s.videos
}
