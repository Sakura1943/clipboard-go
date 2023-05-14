// 路由 /api/document/upload

package document

import (
	"backend/database"
	"backend/utils/msg"
	"backend/utils/random"
	"backend/utils/token"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"strings"
)

// 上传文章的路由
func UploadRoute(ctx *gin.Context) {
	// 初始化响应的消息
	message := msg.Default()

	// 获取当前登录的用户的信息
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

	// 获取用户上传的文件信息
	file, err := ctx.FormFile("file")
	if err != nil {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.UploadError)
		message.WithError("Failed to get field `file`")
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 扩展名
	ext := path.Ext(file.Filename)
	// 文件类型
	contentType := file.Header.Get("Content-Type")
	// 随机路径生成
	randomPath := random.RandStr(5)
	// 最终保存的路径
	savePath := strings.Join([]string{"tmp/", randomPath}, "")

	// 判断临时文件夹是否存在，不存在则创建
	if _, err := os.Stat("tmp/"); err != nil {
		if err := os.MkdirAll("tmp/", os.ModePerm); err != nil {
			// 创建失败，报错
			message.WithCode(http.StatusBadRequest)
			message.WithType(msg.CreateDirError)
			message.WithError(err.Error())
			msg.PrintlnMessage(&message)
			ctx.JSON(http.StatusBadRequest, message)
			return
		}
	}

	// 存储上传的文件，失败则报错
	if err := ctx.SaveUploadedFile(file, savePath); err != nil {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.SaveFileError)
		message.WithError(err.Error())
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 读取保存好的文件的内容
	data, err := os.ReadFile(savePath)
	if err != nil {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.ReadFileError)
		message.WithError(err.Error())
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 读取内容成功后删除临时文件夹，避免占用临时文件夹和数据库中信息这两份空间
	if err := os.RemoveAll("tmp/"); err != nil {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.OperationNotAllowed)
		message.WithError(err.Error())
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

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

	// 初始化响应的文章扩展名为 unknown
	extName := "unknown"

	// 如果前面获得的扩展名长度不为0，则响应的文章扩展名去掉最前面的点号
	if len(ext) != 0 {
		extName = ext[1:]
	}

	// 上传文章数据到数据库，上传失败则报错
	if _, ok := database.CreateData(db, &database.Document{
		Path:     strings.Join([]string{randomPath, "/", extName}, ""),
		Text:     string(data),
		Lang:     extName,
		UserName: tokenClaims.UserName,
	}); !ok {
		message.WithCode(http.StatusBadRequest)
		message.WithType(msg.FileExists)
		message.WithError(err.Error())
		msg.PrintlnMessage(&message)
		ctx.JSON(http.StatusBadRequest, message)
		return
	}

	// 上传成功响应
	message.WithExtra(map[string]interface{}{
		"ext":          extName,
		"content-type": contentType,
		"file-path":    strings.Join([]string{randomPath, "/", extName}, ""),
		"user":         tokenClaims.UserName,
	})
	msg.PrintlnMessage(&message)
	ctx.JSON(http.StatusOK, message)
}
