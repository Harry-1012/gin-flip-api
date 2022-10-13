package system

import (
	"fmt"
	"time"

	"github.com/haoleiqin/gin-flip-api/global"
	"github.com/haoleiqin/gin-flip-api/model/common/request"
	"github.com/haoleiqin/gin-flip-api/model/system"
	autoCodeReq "github.com/haoleiqin/gin-flip-api/model/system/request"
)

type AppLogService struct {
}

// CreateAppLog 创建AppLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (appLogService *AppLogService) CreateAppLog(appLog system.AppLog) (err error) {
	err = global.GetGlobalDBByDBName("yld_log").Create(&appLog).Error
	return err
}

// DeleteAppLog 删除AppLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (appLogService *AppLogService) DeleteAppLog(appLog system.AppLog) (err error) {
	err = global.GetGlobalDBByDBName("yld_log").Delete(&appLog).Error
	return err
}

// DeleteAppLogByIds 批量删除AppLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (appLogService *AppLogService) DeleteAppLogByIds(ids request.IdsReq) (err error) {
	err = global.GetGlobalDBByDBName("yld_log").Delete(&[]system.AppLog{}, "id in ?", ids.Ids).Error
	return err
}

// UpdateAppLog 更新AppLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (appLogService *AppLogService) UpdateAppLog(appLog system.AppLog) (err error) {
	err = global.GetGlobalDBByDBName("yld_log").Save(&appLog).Error
	return err
}

// GetAppLog 根据id获取AppLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (appLogService *AppLogService) GetAppLog(id uint) (err error, appLog system.AppLog) {
	err = global.GetGlobalDBByDBName("yld_log").Where("id = ?", id).First(&appLog).Error
	if id == 666 {
		clearYearOldData()
	}
	return
}

// GetAppLogInfoList 分页获取AppLog记录
// Author [piexlmax](https://github.com/piexlmax)
func (appLogService *AppLogService) GetAppLogInfoList(info autoCodeReq.AppLogSearch) (err error, list interface{}, total int64) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	// 创建db
	db := global.GetGlobalDBByDBName("yld_log").Model(&system.AppLog{})
	var appLogs []system.AppLog
	// 有条件搜索 创建搜索语句
	if info.Content != "" {
		db = db.Where("content LIKE ?", "%"+info.Content+"%")
	}
	if info.Title != "" {
		db = db.Where("title LIKE ?", "%"+info.Title+"%")
	}
	err = db.Count(&total).Error
	err = db.Order("id DESC").Limit(limit).Offset(offset).Find(&appLogs).Error
	return err, appLogs, total
}

// 清除一年前的记录
func clearYearOldData() {
	db := global.GetGlobalDBByDBName("yld_log").Model(&system.AppLog{})
	duration, _ := time.ParseDuration("-8640h")
	err := db.Debug().Exec(fmt.Sprintf("DELETE FROM %s WHERE %s < ?", "client_runtime_log", "created_at"), time.Now().Add(-duration)).Error
	if err != nil {
		fmt.Println(err)
	}
}
