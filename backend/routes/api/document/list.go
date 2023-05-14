// 路由 /api/document/list

package document

import (
	"backend/database"
	"backend/utils/msg"
	"net/http"
	"github.com/gin-gonic/gin"
)

// 列表获取的路由
func ListRoute(ctx *gin.Context) {
	// 初始化响应的消息
	message := msg.Default()
	
	// 初始化文章数据库
	db, err := database.InitDocumentDatabase()
	if err != nil {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.InitDatabaseError)
		message.WithError(err.Error())
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 获取所有文章数据
	datas := database.AllDatas(db)

	// 如果获取的文章数据长度为0，则报错
	if len(*datas) == 0 {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.Empty)
		message.WithError("Empty list")
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 获取成功
	message.WithExtra(datas)
	ctx.JSON(http.StatusOK, message)
}