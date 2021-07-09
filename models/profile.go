package models

import (
	_ "github.com/lib/pq" //DATABASE DRIVER
)

// Model Struct

type Profile struct{
	Id int
	Name string
	PhotoPath string
	User *User `orm:"null;rel(one);on_delete(set_null)"`
}

func init(){
	//var user User
}
