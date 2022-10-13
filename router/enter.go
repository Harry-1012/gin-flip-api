package router

import (
	"github.com/haoleiqin/gin-flip-api/router/autocode"
	"github.com/haoleiqin/gin-flip-api/router/example"
	"github.com/haoleiqin/gin-flip-api/router/system"
)

type RouterGroup struct {
	System   system.RouterGroup
	Example  example.RouterGroup
	Autocode autocode.RouterGroup
}

var RouterGroupApp = new(RouterGroup)
