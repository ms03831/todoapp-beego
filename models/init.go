package models

import (
	"fmt"
	"time"

	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq" //DATABASE DRIVER
)


func init() {
    // register model
    orm.RegisterModel(new(User), new(Profile), new(Task))

    // set default database
    orm.RegisterDataBase("default", 
        "postgres",
        "user=mudasir password=1985 host=localhost port=5432 dbname=dbtest sslmode=disable");
	
	orm.RunSyncdb("default", false, true)
}

func main() {
    o := orm.NewOrm()

    user := User{Email: "mudasir@webstack.ca", Password:"1985"}
	profile := Profile{Name: "Mudasir", PhotoPath: "/Users/mudasir/go/hello/static/img/picture.png", User: &user}
	
	deadline1 := time.Date(2021, time.Month(2), 21, 1, 10, 30, 0, time.UTC)
	deadline2 := time.Date(2022, time.Month(4), 21, 1, 10, 30, 0, time.UTC)
	
	task1 := Task{Title: "Task1", Description: "Meeting", Done: false, Deadline: deadline1, User: &user}
	task2 := Task{Title: "Task2", Description: "Check DB", Done: false, Deadline: deadline2, User: &user}
	
	// insert
    id, err := o.Insert(&user)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)
    
	id, err = o.Insert(&profile)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)

    id, err = o.Insert(&task1)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)

    id, err = o.Insert(&task2)
    fmt.Printf("ID: %d, ERR: %v\n", id, err)
}
