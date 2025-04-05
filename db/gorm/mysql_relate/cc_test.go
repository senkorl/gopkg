package main

import (
	"fmt"
	"log"
	"testing"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestQuery(T *testing.T) {
	dsn := "root:root@tcp(127.0.0.1:33060)/gg?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	u := &User{}
	db.Preload("Orders").Preload("Groups").First(u, "id = 3")
	fmt.Println("User:", u)
}

func TestCreate(T *testing.T) {
	c := 1 % 7
	log.Print(c)
	c1 := 7 % 7
	log.Print(c1)

}
