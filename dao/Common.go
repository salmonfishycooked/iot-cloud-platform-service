package dao

import (
	"iot_backend/util"
)

// 此处存放通用数据库CRUD方法

// QueryById
// @Description: 查询数据 by id
// @param data 绑定 Model 层的数据
// @param id 传入id
// @return int64 查询到的条数
func QueryById(data interface{}, id uint) int64 {
	db, _ := util.GetOrm()
	result := db.First(data, id)

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
