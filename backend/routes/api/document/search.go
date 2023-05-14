// 路由 /api/document/search

package document

import (
	"backend/database"
	"backend/utils/msg"
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
)

// 查询数据的路由的form表单结构体
type searchDataRequest struct {
	// 文章路径
	Path string `json:"path" xml:"path" form:"path" binding:"required"`
	Lang string `json:"lang" xml:"lang" form:"lang" binding:"required"`
}

// 查询文章的路由
func SearchRoute(ctx *gin.Context) {
	// 初始化响应的消息
	message := msg.Default()

	// 获取form表单的数据
	var req searchDataRequest
	if err := ctx.ShouldBind(&req); err != nil {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.FormDataParseError)
		message.WithError(err.Error())
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	queryPath := strings.Join([]string{req.Path, "/", req.Lang}, "")

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

	// 获取文章的信息
	info, ok := database.SearchData(db, queryPath)
	// 获取失败
	if !ok {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.SearchError)
		message.WithError(strings.Join([]string{"Failed to query document `", req.Path, "`"}, ""))
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 获取成功
	message.WithExtra(info)
	ctx.JSON(http.StatusOK, message)
}
