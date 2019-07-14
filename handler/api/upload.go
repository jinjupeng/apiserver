package api

import (
	. "apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/pkg/upload"
	"github.com/gin-gonic/gin"
)

// 上传图片
func UplaodImage(c *gin.Context) {
	data := make(map[string]string)
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		SendResponse(c, errno.InternalServerError, nil)
		return
	}

	if image == nil {
		SendResponse(c, errno.InternalServerError, nil)
		return
	} else {
		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()

		src := fullPath + imageName
		if !(upload.CheckImageExt(imageName) && upload.CheckImageSize(file)) {
			SendResponse(c, errno.ErrUploadCheckImageFormat, nil)
		} else {
			err := upload.CheckImage(fullPath)
			if err != nil {
				SendResponse(c, errno.ErrUploadCheckImageFail, nil)
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				SendResponse(c, errno.ErrUploadSaveImageFail, nil)
			} else {
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = savePath + imageName
				SendResponse(c, errno.OK, data)
			}
		}
	}

}

// 上传视频
func UploadVideo(c *gin.Context) {
	data := make(map[string]string)
	file, video, err := c.Request.FormFile("video")
	if err != nil {
		SendResponse(c, errno.InternalServerError, nil)
		return
	}

	if video == nil {
		SendResponse(c, errno.InternalServerError, nil)
		return
	} else {
		videoName := upload.GetVideoName(video.Filename)
		fullPath := upload.GetVideoFullPath()
		savePath := upload.GetVideoPath()

		src := fullPath + videoName
		if !(upload.CheckVideoExt(videoName) && upload.CheckVideoSize(file)) {
			SendResponse(c, errno.ErrUploadCheckVideoFormat, nil)
		} else {
			err := upload.CheckVideo(fullPath)
			if err != nil {
				SendResponse(c, errno.ErrUploadCheckVideoFail, nil)
			} else if err := c.SaveUploadedFile(video, src); err != nil {
				SendResponse(c, errno.ErrUploadSaveVideoFail, nil)
			} else {
				data["video_url"] = upload.GetVideoFullUrl(videoName)
				data["video_save_url"] = savePath + videoName
				SendResponse(c, errno.OK, data)
			}
		}
	}
}
