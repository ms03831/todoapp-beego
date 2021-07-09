package models

import (
	"fmt"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq" //DATABASE DRIVER
)

type User struct {
    Id   int
	Email string `orm:"unique"`
	Password string
	Profile *Profile `orm:"reverse(one)"`
	Tasks []*Task `orm:"reverse(many)"`
}


func (m *User) Insert() (int64, error) {
	id, err := orm.NewOrm().Insert(m)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (m *User) GetUserInfo(email string) (User, error) {
	var user User
	qs := orm.NewOrm().QueryTable("user")
	qs = qs.Filter("email", email)
	
	err := qs.One(&user)
	fmt.Println(err)
	if err == nil{
		return user, nil
	}
	return user, err
}
 
func (m *User) Read() error {
	if err := orm.NewOrm().Read(m); err != nil {
		return err
	}
	return nil
}


func (m *User) Update(email string, password string) error {
	if _, err := orm.NewOrm().Update(m, email, password); err != nil {
		return err
	}
	return nil
}

func (m *User) Delete() error {
	if _, err := orm.NewOrm().Delete(m); err != nil {
		return err
	}
	return nil
}

func GetAllUsers() orm.QuerySeter {
	return orm.NewOrm().QueryTable("user")
}