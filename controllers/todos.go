package controllers

import (
	"encoding/json"

	"github.com/ms03831/todoapp-beego/models"
)

type TodoController struct {
	BaseController
}

func (c *TodoController) ListTodo() {
	sess := c.GetSession("user")
	if sess != nil {
		var tasks []models.Task
		var task models.Task
		tasks, err := task.GetAllTasks(sess)
		
		if err == nil {
			c.Data["json"] = map[string]interface{}{"status": 0, "tasks":tasks}
		} else{
			c.Data["json"] = map[string]interface{}{"status": -1, "error": "Some error occured"}
		}
	} else {
			c.Data["json"] = map[string]interface{}{"status": -1, "error": "Please login first"}
	}
	c.ServeJSON()
}

func (this *TodoController) CreateTodo() {
	sess := this.GetSession("user")
	if sess != nil {
		var task models.Task
		if err := json.NewDecoder(this.Ctx.Request.Body).Decode(&task); err != nil {
			this.Data["json"] = map[string]interface{}{"status": -1, "error": err } 
		} else {
			//task.Deadline, err = helpers.TimestampToJavaScriptISO(task.Deadline.String())
			id, err := task.Insert(sess)
			
			if err == nil {
				this.Data["json"] = map[string]interface{}{"status": 0, "taskId":id}
			} else {
				this.Data["json"] = map[string]interface{}{"status": -1, "error": "Some error occured"}
			}
		}
	} else {
			this.Data["json"] = map[string]interface{}{"status": -1, "error": "Please login first"}
	}
	this.ServeJSON()
}

func (c *TodoController) UpdateTodo() {
	c.Data["Website"] = "mudasir.me"
	c.Data["Email"] = "mudasir.hanif@securiti.ai"
	c.TplName = "index.tpl"
}

func (c *TodoController) DeleteTodo() {
	c.Data["Website"] = "mudasir.me"
	c.Data["Email"] = "mudasir.hanif@securiti.ai"
	c.TplName = "index.tpl"
}

