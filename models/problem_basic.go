package models

import "gorm.io/gorm"

type ProblemBasic struct {
	gorm.Model
	Identity          string             `gorm:"column:identity" json:"identity"` // 题目标识
	ProblemCategories []*ProblemCategory `gorm:"foreignKey:problem_id;references:identity"`
	Title             string             `gorm:"column:title" json:"title"`                      // 题目标题
	Content           string             `gorm:"column:content" json:"content"`                  // 题目内容
	MaxRuntime        int                `gorm:"column:max_runtime" json:"max_runtime"`          // 最大运行时间
	MaxMem            int                `gorm:"column:max_mem" json:"max_mem"`                  // 最大运行内存
	TestCases         []*TestCase        `gorm:"foreignKey:ProblemIdentity;references:Identity"` // 测试用例
}

func (p *ProblemBasic) TableName() string {
	return "problem_basic"
}

func GetProblemList(keyword, categoryIdentity string) *gorm.DB {
	tx := DB.Model(&ProblemBasic{}).Preload("ProblemCategories").Preload("ProblemCategories.CategoryBasic").
		Where("title like ? or content like ?", "%"+keyword+"%", "%"+keyword+"%")
	if categoryIdentity != "" {
		tx = tx.Joins("right join problem_category on problem_basic.identity = problem_category.problem_id").
			Where("problem_category.category_id = (select category_basic.id from category_basic where category_basic.identity = ? )", categoryIdentity)
	}
	return tx
}
