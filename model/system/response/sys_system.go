package response

import "github.com/haoleiqin/gin-flip-api/config"

type SysConfigResponse struct {
	Config config.Server `json:"config"`
}
