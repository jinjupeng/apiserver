package model

import _ "fmt"

// 视频映射实体
type VideoModel struct {
	BaseModel
	Title string `json:"title"`
	Info string `json:"info"`
}

// 自定义表名,覆盖GORM默认表名
func (c *VideoModel) TableName() string {
	return "tb_videos"
}