package main

import (
	"iot_backend/router"
	"iot_backend/util"
)

func main() {
	cfg, _ := util.ParseConfig()
	app := router.GetApp()
	_, _ = util.GetOrm()

	app.Run(":" + cfg.AppPort)

	//tcp.StartListenTCP() // 开启 TCP 监听
}
