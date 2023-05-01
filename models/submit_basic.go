package models

import "gorm.io/gorm"

type SubmitBasic struct {
	gorm.Model
	Identity        string        `gorm:"column:identity" json:"identity"`                 // 提交标识
	ProblemIdentity string        `gorm:"column:problem_identity" json:"problem_identity"` // 题目标识
	ProblemBasic    *ProblemBasic `gorm:"foreignKey:identity;references:problem_identity"`
	UserIdentity    string        `gorm:"column:user_identity" json:"user_identity"` // 用户标识
	UserBasic       *UserBasic    `gorm:"foreignKey:identity;references:user_identity"`
	Path            string        `gorm:"column:path" json:"path"`     // 提交路径
	Status          int           `gorm:"column:status" json:"status"` // 提交状态
}

func (s *SubmitBasic) TableName() string {
	return "submit_basic"
}

func GetSubmitList(problemIdentity, userIdentity string, status int) *gorm.DB {
	tx := DB.Model(&SubmitBasic{}).Preload("ProblemBasic").Preload("UserBasic")
	if problemIdentity != "" {
		tx.Where("problem_identity = ?", problemIdentity)
	}
	if userIdentity != "" {
		tx.Where("user_identity = ?", userIdentity)
	}
	if status != 0 {
		tx.Where("status = ?", status)
	}
	return tx
}
