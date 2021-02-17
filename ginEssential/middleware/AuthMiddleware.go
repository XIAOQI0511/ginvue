package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"qi.xiao/ginessential/common"
	"qi.xiao/ginessential/model"
)

//中间件 用户认证，保护路由
func AuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		//获取authorization header
		tokenString := ctx.GetHeader("Authorization")
		//token为空或者不是以bearer开头
		if tokenString == "" || !strings.HasPrefix(tokenString, "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			//抛弃当次请求
			ctx.Abort()
			return
		}

		tokenString = tokenString[7:]
		token, claims, err := common.ParseToken(tokenString)
		if err != nil || !token.Valid {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			//抛弃当次请求
			ctx.Abort()
			return
		}
		//验证通过
		userId := claims.UserId
		DB := common.GetDB()
		var user model.User
		DB.First(&user, userId)

		if user.ID == 0 {
			ctx.JSON(http.StatusUnauthorized, gin.H{"code": 401, "msg": "权限不足"})
			//抛弃当次请求
			ctx.Abort()
			return
		}

		//用户存在，将user信息写入上下文
		ctx.Set("user", user)
		ctx.Next()

	}
}
