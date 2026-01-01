package xorm

import (
	"awesomeProject/database/mysql/xorm/tables"
	"fmt"
	"log"
	"testing"
	"time"
)

func TestSync(t *testing.T) {
	xdb := New("root:root@(localhost:33060)/default?charset=utf8mb4")
	if xdb.err != nil {
		log.Fatal(xdb.err.Error())
	}
	xdb.ShowSQL()
	err := xdb.Sync()
	if err != nil {
		log.Fatal(err.Error())
	}
}

func TestInsert(t *testing.T) {
	xdb := New("root:root@(localhost:33060)/default?charset=utf8mb4")
	if xdb.err != nil {
		log.Fatal(xdb.err.Error())
	}
	data := tables.ImSwitchStatusLog{
		AgentId:      12,
		Type:         2,
		Status:       3,
		StartAt:      time.Now().Add(-1 * time.Minute),
		EndAt:        time.Now().Add(5 * time.Minute),
		OperatorID:   1,
		OperatorName: "system",
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	xdb.ShowSQL()
	err := xdb.Insert(&data)
	if err != nil {
		return
	}
	fmt.Println(data.Id)
}

func TestGet(t *testing.T) {
	xdb := New("root:root@(localhost:33060)/default?charset=utf8mb4")
	if xdb.err != nil {
		log.Fatal(xdb.err.Error())
	}
	xdb.ShowSQL()
	get := tables.ImSwitchStatusLog{}
	err := xdb.Get(1, &get)
	if err != nil {
		return
	}
	fmt.Println(get)
}

func TestUpdate(t *testing.T) {
	xdb := New("root:root@(localhost:33060)/default?charset=utf8mb4")
	if xdb.err != nil {
		log.Fatal(xdb.err.Error())
	}
	xdb.ShowSQL()
	update := tables.ImSwitchStatusLog{
		EndAt: time.Now().Add(2 * time.Minute),
	}
	rowsAffected, err := xdb.Update(&update, 2, 1, 3)
	if err != nil {
		return
	}
	fmt.Println(rowsAffected)
}
