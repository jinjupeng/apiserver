package video

import (
	"apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
)

// 创建视频投稿服务
type CreateVideoService struct {
	Title string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Info string `form:"info" json:"info" binding:"max=3000"`
	Url string `form:"url" json:"url"`
}

// Create video
func (service *CreateVideoService) Create() handler.Response {
	video := model.VideoModel{
		Title: service.Title,
		Info: service.Info,
		Url: service.Url,
	}
	err := model.DB.Self.Create(&video).Error
	if err != nil {
		return handler.Response{
			Code:    errno.ErrCreateFail.Code,
			Message: errno.ErrCreateFail.Message,
			Data: nil,
		}
	}

	return handler.Response{
		Code: errno.OK.Code,
		Message: errno.OK.Message,
		Data: video,
	}
}