package main

import (
	"awesomeProject/database/sqlite/gorm/tables"
	"fmt"
)

func main() {
	db := GetInstance("./sqlite.database")
	err := db.AutoMigrate(&tables.Device{}, &tables.DevicePoint{})
	if err != nil {
		panic(err)
	}
	fmt.Println("migrated")
	d := make([]*tables.Device, 0)
	for i := 0; i < 10; i++ {
		d = append(d, &tables.Device{DeviceID: i, DeviceName: fmt.Sprintf("device%d", i)})
	}
	err = db.Create(&d).Error
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("created")
	u := tables.DevicePoint{}
	db.First(&u, 1)
	fmt.Println("found", u)
}
