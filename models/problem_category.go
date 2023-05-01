package models

import "gorm.io/gorm"

type ProblemCategory struct {
	gorm.Model
	ProblemId     uint           `gorm:"column:problem_id" json:"problem_id"`   // 题目ID
	CategoryId    uint           `gorm:"column:category_id" json:"category_id"` // 分类ID
	CategoryBasic *CategoryBasic `gorm:"foreignKey:CategoryId"`
}
