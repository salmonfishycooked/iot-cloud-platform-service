package model

// Battery Battery 表 映射模型
type Battery struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Voltage     string `gorm:"not NULL" json:"voltage"`
	Current     string `gorm:"not NULL" json:"current"`
	Temperature string `gorm:"not NULL" json:"temperature"`
	Created     int64  `gorm:"autoCreateTime"`
	UpdatedAt   int    `json:"updated_at"`
}
