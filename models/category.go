package models

import "gorm.io/gorm"

type Category struct {
	gorm.Model
	Identity string `gorm:"column:identity" json:"identity"`   // 唯一标识
	Name     string `gorm:"column:name" json:"name"`           // 分类名称
	ParentId string `gorm:"column:parent_id" json:"parent_id"` // 父分类ID
}

func (c *Category) TableName() string {
	return "category"
}
