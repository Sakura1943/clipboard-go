// 路由 /api/document/delete/:user_name/:path

package document

import (
	"backend/database"
	"backend/utils/msg"
	"backend/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// 删除文章的路由
func DeleteRoute(ctx *gin.Context) {
	// 初始化响应的消息
	message := msg.Default()

	// 获取响应的消息
	path := ctx.Query("path")

	// 获取token携带的信息
	_tokenClaims, exists := ctx.Get("claims")
	if !exists {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.VarNotExists)
		message.WithError("Failed to get token claims")
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}
	tokenClaims := _tokenClaims.(*token.Claims)

	// 初始化文章保存的数据库
	db, err := database.InitDocumentDatabase()
	if err != nil {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.InitDatabaseError)
		message.WithError(err.Error())
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 判断文章是否存在
	existsData := database.ExistsFile(db, path)
	// 如果不存在，则报错
	if !existsData {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.DataNotExists)
		message.WithError(strings.Join([]string{"Document `", path, "` already exists"}, ""))
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 获取旧的文章数据
	oldData, ok := database.SearchData(db, path)
	// 获取失败
	if !ok {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.SearchError)
		message.WithError("Document data fetch failed")
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 如果旧的文章数据中的文章所有者不要为当前登录的用户，且登录的用户不为admin，则报错
	if oldData.UserName != tokenClaims.UserName && tokenClaims.UserName != "admin" {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.OperationNotAllowed)
		message.WithError("Document deletion failed because the document does not belong to the current user")
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 删除文章
	deleted := database.DeleteData(db, path)

	// 删除失败
	if !deleted {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.DeleteDataError)
		message.WithError(strings.Join([]string{"Deleting document `", path, "` failed"}, ""))
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 删除成功
	message.WithType(msg.DeletedData)
	message.WithMessage(strings.Join([]string{"Delete document `", path, "` successfully"}, ""))
	msg.PrintlnMessage(&message)
	ctx.JSON(http.StatusBadRequest, message)
}
