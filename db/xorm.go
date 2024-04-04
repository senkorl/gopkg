package db

import (
	"fmt"
	"log"
	"time"

	xorm_structs2 "awesomeProject2/db/xorm_structs"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func SyncSql() {
	engine, err := xorm.NewEngine("mysql",
		"qdfresh_test:ZYF5gThaUNdf4QV6@(qa-test-m-out.mysql.zhangbei.rds.aliyuncs.com)/yr_box_im?charset=utf8mb4")
	if err != nil {
		log.Fatal(err.Error())
	}
	err = engine.Sync(new(xorm_structs2.ImSwitchStatusLog))
	if err != nil {
		log.Fatal(err.Error())
	}
}

func UpdateTimeSql() {
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
		log.Fatal(err.Error())
	}
	fmt.Println(id)
}

func GetTaskSql() {
	engine, err := xorm.NewEngine("mysql",
		"qdfresh_test:ZYF5gThaUNdf4QV6@(qa-test-m-out.mysql.zhangbei.rds.aliyuncs.com)/lme_crm_workorder?charset=utf8mb4")
	if err != nil {
		fmt.Println(err.Error())
	}

	engine.ShowSQL(true)
	task := xorm_structs2.Task{}
	_, err = engine.Table("tasks").Where("id = ?", 1011).Get(&task)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(task)
}
