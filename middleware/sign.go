package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/haoleiqin/gin-flip-api/global"
	"github.com/haoleiqin/gin-flip-api/model/common/response"
	"github.com/haoleiqin/gin-flip-api/utils"
)

var (
	SignWrong     = "sign wrong"
	NonceNull     = "nonce null"
	TimestampNull = "timestamp null"
	SignNull      = "sign null"
)

func CheckSign() gin.HandlerFunc {
	return func(c *gin.Context) {
		jsonBody := make(map[string]interface{}) //注意该结构接受的内容
		c.BindJSON(&jsonBody)
		nullRes := ""
		if jsonBody["sign"] == nil {
			nullRes = SignNull
		}
		if jsonBody["timestamp"] == nil {
			nullRes = TimestampNull
		}
		if jsonBody["nonce"] == nil {
			nullRes = NonceNull
		}
		if nullRes != "" {
			response.FailWithDetailed(gin.H{}, nullRes, c)
			c.Abort()
			return
		}
		signReq := fmt.Sprintf("%v", jsonBody["sign"])
		delete(jsonBody, "sign")

		//先按key 排序 升序 ASCII 升序
		keys := make([]string, 0, len(jsonBody))
		if len(jsonBody) > 0 {
			for k := range jsonBody {
				keys = append(keys, k)
			}
		}
		// /排序
		sort.Strings(keys)
		jsonBodySort := make(map[string]interface{})
		for _, k := range keys {
			jsonBodySort[k] = jsonBody[k]
		}
		signStrByte, _ := json.Marshal(jsonBodySort)
		c.Request.Body = io.NopCloser(bytes.NewBuffer(signStrByte)) // 把body再写回去,不然别的地方取不到
		//接受到的消息
		acceptmsg := []byte(signStrByte)
		//接受到的签名
		acceptsign := signReq
		//验证签名
		signVerifyResult := utils.VerifySign(acceptmsg, acceptsign, "public.pem")
		//比较签名
		if !signVerifyResult {
			// 生成签名
			sign := utils.GetSign(signStrByte, "private.pem")
			// fmt.Println("realsign:", string(sign))
			global.GVA_LOG.Warn("api接口验签失败!请求:" + string(signStrByte) + " 接收到的签名:" + acceptsign + " 正确真实签名:" + sign)
			response.Result(401, gin.H{}, SignWrong, c)
			c.Abort()
			return
		}
		c.Next()
	}
}
