package models

import (
	"time"
)

// 系统操作日志
type SysOperationLog struct {
	Model
	ApiDesc    string        `json:"apiDesc" gorm:"comment:'接口说明';size:255"`
	Path       string        `json:"path" gorm:"comment:'访问路径';size:255"`
	Method     string        `json:"method" gorm:"comment:'请求方式';size:128"`
	Body       string        `json:"body" gorm:"type:blob;comment:'请求主体(通过二进制存储节省空间)';size:255"`
	Data       string        `json:"data" gorm:"type:blob;comment:'响应数据(通过二进制存储节省空间)';size:255"`
	Status     int           `json:"status" gorm:"comment:'响应状态码'"`
	Username   string        `json:"username" gorm:"comment:'用户登录名';size:128"`
	RoleName   string        `json:"roleName" gorm:"comment:'用户角色名';size:128"`
	Ip         string        `json:"ip" gorm:"comment:'Ip地址';size:128"`
	IpLocation string        `json:"ipLocation" gorm:"comment:'Ip所在地';size:128"`
	Latency    time.Duration `json:"latency" gorm:"comment:'请求耗时(ms)'"`
	UserAgent  string        `json:"userAgent" gorm:"comment:'浏览器标识';size:255"`
}

func (m SysOperationLog) TableName() string {
	return m.Model.TableName("sys_operlog")
}