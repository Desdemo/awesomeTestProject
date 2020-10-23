package model

import (
	"awesomeTestProject/initDB"
	"log"
)

type Article struct {
	Id      int    `json:"id"`
	Type    string `json:"type"`
	Content string `json:"content"`
}

func (article Article) TableName() string {
	return "article"
}

func (article Article) Insert() int {
	create := initDB.Db.Create(&article)
	if create.Error != nil {
		log.Panicln("文章添加失败", create.Error.Error())
	}
	i := create.Value
	return i.(int)
}
