package controllers

import (
	"github.com/astaxie/beego"
	"github.com/pquerna/ffjson/ffjson"
	"z2665/t12/models"
)

type FromController struct {
	beego.Controller
}

//获取添加项目页面
// @router /api/froms/changs/publish [get]
func (f *FromController) GetFromPublishPage() {
	StaticPageRender("./view/publish.html", f.Ctx.ResponseWriter)

}

//向后台添加项目
// @router /api/froms/changs/publish [post]
func (f *FromController) FromPublishByUser() {
	var from models.From
	beego.Notice(string(f.Ctx.Input.RequestBody))
	beego.Notice(f.GetString("data[title]"))
	ffjson.Unmarshal(f.Ctx.Input.RequestBody, &from)
	err := models.CheckFromIsRight(from)
	if err != nil {
		f.Data["json"] = models.ErrorContext{Err: err.Error()}
		f.ServeJson()
	} else {
		models.AddFrom(from)
		f.Data["json"] = models.ErrorContext{Data: "添加成功"}
		f.ServeJson()
	}

}

//获取项目列表 tips：page是当前页数从0开始
// @router /api/froms/lsit [get]
func (f *FromController) FromGetList() {
	page, err := f.GetInt("page")
	if err != nil {
		f.Abort("404")
	}
	list := models.GetFromList(page)
	buffer, _ := ffjson.Marshal(&list)
	f.Ctx.ResponseWriter.Write(buffer)
}
