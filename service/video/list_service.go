package video

import "apiserver/model"

// ListVideoService 视频列表服务
type ListVideoService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// List 视频列表
func (service ListVideoService) List() error {
	var videos []model.VideoModel

	return model.DB.Self.Find(&videos).Limit(service.Limit).Offset(service.Start).Error
}