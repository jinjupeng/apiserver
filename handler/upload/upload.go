package upload

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
		if ! (upload.CheckImageExt(imageName) && upload.CheckImageSize(file)) {
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
