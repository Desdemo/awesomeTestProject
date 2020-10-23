package model

import (
	"awesomeTestProject/initDB"
	"database/sql"
	"log"
)

type UserModel struct {
	Id            int    `form:"id"`
	Email         string `form:"email" binding:"email"`
	Password      string `form:"password"`
	PasswordAgain string `form:"password-again"`
	Avatar        sql.NullString
}

func (user UserModel) TableName() string {
	return "user"
}

func (user *UserModel) Save() int {
	result := initDB.Db.Create(&user)
	if result.Error != nil {
		log.Panicln("user insert error", result.Error.Error())
	}
	id := result.Value.(int)
	return id
}

func (user *UserModel) QueryByEmail() UserModel {
	u := UserModel{}
	result := initDB.Db.Where("email = ?", user.Email).First(&u)
	if result.Error != nil {
		log.Panicln(result.Error.Error())
	}
	return u
}

func (user *UserModel) QueryById(id int) (UserModel, error) {
	u := UserModel{}
	row := initDB.Db.First(&u, id)
	if row.Error != nil {
		log.Panicln(row.Error.Error())
	}
	return u, row.Error
}

func (user *UserModel) Update(id int) error {
	user.Id = id
	result := initDB.Db.Model(&user).Update("password", "avatar")
	if result.Error != nil {
		log.Panicln("发生了错误", result.Error.Error())
	}
	return result.Error
}
