package model

import (
	"awesomeTestProject/initDB"
	"github.com/jinzhu/gorm"
)

type Comment struct {
	gorm.Model
	Content string
}

func (c Comment) TableName() string {
	return "comment"
}

func init() {
	table := initDB.Db.HasTable(Comment{})
	if !table {
		initDB.Db.CreateTable(Comment{})
	}
}
