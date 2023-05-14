// 路由 /api/user/update

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

// 用户更新路由的参数
type updateRequestForm struct {
	// 用户ID, 默认空则为旧ID
	ID uint64 `json:"id" xml:"id" form:"id"`
	// 新用户名
	Name string `json:"name" xml:"name" form:"name" binding:"required"`
	// 旧的用户名
	OldName string `json:"old_name" xml:"old_name" form:"old_name" binding:"required"`
	// 新密码， 默认空则为旧密码
	Password string `json:"password" xml:"password" form:"password"`
	// 新权限，默认空则为旧权限
	Permission string `json:"permission" xml:"permission" form:"permission"`
}

// 用户信息更新路由
func UpdateRoute(ctx *gin.Context) {
	// 初始化返回的消息
	message := msg.Default()

	// 获取form表单
	var req updateRequestForm
	if err := ctx.ShouldBind(&req); err != nil {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.FormDataParseError)
		message.WithError(err.Error())
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 如果传入的旧用户为空时报错
	if req.OldName == "" {
		message.WithType(msg.FieldNotExists)
		message.WithCode(http.StatusNotAcceptable)
		message.WithMessage("The field `old_name` does not exist")
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusNotAcceptable, message)
		return
	}

	// 获取token携带的信息
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

	/// 判断更改的权限是不是custom和admin
	if req.Permission != permission.Custom && req.Permission != permission.Admin {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.UnknownPermission)
		message.WithError("Only admin and custom permissions are supported")
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	/// 初始化用户信息数据库
	db, err := database.InitUserDatabase()
	if err != nil {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.InitDatabaseError)
		message.WithError(err.Error())
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 旧的用户数据
	oldUserInfo := database.SearchUser(db, req.OldName)

	// 为空则为默认值
	{
		if req.ID == 0 {
			req.ID = oldUserInfo.ID
		}
		if req.Password == "" {
			req.Password = oldUserInfo.Password
		}
		if req.Permission == "" {
			req.Permission = oldUserInfo.Permission
		}
	}

	// 当前用户新信息
	newCurrentToken := database.SearchUser(db, tokenClaims.UserName)

	// Admin阶段
	{
		/// 请求的用户名是admin且修改权限为custom时报错
		if tokenClaims.UserName == "admin" && req.Name == tokenClaims.UserName && req.Permission == permission.Custom {
			message.WithCode(http.StatusBadRequest)
			message.WithType(msg.PermissionDenied)
			message.WithError("The admin user is not allowed to change permissions")
			msg.PrintlnMessage(&message)
			ctx.JSON(http.StatusBadRequest, message)
			return
		}

		// 允许admin执行
		if tokenClaims.UserName == "admin" {
			// 禁止admin用户更改用户名
			if req.OldName == "admin" && req.Name != "admin" {
				message.WithCode(http.StatusBadRequest)
				message.WithType(msg.OperationNotAllowed)
				message.WithError("The admin user is not allowed to change the username")
				msg.PrintlnMessage(&message)
				ctx.JSON(http.StatusBadRequest, message)
				return
			}

			// 更新用户信息
			newUserInfo := database.UpdateUser(db, req.OldName, &database.User{
				ID:         req.ID,
				Name:       req.Name,
				Password:   req.Password,
				Permission: req.Permission,
			})
			// 如果获取新信息失败
			if newUserInfo.Name == "" {
				message.WithCode(http.StatusBadRequest)
				message.WithType(msg.UpdateError)
				message.WithError(strings.Join([]string{"Failed to update user `", req.Name, "`'s information"}, ""))
				msg.PrintlnMessage(&message)
				ctx.JSON(http.StatusBadRequest, message)
				return
			}
			message.WithType(msg.Updated)
			message.WithMessage(strings.Join([]string{"Update the information of the user `", req.Name, "` successfully"}, ""))
			msg.PrintlnMessage(&message)
			ctx.JSON(http.StatusOK, message)
			return
		}
	}

	// 判断新ID是否存在
	if req.Name != req.OldName && database.ExistsUserById(db, req.ID) {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.OperationNotAllowed)
		message.WithError("The new user ID already exists")
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 普通用户阶段
	{
		// 请求修改的用户不等于当前用户时报错
		if req.OldName != req.Name && database.ExistsUser(db, req.Name) {
			message.WithCode(http.StatusBadRequest)
			message.WithType(msg.UserExists)
			message.WithError("The new user name already exists")
			msg.PrintlnMessage(&message)
			ctx.JSON(http.StatusBadRequest, message)
			return
		}
		// 如果旧用户名不等于当前用户名时报错
		if req.OldName != newCurrentToken.Name {
			message.WithCode(http.StatusBadRequest)
			message.WithType(msg.PermissionDenied)
			message.WithError("The current user is not allowed to change the information of other users")
			msg.PrintlnMessage(&message)
			ctx.JSON(http.StatusBadRequest, message)
			return
		}

		// 如何修改admin用户时报错
		if req.OldName == "admin" {
			message.WithCode(http.StatusBadRequest)
			message.WithType(msg.NotAllowedToChangeAdminInfo)
			message.WithError("It is not allowed to change the information of the admin user")
			msg.PrintlnMessage(&message)
			ctx.JSON(http.StatusBadRequest, message)
			return
		}

		// 如果当前用户名等于请求修改的用户名，但是当前用户权限为普通用户，却仍然要修改权限时报错
		if req.OldName == newCurrentToken.Name && newCurrentToken.Permission == permission.Custom && req.Permission == permission.Admin {
			message.WithCode(http.StatusBadRequest)
			message.WithType(msg.OperationNotAllowed)
			message.WithError("The current user is not allowed to change the current user's information")
			msg.PrintlnMessage(&message)
			ctx.JSON(http.StatusNotAcceptable, message)
			return
		}

		// 如果当前用户名等于请求修改的用户名，但是改变当前用户权限，则报错
		if req.OldName == newCurrentToken.Name && newCurrentToken.Permission != req.Permission {
			message.WithCode(http.StatusBadRequest)
			message.WithType(msg.OperationNotAllowed)
			message.WithError("The current user is not allowed to change their permissions")
			msg.PrintlnMessage(&message)
			ctx.JSON(http.StatusNotAcceptable, message)
			return
		}
	}

	// 更新用户信息，并拿取新的用户信息
	newUserInfo := database.UpdateUser(db, req.OldName, &database.User{
		ID:         req.ID,
		Name:       req.Name,
		Password:   req.Password,
		Permission: req.Permission,
	})

	// 如果获取的用户信息的名称为空，则更新失败
	if newUserInfo.Name == "" {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.UpdateError)
		message.WithError(strings.Join([]string{"Failed to update user `", req.Name, "` information"}, ""))
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 否则更新成功
	message.WithType(msg.Updated)
	message.WithMessage(strings.Join([]string{"Update user `", req.Name, "` information successful"}, ""))
	msg.PrintlnMessage(&message)
	ctx.JSON(http.StatusOK, message)
}
