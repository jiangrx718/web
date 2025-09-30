package model

import (
	"time"
	"web/gopkg/gorms"
)

type Demo struct {
	gorms.Entity
	DemoId      string    `gorm:"column:demo_id;type:varchar(32);index:uni_demo_id,unique;comment:案例ID"`
	Name        string    `gorm:"column:name;type:varchar(1024);default:'';comment:案例标题"`
	FileType    int       `gorm:"column:file_type;type:integer;default:0;comment:文档类型,1可研报告,2需求报告"`
	ProjectType int       `gorm:"column:project_type;type:integer;default:0;comment:项目类型,1产品购置类,2开发实施类,3数据工程类,4咨询、运维类"`
	Content     string    `gorm:"column:content;type:varchar(1024);default:'';comment:文档内容"`
	CreatedAt   time.Time `gorm:"column:created_at;type:time;autoCreateTime;index:idx_demo_created_at;comment:创建时间"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:time;autoUpdateTime;index:idx_demo_updated_at;comment:更新时间"`
}

func (r *Demo) TableName() string {
	return "demo"
}
