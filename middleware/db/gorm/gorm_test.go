package gorm

import (
	"database/sql"
	"fmt"
	"log"
	"testing"
	"time"

	"awesomeProject/middleware/db/gorm/gorm_structs"
	"github.com/golang-module/carbon/v2"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	dsn := "root:root@tcp(127.0.0.1:33060)/default?charset=utf8mb4&parseTime=True&loc=Local"
	gdb := New(dsn)
	if gdb.err != nil {
		log.Fatalln(gdb.err.Error())
	}
	email := "foo@qq.com"
	birthday := time.Now()
	data := gorm_structs.UserEnt{
		Name:         "xiaoming11",
		Email:        &email,
		Age:          23,
		Birthday:     &birthday,
		MemberNumber: sql.NullString{String: "ccccccccllll", Valid: true},
		ActivatedAt:  sql.NullTime{Time: time.Now(), Valid: true},
		DeletedAt:    gorm.DeletedAt{Time: carbon.Now().AddDay().ToStdTime(), Valid: false},
	}
	err := gdb.Create(&data)
	if err != nil {
		log.Fatalln(err.Error())
	}
	fmt.Println(data)
}
