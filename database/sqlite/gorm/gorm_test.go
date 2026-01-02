package main

import (
	"awesomeProject/database/sqlite/gorm/tables"
	"fmt"
	"log"
	"testing"

	"gorm.io/gorm"
)

func TestAutoMigrateAndRelateOperate(t *testing.T) {
	db := GetInstance("./devices.database")

	// AutoMigrate（必须先 Device）
	err := db.AutoMigrate(
		&tables.Device{},
		&tables.DevicePoint{},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("migrated")

	// 创建设备 + 点位（一对多）
	deviceData := tables.Device{
		DeviceID: 1001,
		Name:     "device-1001",
		Points: []tables.DevicePoint{
			{PointID: 1, Name: "temperature"},
			{PointID: 2, Name: "humidity"},
		},
	}

	if err := db.Create(&deviceData).Error; err != nil {
		panic(err)
	}

	// 给已有设备新增点位
	var device tables.Device
	db.Where("device_id = ?", 1001).First(&device)

	point := tables.DevicePoint{
		PointID:  3,
		Name:     "pressure",
		DeviceID: device.ID,
	}

	db.Create(&point)

	dev, err := queryDeviceInfo(db, 1001)
	if err != nil {
		panic(err)
	}
	log.Println(dev)

	p, err := queryPointInfo(db, 1)
	if err != nil {
		panic(err)
	}
	log.Println(p)
}

func queryDeviceInfo(db *gorm.DB, deviceID int) (*tables.Device, error) {
	// 查询设备 + 所有点位（Preload）
	var deviceInfo tables.Device
	err := db.Preload("Points").
		Where("device_id = ?", deviceID).
		First(&deviceInfo).Error

	if err != nil {
		panic(err)
	}
	return &deviceInfo, nil
}

func queryPointInfo(db *gorm.DB, pointID int) (*tables.DevicePoint, error) {
	var point tables.DevicePoint

	err := db.
		Preload("Device.Points").
		Where("point_id = ?", pointID).
		First(&point).Error

	if err != nil {
		panic(err)
	}

	// 当前点位
	fmt.Println("current point:", point.Name)

	// 所属设备
	fmt.Println("device:", point.Device.Name)

	// 该设备下的所有点位
	for _, p := range point.Device.Points {
		fmt.Println(" -", p.Name)
	}
	return &point, nil
}
