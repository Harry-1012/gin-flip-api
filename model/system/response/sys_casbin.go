package response

import (
	"github.com/haoleiqin/gin-flip-api/model/system/request"
)

type PolicyPathResponse struct {
	Paths []request.CasbinInfo `json:"paths"`
}
