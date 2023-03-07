package main

import (
	"iot_backend/router"
	"iot_backend/tcp"
	"iot_backend/util"
)

func main() {
	cfg, _ := util.ParseConfig() // 读取配置
	app := router.GetApp()       // 初始化 HTTP 服务框架以及路由

	go tcp.StartListenTCP()    // 开启 TCP 监听
	app.Run(":" + cfg.AppPort) // 运行 HTTP 服务
}
