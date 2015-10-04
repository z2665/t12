package models

import (
	"errors"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"strconv"
)

//ps：所有对json结构体更改之后执行ffjson
//20150924创建最初版本，支持添加新用户和按名称查找
//20150925加入错误上下文结构体
//20150928加入email和输入验证以及登陆检查
//20150929优化完成使用select返回指定字段
//20150930优化数据库连接部分,加入注册用户名验证确保用户名唯一
//20151004去除一些格式判断以后再加入
type User struct {
	UserName    string          `json:"userName"`
	PassWord    string          `json:"passWord"`
	TrueName    string          `json:"trueName"`
	StuNum      string          `json:"stuName"`
	Tags        []string        `json:"tags"`
	School      string          `json:"school"`
	SchoolEx    string          `json:"schoolEx"`
	Email       string          `json:"email"`
	IsNormal    bool            `json:"isNormal"`
	IsGreenCard bool            `json:"isGreenCard"`
	FormList    []bson.ObjectId `json:"fromList"`
	ResumeList  []bson.ObjectId `json:"resumeList"` //简历列表
	MailBox     []MailContext   `json:"mailBox"`
}
type MailContext struct {
}
type ErrorContext struct {
	Err  string      `json:"err"`
	Data interface{} `json:"data"`
}

func AddUser(u User) error {

	beego.Notice(u)
	s1 := GetSession()
	beego.Notice(s1)
	defer s1.Close()
	coll := s1.DB("test").C("users")
	n, err := coll.Find(bson.M{"username": u.UserName}).Count()
	if err != nil {
		return err
	}
	if n > 0 {
		return errors.New("用户名已经被注册了")
	}
	err = coll.Insert(&u)
	if err != nil {
		return err
	}
	return nil
}
func CheckUserIsRight(u User) error {
	if len(u.UserName) == 0 || len(u.TrueName) == 0 {
		return errors.New("用户名或真实姓名不能为空")
	}
	if len(u.PassWord) == 0 {
		return errors.New("密码不能为空")
	}
	if len(u.StuNum) == 0 {
		return errors.New("学号不能为空")
	}
	_, err := strconv.Atoi(u.StuNum)
	if err != nil {
		return errors.New("学号必须是数字")
	}
	return nil
}
func GetUserByUserName(u *User) {
	s1 := GetSession()
	defer s1.Close()
	coll := s1.DB("test").C("users")
	coll.Find(bson.M{"username": u.UserName}).One(u)
}
func CheckUserWhenLogin(u User) error {
	s1 := GetSession()
	defer s1.Close()
	coll := s1.DB("test").C("users")
	var rightuser User
	coll.Find(bson.M{"username": u.UserName}).Select(bson.M{"username": 1, "password": 1}).One(&rightuser)
	if u.UserName == rightuser.UserName && u.PassWord == rightuser.PassWord {
		return nil
	} else {
		return errors.New("登陆失败，用户名或密码错误")
	}

}
