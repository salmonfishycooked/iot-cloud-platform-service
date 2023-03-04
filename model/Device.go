package model

// Device Device 表 映射模型
type Device struct {
	ID                uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	ElectricMachinery bool   `gorm:"not NULL" json:"electric_machinery"`
	Speed             string `gorm:"not NULL" json:"speed"`
	Rfid              uint   `gorm:"not NULL" json:"rfid"`
	Created           int64  `gorm:"autoCreateTime"`
	UpdatedAt         int    `json:"updated_at"`
}
