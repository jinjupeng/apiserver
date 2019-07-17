package video

import (
	"apiserver/handler"
	"apiserver/pkg/errno"
	"apiserver/service/video"
	"apiserver/util"
	"github.com/gin-gonic/gin"
	"github.com/lexkong/log"
	"github.com/lexkong/log/lager"
)

func CreateVideo(c *gin.Context)  {
	log.Info("User Create function called.", lager.Data{"X-Request-Id": util.GetReqID(c)})
	service := video.CreateVideoService{}
	if err := c.ShouldBind(&service); err == nil {
		res := service.Create()
		c.JSON(200, res)
	} else {
		handler.SendResponse(c, errno.ErrBind, nil)
	}
}
