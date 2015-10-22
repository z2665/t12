package models

import (
	"errors"
	"github.com/astaxie/beego"
	"gopkg.in/mgo.v2/bson"
)

type From struct {
	Title   string   `json:"title"`
	Content string   `json:"content"`
	Tags    []string `json:"tags"`
	//?Whohave
	BeginDay           string           `json:"beginDay"`
	EndDay             string           `json:"endDay"`
	Status             string           `json:"status"` //状态有正常，结束，过期，关闭，待审核,未通过
	TotalNeedPeopleNum int              `json:"totalNeedPeopleNum"`
	TotalPeopleNum     int              `json:"totalPeopleNum"`
	PeoPleList         []PeopleContext  `json:"peopleList"`
	CommentList        []CommentContext `json:"commentList"`
	ContactWay         ContactContext   `json:"contactWay"`
}
type CommentContext struct {
}
type PeopleContext struct {
	PeopleType     string          `json:"peopleType"`
	PeopleNumber   int             `json:"peopleNumber"`
	NeedNumber     int             `json:"needNumber"`
	PeopleResumeId []bson.ObjectId `json:"peopleResume"`
}
type ContactContext struct {
	PhoneNumber  string `json:"phoneNumber"`
	QQNumber     string `json:"QQNumber"`
	Email        string `json:"email"`
	WechatNumber string `json:"wechatNumber"`
}

//检查输入合法性。后台永远不要相信用户输入。
//待完成。多数检查还未完成
func CheckFromIsRight(f From) error {
	beego.Notice(f)
	if len(f.Title) == 0 || len(f.Content) == 0 {
		return errors.New("标题或简介不能为空")
	}
	if len(f.BeginDay) == 0 || len(f.EndDay) == 0 {
		return errors.New("起始或结束时间不能为空")
	}
	return nil
}

//将项目加入数据库
func AddFrom(f From) (err error) {
	f.Status = "正常"
	s1 := GetSession()
	defer s1.Close()
	coll := s1.DB("test").C("froms")
	err = coll.Insert(&f)
	if err != nil {
		return err
	}
	return nil
}

//根据页数获取项目列表
func GetFromList(page int) (FromList []From) {
	s1 := GetSession()
	defer s1.Close()
	coll := s1.DB("test").C("froms")
	coll.Find(nil).Sort("_id").Skip(page * 5).Limit(5).All(&FromList)
	return
}
