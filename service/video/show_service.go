package video

import "apiserver/model"

// ShowVideoService 投稿详情的服务
type ShowVideoService struct {

}

// Show 显示视频
func (service *ShowVideoService) Show(id string) error {
	var video model.VideoModel
	return model.DB.Self.First(&video, id).Error
}
