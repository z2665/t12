package main

import (
	"github.com/astaxie/beego"
	"z2665/t12/controllers"
	_ "z2665/t12/routers"
)

//加入静态文件服务目录
func main() {
	beego.SetStaticPath("/static", "./static")
	beego.Errorhandler("404", controllers.Page_not_found)
	beego.Run()

}
