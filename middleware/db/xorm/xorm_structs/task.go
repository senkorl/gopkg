package xorm_structs

import "time"

// 实体
type Task struct {
	AppId                 int       `json:"app_id" xorm:"'app_id' int(11)"`
	CreatorId             int64     `json:"creator_id" xorm:"'creator_id' bigint(20)"`
	CreatorAuthId         int64     `json:"creator_auth_id" xorm:"'creator_auth_id' bigint(20)"`
	CreatorName           string    `json:"creator_name" xorm:"'creator_name' varchar(20)"`
	Id                    int64     `json:"id" xorm:"'id' bigint(20) pk autoincr"`
	Priority              int       `json:"priority" xorm:"'priority' TINYINT"`
	TypeLink              string    `json:"type_link" xorm:"'type_link' varchar(20)"`
	StaffId               int64     `json:"staff_id" xorm:"'staff_id' bigint(20)"`
	StaffName             string    `json:"staff_name" xorm:"'staff_name' varchar(20)"`
	Status                int       `json:"status" xorm:"'status' TINYINT"`
	Type                  int       `json:"type" xorm:"'type' int(11)"`
	UserSource            int       `json:"user_source" xorm:"'user_source' TINYINT"`
	TaskSource            int       `json:"task_source" xorm:"'task_source' int(11)"`
	UserId                string    `json:"user_id" xorm:"'uuid' varchar(64)"` // 注意，对应数据库是 uuid，用于兼容所有用户体系ID
	GroupId               int64     `json:"group_id" xorm:"'group_id' bigint(20)"`
	GroupName             string    `json:"group_name" xorm:"'group_name' varchar(50)"`
	TodoId                int64     `json:"todo_id" xorm:"'todo_id' bigint(20)"`
	UserMobile            string    `json:"user_mobile" xorm:"'user_mobile' varchar(20)"`
	UserName              string    `json:"user_name" xorm:"'user_name' varchar(50)"`
	Source                int       `json:"source" xorm:"'source' tinyint(4)"`
	SourceEmail           string    `json:"source_email" xorm:"'source_email' varchar(100)"` //来源邮箱
	UpdatedAt             time.Time `json:"updated_at" xorm:"'updated_at' timestamp updated"`
	CreatedAt             time.Time `json:"created_at" xorm:"'created_at' timestamp created"`
	DeletedAt             time.Time `json:"deleted_at" xorm:"'deleted_at' timestamp deleted"`
	SessionId             int64     `json:"session_id" xorm:"'session_id' bigint(20)"`
	ReminderCount         int       `json:"reminder_count" xorm:"'reminder_count' tinyint(4)"`
	StoreId               int64     `json:"store_id"  xorm:"'store_id' int(11)"`          //仓库id
	StoreName             string    `json:"store_name"  xorm:"'store_name' varchar(255)"` //仓库名称
	SubjectId             int64     `json:"subject_id"  xorm:"'subject_id' int(11)"`      //学科id
	FeedbackId            int64     `json:"feedback_id" xorm:"'feedback_id' bigint(20)"`  //反馈id
	CloseTime             time.Time `json:"close_time" xorm:"'close_time' datetime"`
	FinishTime            time.Time `json:"finish_time" xorm:"'finish_time' datetime"`
	FirstConsultationTime time.Time `json:"first_consultation_time" xorm:"'first_consultation_time' datetime"` // 首次咨询时间
}
