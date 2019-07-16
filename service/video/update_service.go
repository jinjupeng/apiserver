package video

import (
	"apiserver/model"
)

// UpdateVideoService 更新视频的服务
type UpdateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=30"`
	Info string `form:"info" json:"info" binding:"max=3000"`
}

func (service *UpdateVideoService) Update(id string) error {
	var video model.VideoModel

	video.Title = service.Title
	video.Info = service.Info
	return model.DB.Self.Save(&video).Error
}
