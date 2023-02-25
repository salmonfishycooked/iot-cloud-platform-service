package util

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

type Orm struct {
	*xorm.Engine
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
	conn := database.User + ":" + database.Password + "@tcp(" + database.Host + ":" + database.Port + ")/" + database.DbName + "?charset=" + database.Charset
	engine, err := xorm.NewEngine("mysql", conn)
	if err != nil {
		return err
	}

	engine.ShowSQL(database.ShowSql)

	// 里面放入 model 结构体指针，同步能够部分智能的根据结构体的变动检测表结构的变动，并自动同步。
	//engine.Sync2(&model.User{})

	orm = &Orm{}

	orm.Engine = engine
	return nil
}
