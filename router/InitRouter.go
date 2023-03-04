package router

import (
	"github.com/gin-gonic/gin"
	"iot_backend/controller"
	"iot_backend/util"
)

type App struct {
	*gin.Engine
}

var app *App

// GetApp 获取 Gin 引擎实例（单例模式）
// *App 返回 App 实例
func GetApp() *App {
	if app == nil {
		initApp()
	}
	return app
}

// initApp 初始化 Gin 引擎
func initApp() {
	cfg, err := util.ParseConfig() // 读取 app.json 配置项
	if err != nil {
		panic(err.Error())
	}
	app = &App{gin.Default()}
	gin.SetMode(cfg.AppMode) // 设置 Gin 运行的模式

	initRoutes() // 初始化路由
}

// initRoutes 初始化路由
func initRoutes() {
	v1 := app.Group("/api/v1")
	{
		controller.InitBatteryRoutes(v1)
		controller.InitDeviceRoutes(v1)
	}
}
