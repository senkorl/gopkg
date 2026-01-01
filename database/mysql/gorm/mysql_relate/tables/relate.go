package tables

// User 用户表
type User struct {
	ID     uint    `gorm:"primaryKey"`
	Name   string  `gorm:"size:100;not null"`
	Orders []Order `gorm:"foreignKey:UserID"`      // 一对多关联
	Groups []Group `gorm:"many2many:user_groups;"` // 多对多关联
}

// Order 订单表
type Order struct {
	ID     uint   `gorm:"primaryKey"`
	Item   string `gorm:"size:255;not null"`
	UserID uint   // 外键，指向 User
}

// Group 群组表
type Group struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"size:100;not null"`
	Users []User `gorm:"many2many:user_groups;"` // 反向关联
}
