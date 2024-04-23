package main

// @title 个人主页
// @version 1.0
// @description 用于个人文件管理
// @termsOfService http://swagger.io/terms/

// @contact.name tiny beer
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:9999
// @BasePath /

func main() {
	svc := InitializeServer()
	svc.Start()
}
