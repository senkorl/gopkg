package db

import (
	"fmt"
	"time"

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

func TimeSql() {
	engine, err := xorm.NewEngine("mysql",
		"qdfresh_test:ZYF5gThaUNdf4QV6@(qa-test-m-out.mysql.zhangbei.rds.aliyuncs.com)/yr_box_im?charset=utf8mb4")
	if err != nil {
		fmt.Println(err.Error())
	}
	updateInfo := map[string]interface{}{
		"start_at": time.Now(),
	}
	id, err := engine.Table("im_switch_status_log").In("id", 1).Update(updateInfo)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(id)
}
