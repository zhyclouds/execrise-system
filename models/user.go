package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Identity string `gorm:"column:identity" json:"identity"` // 唯一标识
	Name     string `gorm:"column:name" json:"name"`         // 用户名
	Password string `gorm:"column:password" json:"password"` // 密码
	Phone    string `gorm:"column:phone" json:"phone"`       // 手机号
	Mail     string `gorm:"column:mail" json:"mail"`         // 邮箱
}

func (u *User) TableName() string {
	return "user"
}
