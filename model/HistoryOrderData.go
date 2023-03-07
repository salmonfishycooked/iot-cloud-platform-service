package model

// HistoryOrderData 历史命令数据表 映射模型
type HistoryOrderData struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	DeviceTag   string `gorm:"type:VARCHAR(255);not NULL" json:"device_tag"`
	ActuatorTag string `gorm:"type:VARCHAR(255);not NULL" json:"actuator_tag"`
	Value       string `gorm:"type:VARCHAR(255);" json:"value"`
	Created     int64  `gorm:"autoCreateTime" json:"created"`
}
