package video

import (
	"apiserver/service/video"
	"github.com/gin-gonic/gin"
)

// ShowVideo 视频详情接口
func ShowVideo(c *gin.Context) {
	service := video.ShowVideoService{}
	res := service.Show(c.Param("id"))
	c.JSON(200, res)
}