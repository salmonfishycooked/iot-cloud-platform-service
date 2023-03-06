package model

// Device 设备表 映射模型
type Device struct {
	ID       uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Tag      string `gorm:"type:VARCHAR(255);not NULL;unique" json:"tag"`
	Name     string `gorm:"type:VARCHAR(255);not NULL" json:"name"`
	IsOnline bool   `gorm:"default:false" json:"is_online"`
	Created  int64  `gorm:"autoCreateTime" json:"created"`
	Updated  int64  `gorm:"autoUpdateTime" json:"updated"`
}
