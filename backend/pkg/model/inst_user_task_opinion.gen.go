// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameInstUserTaskOpinion = "inst_user_task_opinion"

// InstUserTaskOpinion mapped from table <inst_user_task_opinion>
type InstUserTaskOpinion struct {
	ID          int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                        // 唯一id
	InstTaskID  string    `gorm:"column:inst_task_id;not null" json:"inst_task_id"`                         // 实例任务id
	NodeTaskID  string    `gorm:"column:node_task_id;not null" json:"node_task_id"`                         // 节点任务id
	UserTaskID  string    `gorm:"column:user_task_id;not null" json:"user_task_id"`                         // 用户任务id
	NodeID      string    `gorm:"column:node_id;not null" json:"node_id"`                                   // 节点id
	OpinionID   string    `gorm:"column:opinion_id;not null" json:"opinion_id"`                             // 意见id
	Opinion     int32     `gorm:"column:opinion;not null;default:1" json:"opinion"`                         // 处理意见【1：未处理；2：已阅；3：同意；4：不同意；5：回退；6：终止】
	OpinionDesc string    `gorm:"column:opinion_desc;not null" json:"opinion_desc"`                         // 处理意见描述
	OpUserID    string    `gorm:"column:op_user_id;not null" json:"op_user_id"`                             // 操作用户id
	OpUserName  string    `gorm:"column:op_user_name;not null" json:"op_user_name"`                         // 操作用户名称
	CreateTime  time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	UpdateTime  time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"`          // 更新时间
	OpinionTime time.Time `gorm:"column:opinion_time;default:CURRENT_TIMESTAMP" json:"opinion_time"`        // 发表意见时间
}

// TableName InstUserTaskOpinion's table name
func (*InstUserTaskOpinion) TableName() string {
	return TableNameInstUserTaskOpinion
}
