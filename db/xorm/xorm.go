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
	err error
}

func New(dsn string) *Mysql {
	engine, err := xorm.NewEngine("mysql", dsn)
	return &Mysql{db: engine, err: err}
}

func (x *Mysql) Sync() error {
	return x.db.Sync(new(xorm_structs.ImSwitchStatusLog))
}

func (x *Mysql) Get(id int64) (*xorm_structs.ImSwitchStatusLog, error) {
	x.db.ShowSQL(true)
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
	updateInfo := map[string]interface{}{
		"start_at": time.Now(),
	}
	id, err := x.db.Table("im_switch_status_log").
		In("id", 1).
		Update(updateInfo)
	if err != nil {
		return 0, err
	}
	return id, nil
}
