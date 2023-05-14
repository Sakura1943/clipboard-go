// 路由 /api/user/delete/:name

package user

import (
	"backend/database"
	"backend/utils/msg"
	"backend/utils/permission"
	"backend/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 用户删除路由的Uri参数结构体
type deleteRequestUri struct {
	// 用户名
	Name string `uri:"name" json:"name" form:"name" xml:"name" binding:"required"`
}

// 用户删除的苦于
func DeleteUserRoute(ctx *gin.Context) {
	// 初始化响应的消息
	message := msg.Default()

	// 获取参数
	var req deleteRequestUri
	if err := ctx.ShouldBindUri(&req); err != nil {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.FormDataParseError)
		message.WithError(err.Error())
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 获取JWT验证车成功的用户信息
	_tokenClaims, ok := ctx.Get("claims")
	if !ok {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.GetTokenError)
		message.WithError("")
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}
	tokenClaims := _tokenClaims.(*token.Claims)

	// 初始化用户信息数据库
	db, err := database.InitUserDatabase()
	if err != nil {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.InitDatabaseError)
		message.WithError(err.Error())
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}
	
	// 如果删除的用户在数据库中不存在，则报错
	if !database.ExistsUser(db, req.Name) {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.UserNotExixts)
		message.WithError(strings.Join([]string{"User `", req.Name, "` is not exists"}, ""))
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 获取请求删除的用户数据
	userData := database.SearchUser(db, req.Name)

	// 请求的用户为普通用户，则报错
	if tokenClaims.Permission == permission.Custom {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.PermissionDenied)
		message.WithError("Permission is denied, admin privilege is required")
		message.WithExtra(map[string]interface{}{"user_name": userData.Name, "permission": userData.Permission})
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}
	
	// 如果删除的用户名为admin，则不允许删除
	if req.Name == "admin" {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.OperationNotAllowed)
		message.WithError("Do not allow the current user to delete the admin user")
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 如果用户名等于当前登录的用户名，则不允许删除
	if req.Name == tokenClaims.UserName {
		message.WithCode(http.StatusNotAcceptable)
		message.WithType(msg.OperationNotAllowed)
		message.WithError("Do not allow the current user to delete their account")
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusNotAcceptable, message)
		return
	}

	// 删除用户
	deleted := database.DeleteUser(db, userData.Name)

	// 删除失败
	if !deleted {
		message.WithCode(http.StatusBadGateway)
		message.WithType(msg.DeleteUserError)
		message.WithError(strings.Join([]string{"Failed to delete user `", req.Name, "`"}, ""))
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 删除成功
	message.WithType(msg.DeletedUser)
	message.WithMessage(strings.Join([]string{"Delete the `", req.Name, "` user successfully"}, ""))
	msg.PrintlnMessage(&message)
	ctx.JSON(http.StatusOK, message)
}
