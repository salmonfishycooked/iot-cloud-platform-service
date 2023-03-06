package model

// HistoryOrderData 历史命令数据表 映射模型
type HistoryOrderData struct {
	ID      uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Tag     string `gorm:"type:VARCHAR(255);not NULL" json:"tag"`
	Value   string `gorm:"type:VARCHAR(255);" json:"value"`
	Created int64  `gorm:"autoCreateTime" json:"created"`
}
