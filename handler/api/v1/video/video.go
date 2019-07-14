package video

import (
	"apiserver/model"
)

// 请求结构体
type CreateRequest struct {
	Title string `json:"username"`
	Info  string `json:"password"`
}

// 响应结构体
type CreateResponse struct {
	Title string `json:"username"`
	Info  string `json:"password"`
}

// 列表分页请求结构体
type ListRequest struct {
	Title  string `json:"username"`
	Info   string `json:"password"`
	Offset int    `json:"offset"`
	Limit  int    `json:"limit"`
}

// 列表分页响应结构体
type ListResponse struct {
	TotalCount uint64              `json:"totalCount"`
	VideoList  []*model.VideoModel `json:"userList"`
}
