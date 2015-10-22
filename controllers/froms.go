package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"time"
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
	tmp := f.GetString("data")
	beego.Notice(tmp)
	json.Unmarshal([]byte(tmp), &from)
	//加入当前时间
	from.BeginDay = time.Now().Format("2006-01-02 15:04:05")
	//获得当前用户
	from.MasterName, _ = f.GetSession("longined").(string)
	//检查数据正确性
	err := models.CheckFromIsRight(from)
	models.FullFrom(&from)
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
	f.Data["json"] = list
	f.ServeJson()
}
