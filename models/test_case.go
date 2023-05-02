package models

import "gorm.io/gorm"

type TestCase struct {
	gorm.Model
	Identity        string `gorm:"identity" json:"identity"`
	ProblemIdentity string `gorm:"problem_identity" json:"problem_identity"`
	Input           string `gorm:"input" json:"input"`
	Output          string `gorm:"output" json:"output"`
}

func (t *TestCase) TableName() string {
	return "test_case"
}
