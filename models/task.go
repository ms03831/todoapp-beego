package models

import (
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq" //DATABASE DRIVER
)

type Task struct{
	Id int
	Title string 
	Description string
	Done bool
	Deadline time.Time
	Created orm.DateTimeField `orm:"auto_now_add;type(datetime)"`
	Updated orm.DateTimeField `orm:"auto_add;type(datetime)"`
	User *User `orm:"rel(fk)"`
}

func (m *Task) Insert(session interface{}) (int64, error) {
	var user User
	userSession := session.(map[string]interface{})
	email := fmt.Sprintf("%v", userSession["email"])
	user, _ = user.GetUserInfo(email)
	m.User = &user 
	m.Updated = orm.DateTimeField(time.Now().Local())
	id, err := orm.NewOrm().Insert(m)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (m* Task) GetAllTasks(session interface{}) ([] Task, error){
	var user User
	var tasks []Task
	userSession := session.(map[string]interface{})
	email := fmt.Sprintf("%v", userSession["email"])
	qs := orm.NewOrm().QueryTable("task")
	user, err := user.GetUserInfo(email)
	if err == nil {
		qs.Filter("user_id", user.Id).All(&tasks)
		return tasks, err
	} else {
		return nil, err 
	}
}
