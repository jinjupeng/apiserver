package video

import (
	"apiserver/handler"
	"apiserver/model"
	"apiserver/pkg/errno"
)

// ShowVideoService 投稿详情的服务
type ShowVideoService struct {

}

// Show 显示视频
func (service *ShowVideoService) Show(id string) handler.Response {
	var video model.VideoModel
	err := model.DB.Self.First(&video, id).Error
	if err != nil {
		return handler.Response{
			Code: errno.ErrNotFind.Code,
			Message: errno.ErrNotFind.Message,
			Data: nil,
		}
	}

	return handler.Response{
		Data: video,
	}
}
