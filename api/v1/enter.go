package v1

import (
	"github.com/haoleiqin/gin-flip-api/api/v1/autocode"
	"github.com/haoleiqin/gin-flip-api/api/v1/example"
	"github.com/haoleiqin/gin-flip-api/api/v1/system"
)

type ApiGroup struct {
	ExampleApiGroup  example.ApiGroup
	SystemApiGroup   system.ApiGroup
	AutoCodeApiGroup autocode.ApiGroup
}

var ApiGroupApp = new(ApiGroup)
