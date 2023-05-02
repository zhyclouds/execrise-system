package models

import "gorm.io/gorm"

type ProblemBasic struct {
	gorm.Model
	Identity   string      `gorm:"column:identity" json:"identity"`       // 题目标识
	Title      string      `gorm:"column:title" json:"title"`             // 题目标题
	Content    string      `gorm:"column:content" json:"content"`         // 题目内容
	MaxRuntime int         `gorm:"column:max_runtime" json:"max_runtime"` // 最大运行时间
	MaxMem     int         `gorm:"column:max_mem" json:"max_mem"`         // 最大运行内存
	TestCases  []*TestCase `gorm:"foreignKey:ProblemIdentity;references:Identity"`
}

func (p *ProblemBasic) TableName() string {
	return "problem_basic"
}

func GetProblemList(keyword string) *gorm.DB {
	return DB.Model(&ProblemBasic{}).Where("title like ? or content like ?", "%"+keyword+"%", "%"+keyword+"%")
}
