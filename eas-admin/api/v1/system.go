package v1

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"admin/global"
	"admin/model/request"
	"admin/model/response"
	"admin/service"
	"admin/utils"
)

func Ping(c *gin.Context) {
	response.SuccessResponse(c, "pong")
}

func InitMySQL(c *gin.Context) {
	if global.EASMySql != nil {
		global.EASLog.Error("已存在数据库配置")
		response.ErrorResponseCustom(c, utils.ErrCodeMySQLError, "已存在数据库配置")
		return
	}

	var param request.InitDB
	if err := c.ShouldBind(&param); err != nil {
		global.EASLog.Error("参数校验不通过", zap.Any("err", err))
	}

	if err := service.InitDB(param); err != nil {
		global.EASLog.Error("自动创建数据库失败", zap.Any("err", err))
		response.ErrorResponseCustom(c, utils.ErrCodeMySQLError, "自动创建数据库失败，请查看后台日志，检查后在进行初始化")
		return
	}

	response.SuccessResponse(c, "数据库初始化成功")
}
