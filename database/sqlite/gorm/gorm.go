package main

import (
	"fmt"
	"sync"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

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
