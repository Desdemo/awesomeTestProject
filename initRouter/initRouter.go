package initRouter

import (
	"awesomeTestProject/handler"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func retHelloGinAndMethod(context *gin.Context)  {
	context.String(http.StatusOK,"hello gin" + strings.ToLower(context.Request.Method)+"method")
}

func SetupRouter() *gin.Engine  {
	router := gin.Default()
	router.Any("/", handler.Index)

	userGroup := router.Group("/user")
	{
		userGroup.GET("/:name", handler.UserSave)
		userGroup.POST("register")
	}


	if mode := gin.Mode(); mode == gin.TestMode {
		router.LoadHTMLGlob("./../templates/*")
	} else {
		router.LoadHTMLGlob("templates/*")
	}
	router.Static("/statics","./statics")
	router.StaticFile("/favicon.ico","./gin.ico")
	return router
}
