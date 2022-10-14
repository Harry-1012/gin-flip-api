package middleware

import (
	"bytes"
	"context"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
	"github.com/haoleiqin/gin-flip-api/global"
	"github.com/haoleiqin/gin-flip-api/model/common/response"
	"github.com/haoleiqin/gin-flip-api/model/system"
	"github.com/haoleiqin/gin-flip-api/utils"
)

var (
	SignWrong  = "sign wrong"
	SignNull   = "sign null"
	ApiKeyNull = "api_key null"
	UserWrong  = "用户异常"
)

func CheckSign() gin.HandlerFunc {
	return func(c *gin.Context) {
		jsonBody := make(map[string]interface{}) //注意该结构接受的内容
		c.BindJSON(&jsonBody)
		if jsonBody["sign"] == nil {
			response.FailWithDetailed(gin.H{}, SignNull, c)
			c.Abort()
			return
		}
		signReq := fmt.Sprintf("%v", jsonBody["sign"])
		delete(jsonBody, "sign")
		signStrByte, _ := json.Marshal(jsonBody)
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(signStrByte)) // 把body再写回去,不然别的地方取不到
		apikeyRedisPrex := "yld_partner_apikey_prefix_"
		apiKey := fmt.Sprintf("%v", jsonBody["api_key"])
		if apiKey == "" {
			response.FailWithDetailed(gin.H{}, ApiKeyNull, c)
			c.Abort()
			return
		}

		apiSecret := ""
		var user system.SysUser
		userJsonGet, redisGetErr := global.GVA_REDIS.Get(context.Background(), apikeyRedisPrex+apiKey).Result()
		if redisGetErr == nil && userJsonGet != "" {
			jsonErr := json.Unmarshal([]byte(userJsonGet), &user)
			if jsonErr != nil {
				fmt.Println(jsonErr, "jsonErr 20221012174241")
			} else {
				apiSecret = user.ApiSecret
			}
		}
		if apiSecret == "" { //缓存里没有,从数据库获取
			userSqlErr := global.GVA_DB.Where("api_key = ? ", apiKey).First(&user).Error
			apiSecret = user.ApiSecret
			if userSqlErr != nil {
				response.FailWithDetailed(gin.H{}, UserWrong, c)
				c.Abort()
				return
			}
			dr, err := utils.ParseDuration(global.GVA_CONFIG.JWT.ExpiresTime)
			if err != nil {
				response.FailWithDetailed(gin.H{}, "获取jwt过期时间配置文件异常", c)
				c.Abort()
				return
			}
			userJson, jsonErr := json.Marshal(user)
			if jsonErr == nil {
				_ = global.GVA_REDIS.Set(context.Background(), apikeyRedisPrex+apiKey, userJson, dr).Err()
			}
		}
		signStr := string(signStrByte) + apiSecret
		signReal := SignEncode(signStr)
		if signReq != signReal {
			response.Result(401, gin.H{}, SignWrong, c)
			global.GVA_LOG.Warn("api接口验签失败!请求sign:" + signReq + " 真实sign:" + signReal + " 签名字符串:" + signStr)
			c.Abort()
			return
		}
		c.Next()
	}
}

// @function: SignEncode
// @description: 生成sign
// @param: message string
// @return: sign string
func SignEncode(message string) string {
	return GetSHA256HashCode(message)
}

// @function: GetSHA256HashCode
// @description: SHA256生成哈希值
// @param: message string
// @return: hashCode string
func GetSHA256HashCode(message string) string {
	messageArr := []byte(message)
	//创建一个基于SHA256算法的hash.Hash接口的对象
	hash := sha256.New()
	//输入数据
	hash.Write(messageArr)
	//计算哈希值
	bytes := hash.Sum(nil)
	//将字符串编码为16进制格式,返回字符串
	hashCode := hex.EncodeToString(bytes)
	//返回哈希值
	return hashCode
}
