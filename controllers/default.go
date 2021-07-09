package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

type AboutController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "mudasir.me"
	c.Data["Email"] = "mudasir.hanif@securiti.ai"
	c.TplName = "index.tpl"
}

func (c *AboutController) Get() {
	c.Data["Website"] = "mudasir.me"
	c.Data["Email"] = "mudasir.hanif@securiti.ai"
	c.TplName = "index.tpl"
}

