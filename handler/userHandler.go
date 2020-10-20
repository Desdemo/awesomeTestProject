package handler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UserSave(context *gin.Context)  {
	username := context.Param("name")
	context.String(http.StatusOK,"用户"+ username+"已保存")
	fmt.Println(username)
}

func UserRegister(context *gin.Context)  {
	email := context.PostForm("email")
	password := context.DefaultPostForm("password","G123456")
	passwordAgain := context.DefaultPostForm("password-again","G123456")
	fmt.Println("email", email, "password", password, "password again", passwordAgain)
}
