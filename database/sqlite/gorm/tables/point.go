package tables

import "gorm.io/gorm"

type DevicePoint struct {
	PointID   int    `gorm:"primaryKey;point_id" json:"point_id"` // 点位ID
	PointName string `gorm:"type:varchar(100)" json:"point_name"` // 点位名称
	DeviceID  int    `gorm:"index" json:"device_id"`              // 设备名称
	*gorm.Model
}

func (DevicePoint) TableName() string {
	return "device_point"
}
