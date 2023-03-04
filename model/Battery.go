package model

// Battery Battery 表 映射模型
type Battery struct {
	ID          uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Voltage     string `json:"voltage"`
	Current     string `json:"current"`
	Temperature string `json:"temperature"`
	UpdatedAt   int    `json:"updated_at"`
}
