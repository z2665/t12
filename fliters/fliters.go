package fliters

import (
	"github.com/astaxie/beego/context"
)

//20150929加入用户登录过滤器
var UserFliter = func(ctx *context.Context) {
	_, ok := ctx.Input.Session("longined").(string)
	if !ok {
		ctx.Redirect(302, "/api/users/login")
	}
}
