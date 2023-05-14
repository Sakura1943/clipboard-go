// 路由 /api/login

package api

import (
	"backend/database"
	"backend/utils/msg"
	"backend/utils/token"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 用户登录路由form表单结构体
type loginRequestForm struct {
	// 用户名
	Name string `json:"name" form:"name" xml:"name" binding:"required"`
	// 密码
	Password string `json:"password" form:"password" xml:"password" binding:"required"`
}

// 用户登录路由
func LoginRoute(ctx *gin.Context) {
	// 初始化响应的消息
	message := msg.Default()

	// 获取form表单数据
	var req loginRequestForm
	if err := ctx.ShouldBind(&req); err != nil {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.FormDataParseError)
		message.WithError(err.Error())
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 初始化用户信息数据库
	db, err := database.InitUserDatabase()
	if err != nil {
		message.WithCode(http.StatusNotAcceptable)
		message.WithType(msg.GetUserInfoError)
		message.WithError("Failed to get username and password")
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusNotAcceptable, message)
		return
	}

	// 获取用户在用户数据库里的信息
	userInfo := database.SearchUser(db, req.Name)

	// 如果用户数据库里的用户密码和用户名为空，则不存在，报错
	if userInfo.Name == "" && userInfo.Password == "" {
		message.WithCode(http.StatusNotAcceptable)
		message.WithType(msg.GetUserInfoError)
		message.WithError("Failed to get user information from database")
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusNotAcceptable, message)
		return
	}

	// 如果用户数据库里的用户密码和表单里的用户密码不一致，则报错
	if userInfo.Password != req.Password {
		message.WithCode(http.StatusNotAcceptable)
		message.WithType(msg.PasswordWrong)
		message.WithError("The password is incorrect")
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusNotAcceptable, message)
		return
	}

	// 获取token字符串
	tokenStr, err := token.GenerateToken(userInfo.Name, userInfo.Password, userInfo.Permission)

	// 获取token失败
	if err != nil {
		message.WithType(msg.GenerateTokenError)
		message.WithCode(http.StatusBadRequest)
		message.WithError(err.Error())
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 登录成功
	message.WithType(msg.Logined)
	message.WithExtra(map[string]interface{}{"name": userInfo.Name, "role": userInfo.Permission, "token": tokenStr})
	msg.PrintlnMessage(&message)
	ctx.JSON(http.StatusOK, message)
}
