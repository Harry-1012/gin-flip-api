package request

import (
	"github.com/haoleiqin/gin-flip-api/model/autocode"
	"github.com/haoleiqin/gin-flip-api/model/common/request"
)

type {{.StructName}}Search struct{
    autocode.{{.StructName}}
    request.PageInfo
}