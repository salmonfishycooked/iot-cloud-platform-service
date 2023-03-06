package model

// Actuator 执行器表 映射模型
type Actuator struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Tag       string `gorm:"type:VARCHAR(255);not NULL;unique" json:"tag"`
	Name      string `gorm:"type:VARCHAR(255);not NULL" json:"name"`
	DeviceTag string `gorm:"type:VARCHAR(255);not NULL" json:"device_tag"`
	Created   int64  `gorm:"autoCreateTime" json:"created"`
	Updated   int64  `gorm:"autoUpdateTime" json:"updated"`
}
