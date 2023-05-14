// 路由 /api/user/register

package user

import (
	"backend/database"
	"backend/utils/msg"
	"backend/utils/permission"
	"backend/utils/token"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// 用户注册路由的form表单结构体
type registerRequestForm struct {
	// 用户名
	Name string `json:"name" form:"name" xml:"name" binding:"required"`
	// 密码
	Password string `json:"password" form:"password" xml:"password"`
	// 权限
	Permission string `json:"permission" form:"permission" xml:"permission"`
}

// 用户信息注册路由
func RegisterRoute(ctx *gin.Context) {
	// 初始化响应的消息
	message := msg.Default()

	// 获取token中携带的信息
	_tokenClaims, ok := ctx.Get("claims")
	if !ok {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.GetTokenError)
		message.WithError("Failed to get token claims")
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}
	tokenClaims := _tokenClaims.(*token.Claims)

	// 仅允许注册用户
	if tokenClaims.UserName != "admin" {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.OperationNotAllowed)
		message.WithError("Only the admin user is allowed to register new users")
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 获取form表单
	var req registerRequestForm
	if err := ctx.ShouldBind(&req); err != nil {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.FormDataParseError)
		message.WithError(err.Error())
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}
	// 如何密码和权限为空，则取默认值
	{
		if req.Password == "" {
			req.Password = "12345678"
			message.WithMessage("Default password: 12345678, default permission: custom")
		}
		if req.Permission == "" {
			req.Permission = permission.Custom
		}
	}

	// 初始化用户信息数据库
	db, err := database.InitUserDatabase()
	if err != nil {
		message.WithError(err.Error())
		message.WithType(msg.InitDatabaseError)
		message.WithCode(http.StatusBadRequest)
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 如果用户已存在数据库中，则报错
	if database.ExistsUser(db, req.Name) {
		message.WithType(msg.UserExists)
		message.WithCode(http.StatusBadRequest)
		message.WithError(strings.Join([]string{"User `", req.Name, "` is exists"}, ""))
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 向数据库中添加新数据
	if _, ok := database.CreateUser(db, &database.User{
		Name:       req.Name,
		Password:   req.Password,
		Permission: req.Permission,
	}); !ok {
		message.WithCode(http.StatusNotAcceptable)
		message.WithType(msg.CreatingUserFailed)
		message.WithError(strings.Join([]string{"Creating user `", req.Name, "` failed"}, ""))
		ctx.JSON(http.StatusNotAcceptable, message)
		return
	}

	// 成功注册
	message.WithType(msg.Registered)
	message.WithExtra(map[string]interface{}{"user_name": req.Name})
	msg.PrintlnMessage(&message)
	ctx.JSON(http.StatusOK, message)
}
