// 路由 /api/user/

package user

import (
	"backend/utils/msg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 路由组的 /api/user 的根路由
func IndexRoute(ctx *gin.Context) {
	// 初始化响应的信息
	message := msg.Default()
	// 成功响应
	message.WithMessage("user management")
	message.WithExtra(map[string]interface{}{
		"paths": map[string]interface{}{
			"/api/delete/:name": map[string]interface{}{
				"method":           "DELETE",
				"description":      "accout delete",
				"use_token":        true,
				"permission":       "admin",
				"change_self_info": false,
			},
			"/api/update": map[string]interface{}{
				"method":           "PUT",
				"description":      "update accout information",
				"use_token":        true,
				"permission":       "admin",
				"change_self_info": true,
			},
			"/api/list": map[string]interface{}{
				"method":           "GET",
				"description":      "Show all of users",
				"use_token":        true,
				"permission":       "all",
				"change_self_info": true,
			},
		},
	})
	msg.PrintlnMessage(&message)
	ctx.JSON(http.StatusOK, message)
}
