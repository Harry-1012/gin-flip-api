package system

import (
	"github.com/gin-gonic/gin"
	v1 "github.com/haoleiqin/gin-flip-api/api/v1"
	"github.com/haoleiqin/gin-flip-api/middleware"
)

type AppLogRouter struct {
}

// InitAppLogRouter 初始化 AppLog 路由信息
func (s *AppLogRouter) InitAppLogRouter(Router *gin.RouterGroup) {
	appLogRouter := Router.Group("appLog").Use(middleware.OperationRecord())
	var appLogApi = v1.ApiGroupApp.SystemApiGroup.AppLogApi
	{
		appLogRouter.POST("createAppLog", appLogApi.CreateAppLog)             // 新建AppLog
		appLogRouter.DELETE("deleteAppLog", appLogApi.DeleteAppLog)           // 删除AppLog
		appLogRouter.DELETE("deleteAppLogByIds", appLogApi.DeleteAppLogByIds) // 批量删除AppLog
		appLogRouter.PUT("updateAppLog", appLogApi.UpdateAppLog)              // 更新AppLog
		appLogRouter.GET("findAppLog", appLogApi.FindAppLog)                  // 根据ID获取AppLog
		appLogRouter.GET("getAppLogList", appLogApi.GetAppLogList)            // 获取AppLog列表
	}
}
