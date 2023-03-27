// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameFormDefInfo = "form_def_info"

// FormDefInfo mapped from table <form_def_info>
type FormDefInfo struct {
	ID         int32     `gorm:"column:id;primaryKey" json:"id"`                                           // 主键
	Name       string    `gorm:"column:name;not null" json:"name"`                                         // 表单定义名称
	FormDefID  string    `gorm:"column:form_def_id;not null" json:"form_def_id"`                           // 表单定义id
	State      int32     `gorm:"column:state;not null" json:"state"`                                       // 状态
	CreateTime time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	CreateUser string    `gorm:"column:create_user;not null" json:"create_user"`                           // 创建人
	UpdateTime time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"`          // 更新时间
	UpdateUser string    `gorm:"column:update_user" json:"update_user"`                                    // 更新人
}

// TableName FormDefInfo's table name
func (*FormDefInfo) TableName() string {
	return TableNameFormDefInfo
}
