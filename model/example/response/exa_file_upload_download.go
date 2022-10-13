package response

import "github.com/haoleiqin/gin-flip-api/model/example"

type ExaFileResponse struct {
	File example.ExaFileUploadAndDownload `json:"file"`
}
