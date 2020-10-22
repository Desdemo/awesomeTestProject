package handler

import (
	"awesomeTestProject/model"
	"awesomeTestProject/utils"
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

func UserSave(context *gin.Context) {
	username := context.Param("name")
	context.String(http.StatusOK, "用户"+username+"已保存")
	fmt.Println(username)
}

func UserRegister(context *gin.Context) {
	var user model.UserModel
	if err := context.ShouldBind(&user); err != nil {
		fmt.Println("err - >", err.Error())
		context.String(http.StatusBadRequest, "输入的数据不合法")

	} else {
		id := user.Save()
		log.Println("id is ", id)
		context.Redirect(http.StatusMovedPermanently, "/")
	}
}

func UserLogin(context *gin.Context) {
	var user model.UserModel
	if e := context.Bind(&user); e != nil {
		log.Panicln("login 绑定错误", e.Error())
	}
	u := user.QueryByEmail()
	if u.Password == user.Password {
		log.Println("登录成功", u.Email)
		context.HTML(http.StatusOK, "index.tmpl", gin.H{
			"email": u.Email,
		})
	}
}

func UserProfile(context *gin.Context) {
	id := context.Query("id")
	var user model.UserModel
	i, err := strconv.Atoi(id)
	u, e := user.QueryById(i)
	if e != nil || err != nil {
		context.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
	}
	context.HTML(http.StatusOK, "user_profile.tmpl", gin.H{
		"user": u,
	})
}

func UpdateUserProfile(ctx *gin.Context) {
	var u model.UserModel
	if err := ctx.ShouldBind(u); err != nil {
		ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": err.Error(),
		})
		log.Panicln("绑定发生错误", err.Error())
	}
	file, e := ctx.FormFile("avatar-file")
	if e != nil {
		ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("文件上传错误", e.Error())
	}
	path := utils.RootPath()
	path = filepath.Join(path, "avatar")
	e = os.MkdirAll(path, os.ModePerm)
	if e != nil {
		ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("无法创建文件夹", e.Error())
	}
	fileName := strconv.FormatInt(time.Now().Unix(), 10) + file.Filename
	e = ctx.SaveUploadedFile(file, path+fileName)
	if e != nil {
		ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
		log.Panicln("无法保存文件", e.Error())
	}

	avatarUrl := "http://localhost:8080/avatar/" + fileName
	u.Avatar = sql.NullString{String: avatarUrl}
	e = u.Update(u.ID)
	if e != nil {
		ctx.HTML(http.StatusOK, "error.tmpl", gin.H{
			"error": e,
		})
	}
	log.Panicln("数据无法更新", e.Error())
	ctx.Redirect(http.StatusMovedPermanently, "/user/profile?id="+strconv.Itoa(u.ID))
}
