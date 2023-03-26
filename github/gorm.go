package github

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/golang-module/carbon/v2"
	"gorm.io/driver/mysql"
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

func Demo() {
	DSN := "root:root@tcp(127.0.0.1:33060)/default?charset=utf8mb4&parseTime=True&loc=Local"
	sqlDB, err := sql.Open("mysql", DSN)
	if err != nil {
		fmt.Printf("%s", err)
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	email := "foo@qq.com"
	birthday := time.Now()
	insertUser := UserEnt{
		Name:         "xiaoming11",
		Email:        &email,
		Age:          23,
		Birthday:     &birthday,
		MemberNumber: sql.NullString{String: "ccccccccllll", Valid: true},
		ActivatedAt:  sql.NullTime{Time: time.Now(), Valid: true},
		DeletedAt:    gorm.DeletedAt{Time: carbon.Now().AddDay().ToStdTime(), Valid: false},
	}
	gormDB.Create(&insertUser)
	fmt.Printf("%v", insertUser)
}
