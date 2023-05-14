package main

import (
	"backend/database"
	"backend/routes"
	"backend/routes/api"
	"backend/routes/api/document"
	"backend/routes/api/user"
	"backend/utils/middlewares"
	"github.com/BurntSushi/toml"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

type Config struct {
	Base `toml:"base" binding:"required"`
}

type Base struct {
	AllowedOrigins []string `toml:"allowed_origins" binding:"required"`
	ServerPort     uint16   `toml:"server_port" binding:"required"`
	ServerHost     string   `toml:"server_host"`
	GinMode        string   `toml:"gin_mode"`
}

func main() {
	// 初始化数据库
	database.InitUserDatabase()
	database.InitDocumentDatabase()

	// 加载配置文件
	var config Config
	if _, err := toml.DecodeFile("config.toml", &config); err != nil {
		log.Fatalln(err)
		return
	}

	// 从env中获取地址信息
	address := strings.Join([]string{config.ServerHost, ":", strconv.Itoa(int(config.ServerPort))}, "")

	// 为空则默认模式为debug
	if config.GinMode == "" {
		config.GinMode = "debug"
	}
	// 设置GIN的运行模式
	gin.SetMode(config.GinMode)

	// 配置跨域请求
	cors := cors.New(cors.Config{
		AllowOrigins:     config.AllowedOrigins,
		AllowMethods:     []string{"GET", "POST", "DELETE", "PUT"},
		AllowHeaders:     []string{"token", "Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	})

	// 日志打印与存储，以日期为文件名存储日志(一天一个文件)
	{
		// 不存在日志保存文件夹则创建
		if _, err := os.Stat("logs"); err != nil {
			err := os.MkdirAll("logs", os.ModePerm)
			if err != nil {
				// 创建失败则panic
				log.Fatalln(err)
				return
			}
		}

		// 禁止彩色日志打印
		gin.DisableConsoleColor()
		// 文件名
		fileName := strings.Join([]string{"logs/", time.Now().Format("2006-01-02"), ".log"}, "")
		// 创建文件(不存在则创建，存在则追加数据)
		file, _ := os.OpenFile(fileName, os.O_CREATE|os.O_APPEND|os.O_RDWR, os.ModePerm)
		// 日志设置输出位置(stdout和日志文件)
		output := io.MultiWriter(file, os.Stdout)
		// 设置GIN的默认日志Writer为output
		gin.DefaultWriter = output
		// 设置log库的输出为output
		log.SetOutput(output)
	}

	// 路由管理
	{
		// 初始化默认理由(携带日志打印和panic捕获)
		router := gin.Default()

		router.Use(cors)

		// 根目录路由
		router.GET("/", routes.IndexRoute)

		// 服务器连接测试理由
		router.GET("/ping", routes.PingRoute)

		// api 路由组
		apiGroup := router.Group("/api")
		{
			apiGroup.GET("/", api.IndexRoute)
			apiGroup.POST("/login", api.LoginRoute)
			apiGroup.POST("/document/search", document.SearchRoute)

			// 用户管理api路由组
			userApiGroup := apiGroup.Group("/user")
			// 使用JWTAuth中间间验证用户是否已登录
			userApiGroup.Use(middlewares.JWTAuth())
			{
				userApiGroup.POST("/register", user.RegisterRoute)
				userApiGroup.GET("/", user.IndexRoute)
				userApiGroup.GET("/auth_confirm", user.AuthRoute)
				userApiGroup.DELETE("/delete/:name", user.DeleteUserRoute)
				userApiGroup.PUT("/update", user.UpdateRoute)
				userApiGroup.GET("/list", user.ListRoute)
			}

			// 剪切板管理api路由组
			documentApiGroup := apiGroup.Group("/document")
			// 使用JWTAuth中间间验证用户是否已登录
			documentApiGroup.Use(middlewares.JWTAuth())
			{
				documentApiGroup.POST("/upload", document.UploadRoute)
				documentApiGroup.GET("/", document.IndexRoute)
				documentApiGroup.GET("/list", document.ListRoute)
				documentApiGroup.DELETE("/delete", document.DeleteRoute)
			}
		}
		// 打印服务已开启的信息
		log.Println(strings.Join([]string{"Listening and serving HTTP on ", address, "\n"}, ""))
		// 运行服务
		router.Run(address)
	}
}
