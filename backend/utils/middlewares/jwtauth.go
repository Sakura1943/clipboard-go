// 中间件
// Token验证中间件

package middlewares

import (
	"backend/database"
	"backend/utils/msg"
	_token "backend/utils/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// Token验证中间件
func JWTAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 初始化返回的信息
		message := msg.Default()

		// 从请求头中获取token
		tokenSrc := ctx.Request.Header.Get("token")
		// 请求头为空，服务端发送错误定位信息
		if tokenSrc == "" {
			message.WithCode(http.StatusUnauthorized)
			message.WithType(msg.Unauthorized)
			message.WithError("No token is carried in request headers, no permission to access")
			ctx.JSON(http.StatusUnauthorized, message)
			// 终止请求的客户端的请求
			ctx.Abort()
			return
		}

		// 解析token信息携带体
		claims, err := _token.ParseToken(tokenSrc)
		// 解析token失败
		if err != nil {
			message.WithCode(http.StatusUnauthorized)
			message.WithType(msg.Unauthorized)
			message.WithError(err.Error())
			ctx.JSON(http.StatusUnauthorized, message)
			// 终止请求的客户端的请求
			ctx.Abort()
			return
		}

		// 比对当前token携带的信息与该用户在数据库中的密码是否相等

		// 初始化用户信息数据库
		db, err := database.InitUserDatabase()
		// 初始化失败
		if err != nil {
			message.WithCode(http.StatusNotAcceptable)
			message.WithType(msg.InitDatabaseError)
			message.WithError(err.Error())
			msg.PrintlnMessage(&message)
			ctx.JSON(http.StatusNotAcceptable, message)
			// 终止请求的客户端的请求
			ctx.Abort()
			return
		}

		// 获取该用户在数据库中的信息
		userInfo := database.SearchUser(db, claims.UserName)

		// 如果获取的用户名为空，则当前登录的用户不存在
		if userInfo.Name == "" {
			message.WithCode(http.StatusBadRequest)
			message.WithType(msg.UserNotExixts)
			message.WithMessage(strings.Join([]string{"The currently authenticated user `", claims.UserName, "` is not exists"}, ""))
			msg.PrintlnMessage(&message)
			ctx.JSON(http.StatusBadRequest, message)
			ctx.Abort()
			return
		}

		// 判断用户在数据库中的密码与token信息携带的密码是否相等
		if userInfo.Password != claims.UserPassword {
			message.WithCode(http.StatusNotAcceptable)
			message.WithType(msg.TokenInfoNotEqual)
			message.WithError("The incoming username and password are not equal to the username and password in the database")
			msg.PrintlnMessage(&message)
			ctx.JSON(http.StatusNotAcceptable, message)
			// 终止请求的客户端的请求
			ctx.Abort()
			return
		}

		// 登录成功，将token携带的信息存储在context(上下文)中
		ctx.Set("claims", claims)
	}
}
