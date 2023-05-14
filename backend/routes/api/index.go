// 路由 /api/

package api

import (
	"backend/utils/msg"
	"net/http"
	"github.com/gin-gonic/gin"
)

// 路由组 /api 根路由
func IndexRoute(ctx *gin.Context) {
	// 初始化响应的消息
	message := msg.Default()
	// 响应成功的信息
	message.WithMessage("/api routes")
	message.WithExtra(map[string]interface{}{
		"/api/user": map[string]interface{}{
			"description":      "user management",
			"use_token":        true,
			"permission":       "all",
			"change_self_info": true,
		},
		"/api/document": map[string]interface{}{
			"description": "document data management",
			"use_token":   true,
			"permission":  "all",
		},
	})
	msg.PrintlnMessage(&message)
	ctx.JSON(http.StatusOK, message)
}
