// AppLog
package system

import (
	"github.com/haoleiqin/gin-flip-api/global"
)

// AppLog 结构体
// 如果含有time.Time 请自行import time包
type AppLog struct {
	global.GVA_MODEL
	Title   string `json:"title" form:"title" gorm:"column:title;comment:请求标题;type:varchar(255);"`
	Content string `json:"content" form:"content" gorm:"column:content;comment:请求内容;type:text;"`
	Ip      string `json:"ip" form:"ip" gorm:"column:ip;comment:请求时间;ip:timestamp"`
}

// TableName AppLog 表名
func (AppLog) TableName() string {
	return "client_runtime_log"
}
