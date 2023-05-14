// 路由 /api/document/

package document

import (
	"backend/utils/msg"
	"net/http"
	"github.com/gin-gonic/gin"
)

// 路由组 /api/clip 根路由
func IndexRoute(ctx *gin.Context) {
	// 初始化响应的信息
	message := msg.Default()
	// 成功响应
	message.WithMessage("clipboard management")
	message.WithExtra(map[string]interface{}{
		"/api/document/upload": map[string]interface{}{
			"method":      "POST",
			"description": "Upload file",
			"use_token":   true,
		},
		"/api/document/delete/:path": map[string]interface{}{
			"method":      "DELETE",
			"description": "Delete file",
			"use_token":   true,
		},
		"/api/document/search": map[string]interface{}{
			"method":      "POST",
			"description": "Delete file",
			"use_token":   false,
		},
		"/api/document/list": map[string]interface{}{
			"method":      "GET",
			"description": "Get all list",
			"use_token":   true,
		},
	})
	msg.PrintlnMessage(&message)
	ctx.JSON(http.StatusOK, message)
}
