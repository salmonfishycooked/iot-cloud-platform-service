package model

// Device Device 表 映射模型
type Device struct {
	ID                uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	ElectricMachinery bool   `json:"electric_machinery"`
	Speed             string `json:"speed"`
	Rfid              uint   `json:"rfid"`
	UpdatedAt         int    `json:"updated_at"`
}
