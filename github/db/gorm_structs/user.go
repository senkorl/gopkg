package gorm_structs

import (
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type UserEnt struct {
	ID           uint           `json:"ID,omitempty"`
	Name         string         `json:"name,omitempty"`
	Email        *string        `json:"email,omitempty"`
	Age          uint8          `json:"age,omitempty"`
	Birthday     *time.Time     `json:"birthday,omitempty"`
	MemberNumber sql.NullString `json:"memberNumber"`
	ActivatedAt  sql.NullTime   `json:"activatedAt"`
	CreatedAt    time.Time      `json:"createdAt"`
	UpdatedAt    time.Time      `json:"updatedAt"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"deletedAt"`
}

func (u UserEnt) TableName() string {
	return "user"
}
