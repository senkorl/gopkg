package tables

import (
	"time"
)

type Device struct {
	ID       int           `gorm:"primaryKey;autoIncrement" json:"id"`
	DeviceID int           `gorm:"uniqueIndex" json:"device_id"`
	Name     string        `gorm:"type:varchar(100)" json:"device_name"`
	Points   []DevicePoint `gorm:"foreignKey:DeviceID;references:DeviceID"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Device) TableName() string {
	return "devices"
}
