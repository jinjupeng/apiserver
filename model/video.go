package model

import (
	_ "fmt"
	"github.com/jinzhu/gorm"
)

// 视频映射实体
type VideoModel struct {
	gorm.Model
	Title string `json:"title"`
	Info string `json:"info"`
	Url string `json:"url"`
}

// 自定义表名,覆盖GORM默认表名
func (c *VideoModel) TableName() string {
	return "tb_videos"
}