package main

import (
	"github.com/gin-gonic/gin"
	"iot_backend/util"
)

func main() {
	cfg, err := util.ParseConfig("./config/app.json")

	if err != nil {
		panic(err.Error())
	}

	app := gin.Default()

	gin.SetMode(cfg.AppMode)

	err = app.Run(cfg.AppHost + ":" + cfg.AppPort)
	if err != nil {
		panic(err.Error())
	}
}
