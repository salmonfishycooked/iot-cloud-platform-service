package service

import "iot_backend/dao"

// 此处存放 Service 层通用方法

// GetInfoById
// @Description: 获取数据 by id
// @param data 绑定的 Model 层数据
// @param id 传入id
// @return int64 查询到的数据条数
func GetInfoById(data interface{}, id uint) int64 {
	counts := dao.QueryById(data, id)
	return counts
}

// CreateData
// @Description: 在数据库对应的表中创建一条记录
// @param data 需要创建的数据（绑定Model层）
func CreateData(data interface{}) error {
	err := dao.CreateData(data)
	return err
}
