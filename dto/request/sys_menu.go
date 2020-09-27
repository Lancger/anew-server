package request

import "anew-server/dto/response"

// 获取菜单列表结构体
type MenuListReq struct {
	Title      string `json:"title" form:"title"`
	Path       string `json:"path" form:"path"`
	Status     *bool  `json:"status" form:"status"`
	Hidden    *bool  `json:"hidden" form:"hidden"`
	Creator    string `json:"creator" form:"creator"`
	response.PageInfo          // 分页参数
}

// 创建菜单结构体
type CreateMenuReq struct {
	Title      string `json:"title" validate:"required"`
	Icon       string `json:"icon"`
	Path       string `json:"path"`
	Sort       int    `json:"sort"`
	Status     *bool  `json:"status"`
	Hidden    *bool  `json:"hidden"`
	ParentId   uint   `json:"parentId"`
	Creator    string `json:"creator"`
}

// 翻译需要校验的字段名称
func (s CreateMenuReq) FieldTrans() map[string]string {
	m := make(map[string]string, 0)
	m["Title"] = "菜单名称"
	return m
}
