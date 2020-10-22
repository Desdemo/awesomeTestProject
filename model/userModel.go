package model

import (
	"awesomeTestProject/initDB"
	"database/sql"
	"log"
)

type UserModel struct {
	ID            int
	Email         string `form:"email" binding:"email"`
	Password      string `form:"password"`
	PasswordAgain string `form:"password-again"`
	Avatar        sql.NullString
}

func (user *UserModel) Save() int64 {
	result, err := initDB.Db.Exec("insert into ginhello.user (email,password) values (?,?)", user.Email, user.Password)
	if err != nil {
		log.Panicln("user insert error", err.Error())
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Panicln("user insert id error", err.Error())
	}
	return id
}

func (user *UserModel) QueryByEmail() UserModel {
	u := UserModel{}
	row := initDB.Db.QueryRow("select * from user where email = ?;", user.Email)
	e := row.Scan(&u.ID, &u.Email, &u.Password)
	if e != nil {
		log.Panicln(e)
	}
	return u
}

func (user *UserModel) QueryById(id int) (UserModel, error) {
	u := UserModel{}
	row := initDB.Db.QueryRow("select * from user where id = ?;", id)
	e := row.Scan(&u.ID, &u.Email, &u.Password, &u.Avatar)
	if e != nil {
		log.Panicln(e)
	}
	return u, e
}

func (user *UserModel) Update(id int) error {
	var stmt, e = initDB.Db.Prepare("update user set password=?,avatar=? where id = ?")
	if e != nil {
		log.Panicln("发生了错误", e.Error())
	}
	_, e = stmt.Exec(user.Password, user.Avatar.Valid, user.ID)
	if e != nil {
		log.Panicln("错误", e.Error())
	}
	return e
}
