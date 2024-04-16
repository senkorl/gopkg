package xorm

import (
	"fmt"
	"time"

	"awesomeProject2/db/xorm/xorm_structs"
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

func (x *Mysql) Get(id int64) (*xorm_structs.ImSwitchStatusLog, error) {
	data := new(xorm_structs.ImSwitchStatusLog)
	_, err := x.db.Id(id).Get(data)
	return data, err
}

func (x *Mysql) Insert(data *xorm_structs.ImSwitchStatusLog) error {
	affected, err := x.db.Insert(data)
	fmt.Println("affected: " + cast.ToString(affected))
	return err
}

func (x *Mysql) Update() (int64, error) {
	param := xorm_structs.ImSwitchStatusLog{
		EndAt: time.Now().Add(2 * time.Minute),
	}
	rowsAffected, err := x.db.In("id", 2, 1, 3).Update(param)
	if err != nil {
		return 0, err
	}
	return rowsAffected, nil
}
