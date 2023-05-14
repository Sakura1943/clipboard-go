// 路由 /api/user/auth

package user

import (
	"backend/utils/msg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 验证当前登录的用户的token状态是否有效
func AuthRoute(ctx *gin.Context) {
	// 初始化相应的信息
	message := msg.Default()

	// 验证成功
	message.WithType(msg.Logined)
	message.WithMessage("logined!")
	msg.PrintlnMessage(&message)
	ctx.JSON(http.StatusOK, message)
}
