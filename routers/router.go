package routers

import (
	"github.com/gin-gonic/gin"
	"go-email/middleware/cors"
	"go-email/pkg/setting"
	"go-email/routers/api"
)

//初始化路由
func InitRouter() *gin.Engine {

	//默认初始化 Gin
	r := gin.New()
	//Logger实例将日志写入gin.DefaultWriter的日志记录器中间件。
	r.Use(gin.Logger())

	//Recovery返回一个中间件，该中间件从任何恐慌中恢复，如果有500，则写入500。
	r.Use(gin.Recovery())

	// 使用跨域中间件
	r.Use(cors.Cors())

	//设置mode-----"debug","release","test"
	gin.SetMode(setting.ServerSetting.RunMode)
	//工程名
	project := r.Group(setting.ServerSetting.ProjectName)
	//健康检查
	project.GET("/welcome", api.Welcome)
	//发送邮件
	project.POST("/send", api.SendMail)
	return r
}
