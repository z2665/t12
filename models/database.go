package models

import (
	"gopkg.in/mgo.v2"
)

var (
	session  *mgo.Session
	url      string = "127.0.0.1:27017"
	basename string = "test"
)

func GetSession() *mgo.Session {
	if session == nil {
		var err error
		session, err = mgo.Dial(url)
		if err != nil {
			panic(err)
		}
	}
	return session.Clone()
}
