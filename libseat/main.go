package main

import (
	_ "libseat/routers" //加载路由init
	_ "libseat/service" //初始化数据库和日志

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors" //允许跨域访问
	"github.com/beego/beego/logs"
)

func init() {
	//允许跨域访问
	//InsertFilter是提供一个过滤函数
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		// 允许访问所有源
		AllowAllOrigins: true,
		// 可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// 指的是允许的Header的种类
		AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// 公开的HTTP标头列表
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// 如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))
}
func init() {
	//------------------------相同日志包混用造成flag报错,故初始化在main包里--------------------------------------------//
	// beego日志文件存储方式
	logs.SetLogger(logs.AdapterFile, `{"filename":"log/log.log","level":7,"maxlines":0,"maxsize":0,"daily":true,"maxdays":10}`)
	logs.SetLogFuncCallDepth(3)
	logs.EnableFuncCallDepth(true)
	beego.BeeLogger.DelLogger("console") //只输出到文件
}

func main() {
	beego.Run()
}
