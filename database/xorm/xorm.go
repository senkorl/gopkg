package xorm

import (
	"awesomeProject/database/xorm/xorm_structs"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/spf13/cast"
)

type Mysql struct {
	db  *xorm.Engine
	dsn string
	err error
}

func New(dsn string) *Mysql {
	engine, err := xorm.NewEngine("mysql", dsn)
	return &Mysql{db: engine, dsn: dsn, err: err}
}

func (x *Mysql) ShowSQL() {
	x.db.ShowSQL(true)
}

func (x *Mysql) Sync() error {
	return x.db.Sync(new(xorm_structs.ImSwitchStatusLog))
}

func (x *Mysql) Get(id int64, data interface{}) error {
	_, err := x.db.Id(id).Get(data)
	return err
}

func (x *Mysql) Insert(data interface{}) error {
	affected, err := x.db.Insert(data)
	fmt.Println("affected: " + cast.ToString(affected))
	return err
}

func (x *Mysql) Update(update interface{}, ids ...int64) (int64, error) {
	rowsAffected, err := x.db.In("id", ids).Update(update)
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
