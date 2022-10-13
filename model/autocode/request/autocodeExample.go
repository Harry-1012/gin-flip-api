// SysDictionaryDetail
package request

import (
	"github.com/haoleiqin/gin-flip-api/model/autocode"
	"github.com/haoleiqin/gin-flip-api/model/common/request"
)

// 如果含有time.Time 请自行import time包
type AutoCodeExampleSearch struct {
	autocode.AutoCodeExample
	request.PageInfo
}
