package test

import (
	"execrise-system/models"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"testing"
)

func TestGormTest(t *testing.T) {
	dsn := "root:Zhy1021@@tcp(101.42.9.7:3306)/gin_gorm_oj?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		t.Fatal(err)
	}
	data := make([]*models.Problem, 0)
	err = db.Find(&data).Error
	if err != nil {
		t.Fatal(err)
	}

	for _, v := range data {
		fmt.Println(v)
	}
}
