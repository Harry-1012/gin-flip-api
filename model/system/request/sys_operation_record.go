package request

import (
	"github.com/haoleiqin/gin-flip-api/model/common/request"
	"github.com/haoleiqin/gin-flip-api/model/system"
)

type SysOperationRecordSearch struct {
	system.SysOperationRecord
	request.PageInfo
}
