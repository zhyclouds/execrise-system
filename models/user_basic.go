package models

import "gorm.io/gorm"

type UserBasic struct {
	gorm.Model
	Identity         string `gorm:"column:identity" json:"identity"`                     // 唯一标识
	Name             string `gorm:"column:name" json:"name"`                             // 用户名
	Password         string `gorm:"column:password" json:"password"`                     // 密码
	Phone            string `gorm:"column:phone" json:"phone"`                           // 手机号
	Mail             string `gorm:"column:mail" json:"mail"`                             // 邮箱
	FinishProblemNum int64  `gorm:"column:finish_problem_num" json:"finish_problem_num"` // 已完成题目数量
	SubmitNum        int64  `gorm:"column:submit_num" json:"submit_num"`                 // 提交次数
	IsAdmin          int    `gorm:"column:is_admin" json:"is_admin"`                     // 是否为管理员
}

func (u *UserBasic) TableName() string {
	return "user_basic"
}
