// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameModelDetail = "model_detail"

// ModelDetail mapped from table <model_detail>
type ModelDetail struct {
	ID           int64     `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`                        // 唯一id
	ModelID      string    `gorm:"column:model_id;not null" json:"model_id"`                                 // 模板id
	ModelName    string    `gorm:"column:model_name;not null" json:"model_name"`                             // 模板名称
	ModelTitle   string    `gorm:"column:model_title;not null" json:"model_title"`                           // 模板标题
	FlowDefID    string    `gorm:"column:flow_def_id;not null" json:"flow_def_id"`                           // 流程定义id
	FormDefID    string    `gorm:"column:form_def_id;not null" json:"form_def_id"`                           // 表单定义id
	ModelGroupID string    `gorm:"column:model_group_id;not null" json:"model_group_id"`                     // 模版组id
	Status       int32     `gorm:"column:status;not null;default:1" json:"status"`                           // 模板状态【1：草稿；2：发布；3：停用】
	Remark       string    `gorm:"column:remark;not null" json:"remark"`                                     // 描述
	CreateTime   time.Time `gorm:"column:create_time;not null;default:CURRENT_TIMESTAMP" json:"create_time"` // 创建时间
	CreateUser   string    `gorm:"column:create_user;not null" json:"create_user"`                           // 创建人
	UpdateTime   time.Time `gorm:"column:update_time;default:CURRENT_TIMESTAMP" json:"update_time"`          // 更新时间
	UpdateUser   string    `gorm:"column:update_user" json:"update_user"`                                    // 更新人
	NoticeURL    string    `gorm:"column:notice_url;not null" json:"notice_url"`                             // 回调通知推送url
	TitleProps   string    `gorm:"column:title_props;not null" json:"title_props"`                           // 标题配置
}

// TableName ModelDetail's table name
func (*ModelDetail) TableName() string {
	return TableNameModelDetail
}