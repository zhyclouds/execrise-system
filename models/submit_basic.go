package models

import "gorm.io/gorm"

type SubmitBasic struct {
	gorm.Model
	Identity        string `gorm:"column:identity" json:"identity"`                 // 提交标识
	ProblemIdentity string `gorm:"column:problem_identity" json:"problem_identity"` // 题目标识
	UserIdentity    string `gorm:"column:user_identity" json:"user_identity"`       // 用户标识
	Path            string `gorm:"column:path" json:"path"`                         // 提交路径
}

func (s *SubmitBasic) TableName() string {
	return "submit_basic"
}
