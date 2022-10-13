package utils

import (
	"fmt"
	"sort"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type RequestHttp struct {
	Ctx    *gin.Context
	Params map[string]interface{}
	lock   sync.Mutex
}

//@function: JoinParamsStr
//@description: 请求参数转换字符串
//@param: message string
//@return: a=1&b=2 string
func (r *RequestHttp) JoinParamsStr() string {
	//先按key 排序 升序 ASCII 升序
	keys := make([]string, 0, len(r.Params))
	for k := range r.Params {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	var params []string
	if len(r.Params) > 0 {
		for _, k := range keys {
			params = append(params, fmt.Sprintf("%s=%v", k, r.Params[k]))
		}
	}
	return strings.Join(params, "&")
}

//@function: RequestParams
//@description: 获取参数集合
//@param: exclude string 排除key
//@return: hashCode string
func (r *RequestHttp) RequestParams(exclude string) string {
	ctx := r.Ctx
	bindParams := map[string]interface{}{}
	if ctx.Request.Method == "POST" {
		contextType := ctx.Request.Header.Get("Content-Type")
		if contextType == "application/json" {
			err := ctx.ShouldBindBodyWith(&bindParams, binding.JSON)
			if err != nil { //报错
				fmt.Printf("nyx_request_mid_error %v,err: %v \n", bindParams, err)
				return ""
			}
			if len(bindParams) > 0 {
				for k, v := range bindParams {
					r.Add(k, v)
				}
			}
		} else {
			_ = ctx.Request.ParseMultipartForm(32 << 20)
			if len(ctx.Request.PostForm) > 0 {
				for k, v := range ctx.Request.PostForm {
					r.Add(k, v[0])
				}
			}
		}
	} else {
		var tmpParams = make(map[string]string)
		err2 := ctx.ShouldBind(&tmpParams)
		if err2 != nil {
			fmt.Printf("nyx_request_mid_error %v,err: %v \n", bindParams, err2)
			return ""
		}
		for k, v := range tmpParams {
			r.Add(k, v)
		}
	}
	r.Delete(exclude)
	return r.JoinParamsStr()
}

//添加参数
func (r *RequestHttp) Add(key string, value interface{}) {
	r.lock.Lock()
	r.Params[key] = value
	r.lock.Unlock()
}

//删除参数
func (r *RequestHttp) Delete(key string) {
	r.lock.Lock()
	if _, ok := r.Params[key]; ok {
		delete(r.Params, key)
	}
	r.lock.Unlock()
}
