package controllers

//20150928加入登陆和注册逻辑

import (
	"github.com/astaxie/beego"
	"z2665/t12/models"
)

type UserController struct {
	beego.Controller
}

// @router /api/users/signup [get]
func (u *UserController) GetSingnUpPage() {
	StaticPageRender("./view/signup.html", u.Ctx.ResponseWriter)
}

// @router /api/users/signup [post]
func (u *UserController) Adduser() {
	var user models.User
	user.UserName = u.GetString("username")
	user.PassWord = u.GetString("password")
	user.TrueName = u.GetString("truename")
	user.StuNum = u.GetString("stunum")
	user.School = u.GetString("school")
	user.SchoolEx = u.GetString("schoolex")
	user.Email = u.GetString("email")
	beego.Notice(user)
	err := models.CheckUserIsRight(user)
	if err != nil {
		u.Data["json"] = models.ErrorContext{Err: err.Error()}
		u.ServeJson()
	} else {
		err = models.AddUser(user)
		if err != nil {
			u.Data["json"] = models.ErrorContext{Err: err.Error()}
			u.ServeJson()
		} else {
			u.Data["json"] = models.ErrorContext{Data: "注册成功"}
			u.ServeJson()
		}

	}
}

// @router /api/users/login [get]
func (u *UserController) GetLoginPage() {
	StaticPageRender("./view/login.html", u.Ctx.ResponseWriter)
}

// @router /api/users/login [post]
func (u *UserController) UserLogin() {
	var user models.User
	user.UserName = u.GetString("username")
	user.PassWord = u.GetString("password")
	beego.Notice(user)
	err := models.CheckUserWhenLogin(user)
	if err != nil {
		u.Data["json"] = models.ErrorContext{Err: err.Error()}
		u.ServeJson()
	} else {
		u.SetSession("longined", user.UserName)
		u.Data["json"] = models.ErrorContext{Data: "登陆成功"}
		u.ServeJson()
	}
}
