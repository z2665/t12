package routers

import (
	"github.com/astaxie/beego"
	"z2665/t12/controllers"
	"z2665/t12/fliters"
)

//20150928移除命名空间只使用注解路由
//20150929将表单控制器加入路由，将过滤器加入路由
//20151004加入index路由
func init() {
	beego.InsertFilter("/api/froms/changs/*", beego.BeforeRouter, fliters.UserFliter)
	beego.InsertFilter("/api/users/changs/*", beego.BeforeRouter, fliters.UserFliter)
	beego.Include(&controllers.UserController{})
	beego.Include(&controllers.FromController{})
	beego.Include(&controllers.IndexController{})

}
