package system

import (
	"github.com/gin-gonic/gin"
	"github.com/haoleiqin/gin-flip-api/global"
	"github.com/haoleiqin/gin-flip-api/model/common/request"
	"github.com/haoleiqin/gin-flip-api/model/common/response"
	"github.com/haoleiqin/gin-flip-api/model/system"
	autocodeReq "github.com/haoleiqin/gin-flip-api/model/system/request"
	"github.com/haoleiqin/gin-flip-api/service"
	"go.uber.org/zap"
)

type AppLogApi struct {
}

var appLogService = service.ServiceGroupApp.SystemServiceGroup.AppLogService

// CreateAppLog 创建AppLog
// @Tags AppLog
// @Summary 创建AppLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.AppLog true "创建AppLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appLog/createAppLog [post]
func (appLogApi *AppLogApi) CreateAppLog(c *gin.Context) {
	var appLog system.AppLog
	_ = c.ShouldBindJSON(&appLog)
	if err := appLogService.CreateAppLog(appLog); err != nil {
		global.GVA_LOG.Error("创建失败!", zap.Any("err", err))
		response.FailWithMessage("创建失败", c)
	} else {
		response.OkWithMessage("创建成功", c)
	}
}

// DeleteAppLog 删除AppLog
// @Tags AppLog
// @Summary 删除AppLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.AppLog true "删除AppLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"删除成功"}"
// @Router /appLog/deleteAppLog [delete]
func (appLogApi *AppLogApi) DeleteAppLog(c *gin.Context) {
	var appLog system.AppLog
	_ = c.ShouldBindJSON(&appLog)
	if err := appLogService.DeleteAppLog(appLog); err != nil {
		global.GVA_LOG.Error("删除失败!", zap.Any("err", err))
		response.FailWithMessage("删除失败", c)
	} else {
		response.OkWithMessage("删除成功", c)
	}
}

// DeleteAppLogByIds 批量删除AppLog
// @Tags AppLog
// @Summary 批量删除AppLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body request.IdsReq true "批量删除AppLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"批量删除成功"}"
// @Router /appLog/deleteAppLogByIds [delete]
func (appLogApi *AppLogApi) DeleteAppLogByIds(c *gin.Context) {
	var IDS request.IdsReq
	_ = c.ShouldBindJSON(&IDS)
	if err := appLogService.DeleteAppLogByIds(IDS); err != nil {
		global.GVA_LOG.Error("批量删除失败!", zap.Any("err", err))
		response.FailWithMessage("批量删除失败", c)
	} else {
		response.OkWithMessage("批量删除成功", c)
	}
}

// UpdateAppLog 更新AppLog
// @Tags AppLog
// @Summary 更新AppLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data body system.AppLog true "更新AppLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"更新成功"}"
// @Router /appLog/updateAppLog [put]
func (appLogApi *AppLogApi) UpdateAppLog(c *gin.Context) {
	var appLog system.AppLog
	_ = c.ShouldBindJSON(&appLog)
	if err := appLogService.UpdateAppLog(appLog); err != nil {
		global.GVA_LOG.Error("更新失败!", zap.Any("err", err))
		response.FailWithMessage("更新失败", c)
	} else {
		response.OkWithMessage("更新成功", c)
	}
}

// FindAppLog 用id查询AppLog
// @Tags AppLog
// @Summary 用id查询AppLog
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query system.AppLog true "用id查询AppLog"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"查询成功"}"
// @Router /appLog/findAppLog [get]
func (appLogApi *AppLogApi) FindAppLog(c *gin.Context) {
	var appLog system.AppLog
	_ = c.ShouldBindQuery(&appLog)
	if err, reappLog := appLogService.GetAppLog(appLog.ID); err != nil {
		global.GVA_LOG.Error("查询失败!", zap.Any("err", err))
		response.FailWithMessage("查询失败", c)
	} else {
		response.OkWithData(gin.H{"reappLog": reappLog}, c)
	}
}

// GetAppLogList 分页获取AppLog列表
// @Tags AppLog
// @Summary 分页获取AppLog列表
// @Security ApiKeyAuth
// @accept application/json
// @Produce application/json
// @Param data query autocodeReq.AppLogSearch true "分页获取AppLog列表"
// @Success 200 {string} string "{"success":true,"data":{},"msg":"获取成功"}"
// @Router /appLog/getAppLogList [get]
func (appLogApi *AppLogApi) GetAppLogList(c *gin.Context) {
	var pageInfo autocodeReq.AppLogSearch
	_ = c.ShouldBindQuery(&pageInfo)
	if err, list, total := appLogService.GetAppLogInfoList(pageInfo); err != nil {
		global.GVA_LOG.Error("获取失败!", zap.Any("err", err))
		response.FailWithMessage("获取失败", c)
	} else {
		response.OkWithDetailed(response.PageResult{
			List:     list,
			Total:    total,
			Page:     pageInfo.Page,
			PageSize: pageInfo.PageSize,
		}, "获取成功", c)
	}
}
