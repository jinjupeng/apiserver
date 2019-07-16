package video

import (
	"apiserver/model"
	"apiserver/pkg/errno"
)

// 创建视频投稿服务
type CreateVideoService struct {
	model.BaseModel
	Title string `form:"title" json:"title" binding:"required,min=2,max=100"`
	Info string `form:"info" json:"info" binding:"max=3000"`
	Url string `form:"url" json:"url"`
}

// Create video
func (service *CreateVideoService) Create() errno.Errno {
	video := model.VideoModel{
		Title: service.Title,
		Info: service.Info,
		Url: service.Url,
	}
	err := model.DB.Self.Create(&video).Error
	if err != nil {
		return errno.Errno{
			Code:    errno.ErrCreateFail.Code,
			Message: errno.OK.Message,
		}
	}

	return errno.Errno{
		Code: errno.OK.Code,
		Message: errno.OK.Message,
	}
}