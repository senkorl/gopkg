package db

import (
	"fmt"

	"awesomeProject2/db/xorm_structs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func SyncSql() {
	engine, err := xorm.NewEngine("mysql",
		"qdfresh_test:ZYF5gThaUNdf4QV6@(qa-test-m-out.mysql.zhangbei.rds.aliyuncs.com)/yr_box_im?charset=utf8mb4")
	if err != nil {
		fmt.Println(err.Error())
	}
	err = engine.Sync(new(xorm_structs.ImSwitchStatusLog))
	if err != nil {
		fmt.Println(err.Error())
	}
}
