package model

// Sensor 传感器表 映射模型
type Sensor struct {
	ID        uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Tag       string `gorm:"type:VARCHAR(255);not NULL;unique" json:"tag"`
	Name      string `gorm:"type:VARCHAR(255);not NULL" json:"name"`
	Value     string `gorm:"type:VARCHAR(255);" json:"value"`
	DeviceTag string `gorm:"type:VARCHAR(255);not NULL" json:"device_tag"`
	Created   int64  `gorm:"autoCreateTime" json:"created"`
	Updated   int64  `gorm:"autoUpdateTime" json:"updated"`
}
