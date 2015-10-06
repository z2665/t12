package controllers

//20150928加入登陆和注册逻辑

import (
	"fmt"
	"github.com/astaxie/beego"
	"io"
	"os"
	"strings"
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

//20151006添加添加用户头像逻辑
// @router /api/users/changs/userheadpic [post]
func (u *UserController) AddUserHeadPic() {
	picfliter := [2]string{"jpg", "png"}
	var uc models.UserCore
	var ok bool
	uc.UserName, ok = u.GetSession("longined").(string)
	if !ok {
		u.Data["json"] = models.ErrorContext{Err: "发生了未知错误"}
		u.ServeJson()
	} else {
		f, fh, err := u.GetFile("userheadpic")
		if err != nil {
			beego.Error(err.Error())
			u.Abort("404")
		}
		defer f.Close()
		arr := strings.Split(fh.Filename, ".")
		ok = false
		for _, v := range picfliter {
			if v == arr[len(arr)-1] {
				ok = true
			}
		}
		if ok {
			picname := fmt.Sprintf("%s.%s", uc.UserName, arr[len(arr)-1])
			picpath := "./static/img/UserPic/headPic/" + picname
			beego.Notice(picpath)
			uc.UserPic = picname
			fheadle, err := os.OpenFile(picpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
			if err != nil {
				beego.Error(err.Error())
			}
			defer fheadle.Close()
			io.Copy(fheadle, f)
			err = models.AddUserPic(uc)
			if err != nil {
				beego.Alert(err.Error())
				u.Data["json"] = models.ErrorContext{Err: "发生了未知错误"}
				u.ServeJson()
			} else {
				u.Data["json"] = models.ErrorContext{Data: "上传成功"}
				u.ServeJson()
			}
		}
	}
}
