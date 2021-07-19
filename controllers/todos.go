package controllers

import (
	"encoding/json"

	models "github.com/ms03831/todoapp-beego/models"
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


func (c *TodoController) CreateTodo() {
	sess := c.GetSession("user")
	if sess != nil {
		var task models.Task
		if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&task); err != nil {
			c.Data["json"] = map[string]interface{}{"status": -1, "error": err } 
		} else {
			//task.Deadline, err = helpers.TimestampToJavaScriptISO(task.Deadline.String())
			id, err := task.Insert(sess)
			
			if err == nil {
				c.Data["json"] = map[string]interface{}{"status": 0, "taskId":id}
			} else {
				c.Data["json"] = map[string]interface{}{"status": -1, "error": "Some error occured"}
			}
		}
	} else {
			c.Data["json"] = map[string]interface{}{"status": -1, "error": "Please login first"}
	}
	c.ServeJSON()
}

func (c *TodoController) DeleteTodo() {
	sess := c.GetSession("user")
	if sess != nil {
		var task models.Task
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &task)
		if err != nil {
			c.Data["json"] = map[string]interface{}{"status": -1, "error": err } 
		} else {
			err := task.Delete(sess, task.Id)
			
			if err == nil {
				c.Data["json"] = map[string]interface{}{"status": 0, "taskId": task.Id}
			} else {
				c.Data["json"] = map[string]interface{}{"status": -1, "error": "Some error occured"}
			}
		}
	} else {
			c.Data["json"] = map[string]interface{}{"status": -1, "error": "Please login first"}
	}
	c.ServeJSON()
}




func (c *TodoController) ChangeStatus() {
	sess := c.GetSession("user")
	if sess != nil {
		var task models.Task
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &task)
		if err != nil {
			c.Data["json"] = map[string]interface{}{"status": -1, "error": err } 
		} else {
			err := models.ChangeTaskDone(sess, task.Id)
			if err == nil {
				c.Data["json"] = map[string]interface{}{"status": 0, "taskId": task.Id}
			} else {
				c.Data["json"] = map[string]interface{}{"status": -1, "error": "Some error occured"}
			}
		}
	} else {
			c.Data["json"] = map[string]interface{}{"status": -1, "error": "Please login first"}
	}
	c.ServeJSON()
}

func (c *TodoController) ChangeDeadline() {
	sess := c.GetSession("user")
	if sess != nil {
		var task models.Task
		err := json.Unmarshal(c.Ctx.Input.RequestBody, &task)
		if err != nil {
			c.Data["json"] = map[string]interface{}{"status": -1, "error": err } 
		} else {
			err := task.ChangeTaskDeadline(sess, task.Id, task.Deadline)
			
			if err == nil{
				c.Data["json"] = map[string]interface{}{"status": 0, "taskId": task.Id}
			} else {
				c.Data["json"] = map[string]interface{}{"status": -1, "error": "Some error occured"}
			}
		}
	} else {
			c.Data["json"] = map[string]interface{}{"status": -1, "error": "Please login first"}
	}
	c.ServeJSON()
}


