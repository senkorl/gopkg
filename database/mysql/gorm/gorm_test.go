package main

import (
	"awesomeProject/database/mysql/gorm/tables"
	"fmt"
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCreate(T *testing.T) {
	dsn := "root:root@tcp(127.0.0.1:33060)/gg?charset=utf8mb4&parseTime=True&loc=Local"
	gdb := New(dsn)
	if gdb.err != nil {
		log.Fatal(gdb.err)
	}
	user := tables.User{
		Name: "Foo",
	}
	err := gdb.Create(&user)
	if err != nil {
		log.Fatal(err)
	}
}

func TestQuery(T *testing.T) {
	dsn := "root:root@tcp(127.0.0.1:33060)/gg?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	u := &tables.User{}
	db.Preload("Orders").Preload("Groups").First(u, "id = 3")
	fmt.Println("User:", u)
}

func TestRelate(T *testing.T) {
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
