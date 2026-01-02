package main

import (
	"awesomeProject/database/mysql/gorm/tables"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// 连接 MySQL
	dsn := "root:root@tcp(127.0.0.1:33060)/gg?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 自动迁移表结构
	db.AutoMigrate(&tables.User{}, &tables.Order{}, &tables.Group{})

	// 创建用户
	user := tables.User{Name: "Alice"}
	db.Create(&user)

	// 创建订单并关联到用户
	order := tables.Order{Item: "Laptop", UserID: user.ID}
	db.Create(&order)

	// 创建群组
	group1 := tables.Group{Name: "VIP"}
	group2 := tables.Group{Name: "Premium"}
	db.Create(&group1)
	db.Create(&group2)

	// 关联用户到群组（多对多）
	db.Model(&user).Association("Groups").Append(&group1, &group2)

	// 查询用户及其关联数据
	var result tables.User
	db.Preload("Orders").Preload("Groups").First(&result, user.ID)

	// 打印查询结果
	fmt.Println("User:", result.Name)
	for _, order := range result.Orders {
		fmt.Println("Order:", order.Item)
	}
	for _, group := range result.Groups {
		fmt.Println("Group:", group.Name)
	}
}
