package service

import "github.com/paradise98425/go-gin/entity"

type VideoService interface {
	Save(entity.Video) entity.Video
	FindAll() []entity.Video
}

type videoService struct {
	videos []entity.Video //videos as slice of video
}

//function to get the new instance of video
func New() VideoService {
	//returning a pointer to struct
	return &videoService{
		videos: []entity.Video{},
	}
}

//pass the struct videoService to implement the interface
func (service *videoService) Save(video entity.Video) entity.Video {
	service.videos = append(service.videos, video)
	return video
}

//pass the struct videoService
func (service *videoService) FindAll() []entity.Video {
	return service.videos
}
