package video

import (
	"apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
)

// ListVideoService 视频列表服务
type ListVideoService struct {
	Limit int `form:"limit"`
	Start int `form:"start"`
}

// List 视频列表
func (service ListVideoService) List() handler.Response {
	var videos []model.VideoModel

	err := model.DB.Self.Find(&videos).Limit(service.Limit).Offset(service.Start).Error

	if err != nil {
		return handler.Response{
			Code: errno.ErrDatabase.Code,
			Message: errno.ErrDatabase.Message,
			Data: nil,
		}
	}
	return handler.Response{
		Code: errno.OK.Code,
		Message: errno.OK.Message,
		Data: videos,
	}
}