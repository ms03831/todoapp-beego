package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/ms03831/todoapp-beego/models"
)

type UserController struct {
	BaseController
}

func (c *UserController) Login() {
	var user models.User
	
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &user)
	
	if err != nil {
        fmt.Println("login error: ", err)
		c.Data["json"] = map[string]interface{}{"status": -1, "error": err } 
    } else {
		user, err := user.GetUserInfo(user.Email)
		if err != nil {
			fmt.Println("login error: ", err)
			c.Data["json"] = map[string]interface{}{"status": -1, "error": err }
		} else {
			c.Data["json"] = map[string]interface{}{"status": 0, "id": user.Id }
		}
	}
	m := make(map[string]interface{})
	m["email"] = user.Email
	c.SetSession("user", m)

    c.ServeJSON()
}

func (c *UserController) Register() {
	var user models.User
	
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, &user); err != nil {
        fmt.Println("registration error: ", err)
		c.Data["json"] = map[string]interface{}{"status": -1, "error": err } 
    } else {
		userid, err := user.Insert()
		if err != nil {
			fmt.Println("registration error: ", err)
			c.Data["json"] = map[string]interface{}{"status": -1, "error": err }
		} else {
			c.Data["json"] = map[string]interface{}{"status": 0, "id": userid }
		}
	}

    c.ServeJSON()
}

func (c *UserController) CreateProfile() {
	c.Data["Website"] = "mudasir.me"
	c.Data["Email"] = "mudasir.hanif@securiti.ai"
	c.TplName = "index.tpl"
}

func (c *UserController) ViewProfile() {
	c.Data["Website"] = "mudasir.me"
	c.Data["Email"] = "mudasir.hanif@securiti.ai"
	c.TplName = "index.tpl"
}

func (c *UserController) UpdateProfile() {
	c.Data["Website"] = "mudasir.me"
	c.Data["Email"] = "mudasir.hanif@securiti.ai"
	c.TplName = "index.tpl"
}

func (c *UserController) Logout() {
	c.Data["Website"] = "mudasir.me"
	c.Data["Email"] = "mudasir.hanif@securiti.ai"
	c.TplName = "index.tpl"
}

