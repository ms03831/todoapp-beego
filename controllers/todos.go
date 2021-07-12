package controllers

import (
	"encoding/json"
	"strconv"

	"github.com/ms03831/todoapp-beego/helpers"
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
		var customData map[string]string
		if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&customData); err != nil {
			c.Data["json"] = map[string]interface{}{"status": -1, "error": err } 
		} else {
			id, err1 := strconv.Atoi(customData["id"])
			err2 := task.Delete(sess, id)
			
			if err1 == nil && err2 == nil {
				c.Data["json"] = map[string]interface{}{"status": 0, "taskId": id}
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
		var customData map[string]string
		if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&customData); err != nil {
			c.Data["json"] = map[string]interface{}{"status": -1, "error": err } 
		} else {
			id, err1 := strconv.Atoi(customData["id"])
			err2 := models.ChangeTaskDone(sess, id)
			
			if err1 == nil && err2 == nil {
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

func (c *TodoController) ChangeDeadline() {
	sess := c.GetSession("user")
	var task models.Task
	if sess != nil {
		var customData map[string]string
		if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&customData); err != nil {
			c.Data["json"] = map[string]interface{}{"status": -1, "error": err } 
		} else {
			deadline, err1 := helpers.TimestampToJavaScriptISO(customData["deadline"])
			id, err2 := strconv.Atoi((customData["id"]))
			err3 := task.ChangeTaskDeadline(sess, id, deadline)
			
			if err1 == nil && err2 == nil && err3 == nil{
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


