package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB = Init()

func Init() *gorm.DB {
	dsn := "root:Zhy1021@@tcp(101.42.9.7:3306)/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("数据库连接失败")
		panic(err)
	}
	return db
}
