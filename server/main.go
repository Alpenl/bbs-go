package main

import (
	_ "bbs-go/docs" // 导入生成的 swagger 文档
	"bbs-go/internal/server"
	_ "bbs-go/internal/services/eventhandler"
)

// @title           BBS-Go API
// @version         1.0
// @description     BBS-Go 论坛系统的 API 文档
// @termsOfService  http://swagger.io/terms/
// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io
// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html
// @host      localhost:8082
// @BasePath  /api

func main() {
	server.Init()
	server.NewServer()
}
