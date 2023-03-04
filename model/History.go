package model

type History struct {
	ID      uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Rfid    uint   `gorm:"not NULL" json:"rfid"`
	Type    string `gorm:"not NULL" json:"type"`
	Value   string `gorm:"not NULL" json:"value"`
	Created int64  `gorm:"autoCreateTime"`
}
