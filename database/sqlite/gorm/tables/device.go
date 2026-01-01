package tables

import "gorm.io/gorm"

type Device struct {
	DeviceID   int    `gorm:"index" json:"device_id"`                          // 设备ID
	DeviceName string `gorm:"type:varchar(100);primaryKey" json:"device_name"` // 设备名称
	*gorm.Model
}

func (Device) TableName() string {
	return "devices"
}
