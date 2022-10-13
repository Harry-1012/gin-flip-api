package system

import (
	"github.com/haoleiqin/gin-flip-api/global"
)

type JwtBlacklist struct {
	global.GVA_MODEL
	Jwt string `gorm:"type:text;comment:jwt"`
}
