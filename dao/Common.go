package dao

import (
	"iot_backend/util"
)

// 此处存放通用数据库CRUD方法

// QueryByTag
// @Description: 查询数据 by Tag
// @param data 绑定 Model 层的数据
// @param tag 传入Tag
// @return int64 查询到的条数
func QueryByTag(data interface{}, tag string) int64 {
	db, _ := util.GetOrm()
	result := db.Where("tag = ?", tag).Find(&data)

	counts := result.RowsAffected

	return counts
}

// CreateData
// @Description: 创建数据
// @param data 基于 Model 层创建的数据
func CreateData(data interface{}) error {
	db, _ := util.GetOrm()
	result := db.Create(data)

	err := result.Error

	return err
}
