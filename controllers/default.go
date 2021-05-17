package controllers

import (
	"strconv"
	"strings"
	"time"
	"tododat/models"

	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

type TaskController struct {
	beego.Controller
	Action string
	Method string
}

func (c *MainController) Get() {
	doingTasks := models.ReadAllByStatus("doing")
	c.Data["doingTasks"] = doingTasks
	c.TplName = "index.tpl"
}

func (c *TaskController) Post() {
	if c.Action == "done" && c.Method == "post" {
		strTaskId := c.GetString("taskid")
		if strTaskId == "" {
			c.Ctx.WriteString("taskid is empty")
			return
		}
		taskId, _ := strconv.Atoi(strings.Trim(strTaskId, " "))

		models.DoneTask(models.Task{Id: taskId})
		c.Redirect("/", 301)
	} else if c.Action == "add" && c.Method == "post" {
		taskContent := c.GetString("task-content")
		taskProjectId := 1
		createdBy := 0
		strTaskPriority := c.GetString("task-priority")
		taskPriority, _ := strconv.Atoi(strTaskPriority)
		newTask := models.Task{
			Content:   taskContent,
			ProjectId: taskProjectId,
			CreatedAt: time.RFC3339,
			CreatedBy: createdBy,
			Status:    "doing",
			Priority:  taskPriority,
		}
		models.CreateTask(newTask)
		c.Redirect("/", 301)
	}
}