package controllers

import (
	"../common"
	"gopkg.in/mgo.v2"
)
type Context struct {
	MongoSession *mgo.Session
	User         string
}

func NewContext() *Context {
	session := common.GetSession().Copy()
	context := &Context{
		MongoSession: session,
	}
	return context
}

func (c *Context) Close() {
	c.MongoSession.Close()
}

// DbCollection returns mgo.collection for the given name
func (c *Context) DbCollection(name string) *mgo.Collection {
	return c.MongoSession.DB(common.AppConfig.Database).C(name)
}