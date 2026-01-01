package main

import (
	"fmt"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db := GetInstance("./sqlite.database")
	err := db.AutoMigrate(&Device{}, &DevicePoint{})
	if err != nil {
		panic(err)
	}
	fmt.Println("migrated")
	d := make([]*Device, 0)
	for i := 0; i < 10; i++ {
		d = append(d, &Device{DeviceID: i, DeviceName: fmt.Sprintf("device%d", i)})
	}
	err = db.Create(&d).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("created")
	u := DevicePoint{}
	db.First(&u, 1)
	fmt.Println("found", u)
}

var (
	database *gorm.DB
	dbOnce   sync.Once
)

func GetInstance(dbPath string) *gorm.DB {
	dbOnce.Do(func() {
		db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{
			Logger: logger.Default.LogMode(logger.Info), // 启用调试模式
		})
		if err != nil {
			fmt.Printf("open sqlite failed:%s", err.Error())
		} else {
			database = db
		}
	})
	return database
}
