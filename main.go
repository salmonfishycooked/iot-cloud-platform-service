package main

import (
	"fmt"
	"iot_backend/router"
	"iot_backend/util"
)

func main() {
	//go tcp.StartListenTCP() // 开启 TCP 监听
	cfg, _ := util.ParseConfig()
	app := router.GetApp()

	orm, _ := util.GetOrm()
	fmt.Println(orm)

	app.Run(":" + cfg.AppPort)
}
