package tables

import (
	"time"
)

type DevicePoint struct {
	ID      int    `gorm:"primaryKey;autoIncrement" json:"id"`
	PointID int    `gorm:"uniqueIndex" json:"point_id"`
	Name    string `gorm:"type:varchar(100)" json:"point_name"`

	DeviceID int    `gorm:"index" json:"device_id"`
	Device   Device `gorm:"foreignKey:DeviceID;references:DeviceID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (DevicePoint) TableName() string {
	return "device_point"
}
