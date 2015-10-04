package controllers

import (
	"github.com/astaxie/beego"
)

//20151004加入index主页
type IndexController struct {
	beego.Controller
}

// @router / [get]
func (i *IndexController) GetIndexPage() {
	StaticPageRender("./view/index.html", i.Ctx.ResponseWriter)
}
