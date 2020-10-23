package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		authHeader := context.GetHeader("Authorization")
		if authHeader == "" {
			context.JSON(http.StatusOK, gin.H{
				"code": 2003,
				"msg":  "用户验证失败",
			})
			context.Abort()
			context.HTML(http.StatusUnauthorized, "401.tmpl", nil)
			return
		} else {
			context.Next()
		}

	}
}
