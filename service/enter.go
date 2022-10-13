package service

import (
	"github.com/haoleiqin/gin-flip-api/service/autocode"
	"github.com/haoleiqin/gin-flip-api/service/example"
	"github.com/haoleiqin/gin-flip-api/service/system"
)

type ServiceGroup struct {
	SystemServiceGroup   system.ServiceGroup
	ExampleServiceGroup  example.ServiceGroup
	AutoCodeServiceGroup autocode.ServiceGroup
}

var ServiceGroupApp = new(ServiceGroup)
