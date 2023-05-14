// 路由 /ping

package routes

import (
	"backend/utils/msg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 服务端与客户端连接状态验证路由
func PingRoute(ctx *gin.Context) {
	// 初始化响应的消息
	message := msg.Default()
	message.WithMessage("pong")
	msg.PrintlnMessage(&message)
	ctx.JSON(http.StatusOK, message)
}
