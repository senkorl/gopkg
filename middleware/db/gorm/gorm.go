package gorm

import (
	"database/sql"

	"awesomeProject/middleware/db/gorm/gorm_structs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	db  *gorm.DB
	dsn string
	err error
}

func New(dsn string) *Mysql {
	sqlDB, err := sql.Open("mysql", dsn)
	if err != nil {
		return &Mysql{err: err}
	}
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})
	return &Mysql{db: gormDB, dsn: dsn, err: err}
}

func (g *Mysql) Create(param *gorm_structs.UserEnt) *gorm_structs.UserEnt {
	g.db.Create(param)
	return param
}
