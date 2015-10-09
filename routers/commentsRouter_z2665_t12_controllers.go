package routers

import (
	"github.com/astaxie/beego"
)

func init() {

	beego.GlobalControllerRouter["z2665/t12/controllers:FromController"] = append(beego.GlobalControllerRouter["z2665/t12/controllers:FromController"],
		beego.ControllerComments{
			"GetFromAddPage",
			`/api/froms/changs/add`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["z2665/t12/controllers:FromController"] = append(beego.GlobalControllerRouter["z2665/t12/controllers:FromController"],
		beego.ControllerComments{
			"FromAddByUser",
			`/api/froms/changs/add`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["z2665/t12/controllers:FromController"] = append(beego.GlobalControllerRouter["z2665/t12/controllers:FromController"],
		beego.ControllerComments{
			"FromGetList",
			`/api/froms/lsit`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["z2665/t12/controllers:IndexController"] = append(beego.GlobalControllerRouter["z2665/t12/controllers:IndexController"],
		beego.ControllerComments{
			"GetIndexPage",
			`/`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["z2665/t12/controllers:UserController"] = append(beego.GlobalControllerRouter["z2665/t12/controllers:UserController"],
		beego.ControllerComments{
			"GetSingnUpPage",
			`/api/users/signup`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["z2665/t12/controllers:UserController"] = append(beego.GlobalControllerRouter["z2665/t12/controllers:UserController"],
		beego.ControllerComments{
			"Adduser",
			`/api/users/signup`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["z2665/t12/controllers:UserController"] = append(beego.GlobalControllerRouter["z2665/t12/controllers:UserController"],
		beego.ControllerComments{
			"GetLoginPage",
			`/api/users/login`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["z2665/t12/controllers:UserController"] = append(beego.GlobalControllerRouter["z2665/t12/controllers:UserController"],
		beego.ControllerComments{
			"UserLogin",
			`/api/users/login`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["z2665/t12/controllers:UserController"] = append(beego.GlobalControllerRouter["z2665/t12/controllers:UserController"],
		beego.ControllerComments{
			"AddUserHeadPic",
			`/api/users/changs/userheadpic`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["z2665/t12/controllers:UserController"] = append(beego.GlobalControllerRouter["z2665/t12/controllers:UserController"],
		beego.ControllerComments{
			"GetUserCore",
			`/api/users/nowusercore`,
			[]string{"get"},
			nil})

}
