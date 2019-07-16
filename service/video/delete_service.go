package video

import "apiserver/model"

// Delete Video
func Delete(id uint) error {
	var video model.VideoModel
	video.ID = id
	return model.DB.Self.Delete(&video).Error
}
