// 路由 /

package routes

import (
	"backend/utils/msg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 根路径路由
func IndexRoute(ctx *gin.Context) {
	// 初始化响应的消息
	message := msg.Default()
	message.WithMessage("Online clipboard written using Golang")
	msg.PrintlnMessage(&message)
	ctx.JSON(http.StatusOK, message)
}
