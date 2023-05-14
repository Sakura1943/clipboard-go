// 路由 /api/user/list
package user

import (
	"backend/database"
	"backend/utils/msg"
	"github.com/gin-gonic/gin"
	"net/http"
)

// list路由响应数据结构体
type listReponseData struct {
	// ID
	ID         uint64 `json:"id" form:"id" xml:"id"`
	// 用户名
	Name       string `json:"name" form:"name" xml:"name"`
	// 权限
	Permission string `json:"permission" form:"permission" xml:"permission"`
}

// 用户信息列表路由
func ListRoute(ctx *gin.Context) {
	// 初始化响应的信息
	message := msg.Default()

	// 初始化用户信息数据库
	db, err := database.InitUserDatabase()
	if err != nil {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.InitDatabaseError)
		message.WithError(err.Error())
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 获取全部用户信息
	users := database.AllUsers(db)

	// 如果用户信息数据长度为0，则为找不到用户，报错
	if len(*users) == 0 {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.Empty)
		message.WithError("Empty list")
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 初始化用响应的户信息列表
	var res []listReponseData

	// 将用户信息添加进列表
	for _, user := range *users {
		res = append(res, listReponseData{
			ID:         user.ID,
			Name:       user.Name,
			Permission: user.Permission,
		})
	}

	// 成功获取
	message.WithExtra(res)
	ctx.JSON(http.StatusOK, message)
}
