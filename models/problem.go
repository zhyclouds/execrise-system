package models

import "gorm.io/gorm"

type Problem struct {
	gorm.Model
	Identity   string `gorm:"column:identity" json:"identity"`       // 题目标识
	CategoryId string `gorm:"column:category_id" json:"category_id"` // 题目分类
	Title      string `gorm:"column:title" json:"title"`             // 题目标题
	Content    string `gorm:"column:content" json:"content"`         // 题目内容
	MaxRuntime int    `gorm:"column:max_runtime" json:"max_runtime"` // 最大运行时间
	MaxMem     int    `gorm:"column:max_mem" json:"max_mem"`         // 最大运行内存
}

func (p *Problem) TableName() string {
	return "problem"
}
