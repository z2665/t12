package models

import (
	"errors"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
	"strconv"
	"time"
)

//20150924创建最初版本，支持添加新用户和按名称查找
//20150925加入错误上下文结构体
//20150928加入email和输入验证以及登陆检查
//20150929优化完成使用select返回指定字段
//20150930优化数据库连接部分,加入注册用户名验证确保用户名唯一
//20151004去除一些格式判断以后再加入
//20151006加入用户头像，加入核心用户列
//20151022移除ffjson依赖
//20151112加入了完整的找回密码过程，不过存在部分问题
type User struct {
	UserName    string          `json:"userName"`
	PassWord    string          `json:"passWord"`
	TrueName    string          `json:"trueName"`
	StuNum      string          `json:"stuName"`
	Tags        []string        `json:"tags"`
	School      string          `json:"school"`
	SchoolEx    string          `json:"schoolEx"`
	Email       string          `json:"email"`
	UserPic     string          `json:"userPic"`
	IsNormal    bool            `json:"isNormal"`
	IsGreenCard bool            `json:"isGreenCard"`
	FormList    []bson.ObjectId `json:"fromList"`
	ResumeList  []bson.ObjectId `json:"resumeList"` //简历列表
	MailBox     []MailContext   `json:"mailBox"`
}

//描述一个用户需要的最少资料
type UserCore struct {
	UserName string `json:"userName"`
	UserPic  string `json:"userPic"`
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

//添加用户头像
func AddUserPic(uc UserCore) error {
	s1 := GetSession()
	defer s1.Close()
	coll := s1.DB("test").C("users")
	err := coll.Update(bson.M{"username": uc.UserName}, bson.M{"$set": bson.M{"userpic": uc.UserPic}})
	return err
}
func GetCoreUser(uc *UserCore) error {
	s1 := GetSession()
	defer s1.Close()
	coll := s1.DB("test").C("users")
	coll.Find(bson.M{"username": uc.UserName}).Select(bson.M{"username": 1, "userpic": 1}).One(uc)
	return nil
}

//通过用户名获取邮箱
func GetUserEmailByName(u *User) {
	s1 := GetSession()
	defer s1.Close()
	coll := s1.DB("test").C("users")
	coll.Find(bson.M{"username": u.UserName}).Select(bson.M{"username": 1, "email": 1}).One(u)
}

//检查用户名和邮箱是否匹配
func CheckUSerNameAndEmail(u User) error {
	var user User
	if u.Email == "" || u.UserName == "" {
		return errors.New("用户名或者邮箱为空")
	}
	user.UserName = u.UserName
	GetUserEmailByName(&user)
	if u.Email == user.Email {
		return nil
	} else {
		return errors.New("用户名和邮箱不匹配")
	}
}

//用来存放验证信息
type Pair struct {
	UserName string
	Vcode    string
	Nowtime  string
}

//存储验证信息
func SaveVcode(vcode Pair) {
	s1 := GetSession()
	defer s1.Close()
	coll := s1.DB("test").C("uvcode")
	n, _ := coll.Find(bson.M{"username": vcode.UserName}).Count()
	if n > 0 {
		//移除冗余
		coll.Remove(bson.M{"username": vcode.UserName})
	}
	coll.Insert(&vcode)
}

//检查验证信息是否过期和是否正确存在
func CheckVcode(ucode Pair) (string, error) {
	if ucode.Vcode == "" {
		return "", errors.New("效验id为空，可能是非法访问")
	}
	var vcode Pair
	s1 := GetSession()
	defer s1.Close()
	coll := s1.DB("test").C("uvcode")
	err := coll.Find(bson.M{"vcode": ucode.Vcode}).One(&vcode)
	if err != nil {
		beego.Notice(err.Error())
		return "", err
	}
	//按照本地时区解析时间
	loc, _ := time.LoadLocation("Local")
	t1, _ := time.ParseInLocation("2006-01-02 15:04:05", vcode.Nowtime, loc)
	beego.Notice(t1.Format("2006-01-02 15:04:05"))
	t2 := time.Now().Local()
	beego.Notice(t2.Format("2006-01-02 15:04:05"))
	u1 := t2.Sub(t1)
	beego.Notice(u1.String())
	//设置过期时间为5分钟
	if u1.Minutes() > 5 {
		coll.Remove(bson.M{"vcode": ucode.Vcode})
		return "", errors.New("该密匙已经过期")
	}
	beego.Notice(vcode.UserName)
	//验证完成后从数据库里移除
	coll.Remove(bson.M{"vcode": ucode.Vcode})
	return vcode.UserName, nil
}

//重置密码
func UserRestPassWord(u User) error {
	s1 := GetSession()
	defer s1.Close()
	coll := s1.DB("test").C("users")
	err := coll.Update(bson.M{"username": u.UserName}, bson.M{"$set": bson.M{"password": u.PassWord}})
	return err
}
