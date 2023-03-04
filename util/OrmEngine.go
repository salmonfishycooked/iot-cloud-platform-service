package util

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"iot_backend/model"
)

type Orm struct {
	*gorm.DB
}

var orm *Orm

// GetOrm 获取 orm 实例
func GetOrm() (*Orm, error) {
	if orm == nil {
		err := ormEngineInit(cfg)
		if err != nil {
			return nil, err
		}
	}
	return orm, nil
}

// ormEngineInit 连接数据库，并将 orm 实例赋值给全局变量 orm
func ormEngineInit(cfg *Config) error {
	database := cfg.Database
	dsn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.DbName + "?charset=" + database.Charset
	engine, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}

	orm = &Orm{engine}

	// 数据库建表
	orm.AutoMigrate(&model.Battery{}) // Battery 表
	orm.AutoMigrate(&model.Device{})  // Device 表
	orm.AutoMigrate(&model.History{}) // History 表

	return nil
}
