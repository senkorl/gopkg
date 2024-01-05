package xorm_structs

import "time"

const TableName = "im_switch_status_log"

// 实体
type ImSwitchStatusLog struct {
	Id           int64     `json:"id" xorm:"'id' bigint(20) pk autoincr"`
	AgentId      int64     `json:"agent_id" xorm:"'agent_id' bigint(20)"`
	Type         int       `json:"type" xorm:"'type' TINYINT"`
	Status       int       `json:"status" xorm:"'status' TINYINT"`
	StartAt      time.Time `json:"start_at" xorm:"'start_at' timestamp"`
	EndAt        time.Time `json:"end_at" xorm:"'end_at' timestamp"`
	OperatorID   int64     `json:"operator_id" xorm:"'operator_id' bigint(20)"`
	OperatorName string    `json:"operator_name" xorm:"'operator_name' varchar(30)"`
	CreatedAt    time.Time `json:"created_at" xorm:"'created_at' timestamp created"`
	UpdatedAt    time.Time `json:"updated_at" xorm:"'updated_at' timestamp updated"`
}

func (m *ImSwitchStatusLog) TableName() string {
	return TableName
}
